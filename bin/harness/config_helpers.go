package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"harness/internal/config"
	"harness/internal/llm"
)

func getRoot() string {
	root, err := config.FindWorkspaceRoot()
	if err != nil {
		root, _ = os.Getwd()
	}
	config.LoadEnvFile(root)
	config.LoadEnvFile(filepath.Join(root, "bin", "harness"))
	return root
}

func loadAppCfg(root string) *config.AppConfig {
	cfgPath := filepath.Join(root, ".harness", "config.json")
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		// Dynamic fallback when .harness/config.json is missing
		cfg = buildDefaultAppConfig()
		fmt.Printf("\033[38;5;99m│\033[0m  📡 Usando configuração padrão baseada em variáveis de ambiente (.env)\n")
	} else {
		// Merge custom config with defaults to ensure all standard providers are available
		defaults := buildDefaultAppConfig()
		if cfg.Providers == nil {
			cfg.Providers = make(map[string]config.ProviderConfig)
		}
		for name, defProv := range defaults.Providers {
			if _, exists := cfg.Providers[name]; !exists {
				cfg.Providers[name] = defProv
			}
		}
	}

	// Apply overrides from env vars
	if provider := os.Getenv("HARNESS_PROVIDER"); provider != "" {
		cfg.ActiveProvider = provider
	} else if provider := os.Getenv("ACTIVE_PROVIDER"); provider != "" {
		cfg.ActiveProvider = provider
	}

	// Make sure the active provider exists in our configuration, if not, add a default entry
	if _, exists := cfg.Providers[cfg.ActiveProvider]; !exists {
		cfg.Providers[cfg.ActiveProvider] = defaultProviderConfig(cfg.ActiveProvider)
	}

	if model := os.Getenv("HARNESS_MODEL"); model != "" {
		if pc, exists := cfg.Providers[cfg.ActiveProvider]; exists {
			pc.BodyTemplate = setModelInTemplate(pc.BodyTemplate, model)
			cfg.Providers[cfg.ActiveProvider] = pc
		}
	} else if model := os.Getenv("ACTIVE_MODEL"); model != "" {
		if pc, exists := cfg.Providers[cfg.ActiveProvider]; exists {
			pc.BodyTemplate = setModelInTemplate(pc.BodyTemplate, model)
			cfg.Providers[cfg.ActiveProvider] = pc
		}
	}

	return cfg
}

func defaultProviderConfig(name string) config.ProviderConfig {
	defaults := buildDefaultAppConfig()
	if defProv, exists := defaults.Providers[name]; exists {
		return defProv
	}
	// Generic OpenAI-compatible fallback config
	return config.ProviderConfig{
		URL:          fmt.Sprintf("https://api.%s.com/v1/chat/completions", name),
		Headers:      map[string]string{"Content-Type": "application/json", "Authorization": fmt.Sprintf("Bearer {{%s_API_KEY}}", strings.ToUpper(name))},
		BodyTemplate: fmt.Sprintf(`{"model":"%s-model","messages":[{"role":"user","content":"{{prompt}}"}]}`, name),
		ResponsePath: "choices.0.message.content",
	}
}

func buildDefaultAppConfig() *config.AppConfig {
	ollamaHost := os.Getenv("OLLAMA_HOST")
	if ollamaHost == "" {
		ollamaHost = "http://localhost:11434"
	}
	ollamaModel := os.Getenv("OLLAMA_MODEL")
	if ollamaModel == "" {
		ollamaModel = "glm-4.7-flash:latest"
	}

	providers := map[string]config.ProviderConfig{
		"gemini": {
			URL:          "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-pro:generateContent?key={{GEMINI_API_KEY}}",
			Headers:      map[string]string{"Content-Type": "application/json"},
			BodyTemplate: `{"contents":[{"parts":[{"text":"{{prompt}}"}]}]}`,
			ResponsePath: "candidates.0.content.parts.0.text",
		},
		"deepseek": {
			URL:          "https://api.deepseek.com/chat/completions",
			Headers:      map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{DEEPSEEK_API_KEY}}"},
			BodyTemplate: `{"model":"deepseek-chat","messages":[{"role":"user","content":"{{prompt}}"}]}`,
			ResponsePath: "choices.0.message.content",
			Models: []config.ModelConfig{
				{ID: "deepseek-chat", Name: "DeepSeek Chat", Cost: config.ModelCost{Input: 0.14, Output: 0.28}},
				{ID: "deepseek-reasoner", Name: "DeepSeek Reasoner", Cost: config.ModelCost{Input: 0.55, Output: 2.19}},
			},
		},
		"openai": {
			URL:          "https://api.openai.com/v1/chat/completions",
			Headers:      map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{OPENAI_API_KEY}}"},
			BodyTemplate: `{"model":"gpt-4o","messages":[{"role":"user","content":"{{prompt}}"}]}`,
			ResponsePath: "choices.0.message.content",
		},
		"anthropic": {
			URL:          "https://api.anthropic.com/v1/messages",
			Headers:      map[string]string{"Content-Type": "application/json", "x-api-key": "{{ANTHROPIC_API_KEY}}", "anthropic-version": "2023-06-01"},
			BodyTemplate: `{"model":"claude-3-5-sonnet-20241022","max_tokens":4096,"messages":[{"role":"user","content":"{{prompt}}"}]}`,
			ResponsePath: "content.0.text",
		},
		"groq": {
			URL:          "https://api.groq.com/openai/v1/chat/completions",
			Headers:      map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{GROQ_API_KEY}}"},
			BodyTemplate: `{"model":"llama-3.3-70b-versatile","messages":[{"role":"user","content":"{{prompt}}"}]}`,
			ResponsePath: "choices.0.message.content",
		},
		"openrouter": {
			URL:          "https://openrouter.ai/api/v1/chat/completions",
			Headers:      map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{OPENROUTER_API_KEY}}"},
			BodyTemplate: `{"model":"auto","messages":[{"role":"user","content":"{{prompt}}"}]}`,
			ResponsePath: "choices.0.message.content",
		},
		"ollama": {
			URL:          ollamaHost + "/api/generate",
			Headers:      map[string]string{"Content-Type": "application/json"},
			BodyTemplate: fmt.Sprintf(`{"model":"%s","prompt":"{{prompt}}","stream":false}`, ollamaModel),
			ResponsePath: "response",
		},
	}

	// Auto-detect active provider from environment variables
	active := "gemini" // default fallback
	if os.Getenv("DEEPSEEK_API_KEY") != "" {
		active = "deepseek"
	} else if os.Getenv("GEMINI_API_KEY") != "" {
		active = "gemini"
	} else if os.Getenv("OPENAI_API_KEY") != "" {
		active = "openai"
	} else if os.Getenv("ANTHROPIC_API_KEY") != "" {
		active = "anthropic"
	} else if os.Getenv("GROQ_API_KEY") != "" {
		active = "groq"
	} else if os.Getenv("OPENROUTER_API_KEY") != "" {
		active = "openrouter"
	} else if os.Getenv("OLLAMA_HOST") != "" || os.Getenv("OLLAMA_MODEL") != "" {
		active = "ollama"
	}

	return &config.AppConfig{
		ActiveProvider: active,
		Providers:      providers,
	}
}

// setModelInTemplate replaces the "model" field in a body_template JSON string.
func setModelInTemplate(tmpl, model string) string {
	if tmpl == "" || model == "" {
		return tmpl
	}
	// Match "model": "..." and replace with new model
	oldModel := regexp.MustCompile(`"model"\s*:\s*"[^"]*"`)
	if oldModel.MatchString(tmpl) {
		return oldModel.ReplaceAllString(tmpl, fmt.Sprintf(`"model": "%s"`, model))
	}
	// No model field found: inject before the first message
	return strings.Replace(tmpl, `"messages"`, fmt.Sprintf(`"model": "%s", "messages"`, model), 1)
}

func buildProviderReg(cfg *config.AppConfig) *llm.ProviderRegistry {
	reg := llm.NewProviderRegistry()
	order := registerAvailableProviders(reg, cfg)
	if len(order) == 0 {
		fmt.Println("No providers configured.")
		return nil
	}
	reg.SetFallbackOrder(order)
	return reg
}

// extractCurrentModel returns the current model name for the active provider.
func extractCurrentModel(cfg *config.AppConfig) string {
	if cfg == nil {
		return ""
	}
	pc, ok := cfg.Providers[cfg.ActiveProvider]
	if !ok {
		return ""
	}
	return extractModelFromTemplate(pc.BodyTemplate)
}

func loadMandates(root string) string {
	var b strings.Builder
	for _, mf := range []string{"AGENTS.md", "CLAUDE.md"} {
		mfPath := filepath.Join(root, mf)
		if content, err := os.ReadFile(mfPath); err == nil {
			fmt.Fprintf(&b, "# MANDATES (%s)\n%s\n\n---\n\n", mf, string(content))
		}
	}
	return b.String()
}

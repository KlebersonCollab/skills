package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"harness/internal/config"
	"harness/internal/llm"
)
func registerAvailableProviders(reg *llm.ProviderRegistry, appCfg *config.AppConfig) []string {
	order := make([]string, 0)
	active := appCfg.ActiveProvider
	if p, err := buildSingleProvider(active, appCfg); err == nil {
		reg.Register(p)
		order = append(order, active)
	}
	others := []string{"gemini", "deepseek", "openai", "anthropic", "groq", "openrouter", "together", "fireworks", "mistral", "xai", "cerebras", "deepinfra", "huggingface", "ollama"}
	for _, n := range others {
		if n == active {
			continue
		}
		if _, ok := appCfg.Providers[n]; !ok {
			continue
		}
		if p, err := buildSingleProvider(n, appCfg); err == nil {
			reg.Register(p)
			order = append(order, n)
		}
	}
	return order
}

func buildSingleProvider(name string, appCfg *config.AppConfig) (llm.LLMProvider, error) {
	pc, ok := appCfg.Providers[name]
	if !ok {
		return nil, fmt.Errorf("not configured: %s", name)
	}
	url := config.ResolveEnvVars(pc.URL)

	switch name {
	case "gemini":
		model := extractModelFromTemplate(pc.BodyTemplate)
		return llm.NewGeminiProvider(os.Getenv("GEMINI_API_KEY"), model), nil
	case "deepseek":
		model := extractModelFromTemplate(pc.BodyTemplate)
		return llm.NewDeepSeekProvider(os.Getenv("DEEPSEEK_API_KEY"), model, pc.Models), nil
	case "openai":
		model := extractModelFromTemplate(pc.BodyTemplate)
		return llm.NewOpenAIProvider(os.Getenv("OPENAI_API_KEY"), model), nil
	case "anthropic":
		model := extractModelFromTemplate(pc.BodyTemplate)
		return llm.NewAnthropicProvider(os.Getenv("ANTHROPIC_API_KEY"), model), nil
	case "groq":
		model := extractModelFromTemplate(pc.BodyTemplate)
		return llm.NewGroqProvider(os.Getenv("GROQ_API_KEY"), model), nil
	case "openrouter":
		return oacp("openrouter", "OPENROUTER_API_KEY", "https://openrouter.ai/api/v1/chat/completions", appCfg)
	case "together":
		return oacp("together", "TOGETHER_API_KEY", "https://api.together.xyz/v1/chat/completions", appCfg)
	case "fireworks":
		return oacp("fireworks", "FIREWORKS_API_KEY", "https://api.fireworks.ai/inference/v1/chat/completions", appCfg)
	case "mistral":
		return oacp("mistral", "MISTRAL_API_KEY", "https://api.mistral.ai/v1/chat/completions", appCfg)
	case "xai":
		return oacp("xai", "XAI_API_KEY", "https://api.x.ai/v1/chat/completions", appCfg)
	case "cerebras":
		return oacp("cerebras", "CEREBRAS_API_KEY", "https://api.cerebras.ai/v1/chat/completions", appCfg)
	case "deepinfra":
		return oacp("deepinfra", "DEEPINFRA_API_KEY", "https://api.deepinfra.com/v1/openai/chat/completions", appCfg)
	case "huggingface":
		return oacp("huggingface", "HF_API_KEY", "https://api-inference.huggingface.co/v1/chat/completions", appCfg)
	case "ollama":
		host := url
		if idx := strings.Index(host, "/api/generate"); idx > 0 {
			host = host[:idx]
		}
		model := extractModelFromTemplate(pc.BodyTemplate)
		return llm.NewOllamaProvider(host, model), nil
	default:
		return &genericProvider{cfg: pc}, nil
	}
}

// extractCurrentModel returns the current model name for the active provider.

func oacp(name, envKey, url string, appCfg *config.AppConfig) (llm.LLMProvider, error) {
	key := os.Getenv(envKey)
	if key == "" {
		return nil, fmt.Errorf("%s not set", envKey)
	}
	var model string
	if pc, ok := appCfg.Providers[name]; ok {
		model = extractModelFromTemplate(pc.BodyTemplate)
		url = config.ResolveEnvVars(pc.URL)
	}
	return llm.NewOpenAICompatibleProvider(name, key, url, model), nil
}

// extractModelFromTemplate parses a body_template JSON to find the "model" field.
func extractModelFromTemplate(tmpl string) string {
	if tmpl == "" {
		return ""
	}
	// Replace {{prompt}} and {{history}} with empty strings to make it parseable JSON
	sanitized := strings.ReplaceAll(tmpl, "{{prompt}}", "")
	sanitized = strings.ReplaceAll(sanitized, "{{history}}", "")
	var fields map[string]interface{}
	if err := json.Unmarshal([]byte(sanitized), &fields); err != nil {
		return ""
	}
	if model, ok := fields["model"].(string); ok {
		return model
	}
	return ""
}


type genericProvider struct {
	cfg config.ProviderConfig
}

func (g *genericProvider) Name() string { return "generic" }

func (g *genericProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	body := strings.ReplaceAll(g.cfg.BodyTemplate, "{{prompt}}", jsonEscape(prompt))
	body = strings.ReplaceAll(body, "{{history}}", jsonEscape(""))
	headers := make(map[string]string)
	for k, v := range g.cfg.Headers {
		headers[k] = config.ResolveEnvVars(v)
	}
	resp, err := llm.DoRequest(ctx, g.cfg.URL, headers, []byte(body))
	if err != nil {
		return "", err
	}
	return llm.ExtractToken(string(resp), g.cfg.ResponsePath), nil
}

func (g *genericProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	res, err := g.Complete(ctx, prompt, opts)
	if err != nil {
		return nil, err
	}
	ch := make(chan string, 32)
	go func() {
		defer close(ch)
		words := strings.Split(res, " ")
		for i, w := range words {
			select {
			case <-ctx.Done():
				return
			default:
			}
			chunk := w
			if i < len(words)-1 {
				chunk += " "
			}
			ch <- chunk
			time.Sleep(10 * time.Millisecond) // simulated stream rate
		}
	}()
	return ch, nil
}

func jsonEscape(s string) string {
	b, _ := json.Marshal(s)
	if len(b) >= 2 {
		return string(b[1 : len(b)-1])
	}
	return s
}

// ── init wizard ──────────────────────────────────────────────────────────


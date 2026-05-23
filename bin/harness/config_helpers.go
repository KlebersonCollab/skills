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
		fmt.Println("Config not found. Run 'harness init' first.")
		return nil
	}

	// Apply CLI overrides from env vars
	if provider := os.Getenv("HARNESS_PROVIDER"); provider != "" {
		if _, exists := cfg.Providers[provider]; exists {
			cfg.ActiveProvider = provider
		}
	}
	if model := os.Getenv("HARNESS_MODEL"); model != "" {
		if pc, exists := cfg.Providers[cfg.ActiveProvider]; exists {
			pc.BodyTemplate = setModelInTemplate(pc.BodyTemplate, model)
			cfg.Providers[cfg.ActiveProvider] = pc
		}
	}

	return cfg
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
			b.WriteString("# MANDATES (" + mf + ")\n" + string(content) + "\n\n---\n\n")
		}
	}
	return b.String()
}


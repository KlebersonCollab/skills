package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"harness/internal/config"
)
func runInitWizard() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("LLM Provider Setup")
	root := getRoot()
	configPath := filepath.Join(root, ".harness", "config.json")
	fmt.Print("Provider (1=gemini, 2=ollama, 3=deepseek, 4=openai, 5=anthropic, 6=groq, 7=custom) [1]: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	if choice == "" {
		choice = "1"
	}
	active := ""
	providers := make(map[string]config.ProviderConfig)
	switch choice {
	case "1":
		active = "gemini"
		fmt.Print("API Key: ")
		if k, _ := reader.ReadString('\n'); strings.TrimSpace(k) != "" {
			os.Setenv("GEMINI_API_KEY", strings.TrimSpace(k))
		}
		providers["gemini"] = config.ProviderConfig{URL: "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-pro:generateContent?key={{GEMINI_API_KEY}}", Headers: map[string]string{"Content-Type": "application/json"}, BodyTemplate: `{"contents":[{"parts":[{"text":"{{prompt}}"}]}]}`, ResponsePath: "candidates.0.content.parts.0.text"}
	case "2":
		active = "ollama"
		fmt.Print("Host [http://localhost:11434]: ")
		host, _ := reader.ReadString('\n')
		host = strings.TrimSpace(host)
		if host == "" {
			host = "http://localhost:11434"
		}
		fmt.Print("Model [llama3]: ")
		model, _ := reader.ReadString('\n')
		model = strings.TrimSpace(model)
		if model == "" {
			model = "llama3"
		}
		providers["ollama"] = config.ProviderConfig{URL: host + "/api/generate", Headers: map[string]string{"Content-Type": "application/json"}, BodyTemplate: fmt.Sprintf(`{"model":"%s","prompt":"{{prompt}}","stream":false}`, model), ResponsePath: "response"}
	case "3":
		active = "deepseek"
		fmt.Print("API Key: ")
		if k, _ := reader.ReadString('\n'); strings.TrimSpace(k) != "" {
			os.Setenv("DEEPSEEK_API_KEY", strings.TrimSpace(k))
		}
		providers["deepseek"] = config.ProviderConfig{URL: "https://api.deepseek.com/chat/completions", Headers: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{DEEPSEEK_API_KEY}}"}, BodyTemplate: `{"model":"deepseek-chat","messages":[{"role":"user","content":"{{prompt}}"}]}`, ResponsePath: "choices.0.message.content"}
	case "4":
		active = "openai"
		fmt.Print("API Key: ")
		if k, _ := reader.ReadString('\n'); strings.TrimSpace(k) != "" {
			os.Setenv("OPENAI_API_KEY", strings.TrimSpace(k))
		}
		providers["openai"] = config.ProviderConfig{URL: "https://api.openai.com/v1/chat/completions", Headers: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{OPENAI_API_KEY}}"}, BodyTemplate: `{"model":"gpt-4o","messages":[{"role":"user","content":"{{prompt}}"}]}`, ResponsePath: "choices.0.message.content"}
	case "5":
		active = "anthropic"
		fmt.Print("API Key: ")
		if k, _ := reader.ReadString('\n'); strings.TrimSpace(k) != "" {
			os.Setenv("ANTHROPIC_API_KEY", strings.TrimSpace(k))
		}
		providers["anthropic"] = config.ProviderConfig{URL: "https://api.anthropic.com/v1/messages", Headers: map[string]string{"Content-Type": "application/json", "x-api-key": "{{ANTHROPIC_API_KEY}}", "anthropic-version": "2023-06-01"}, BodyTemplate: `{"model":"claude-sonnet-4-20250514","max_tokens":4096,"messages":[{"role":"user","content":"{{prompt}}"}]}`, ResponsePath: "content.0.text"}
	case "6":
		active = "groq"
		fmt.Print("API Key: ")
		if k, _ := reader.ReadString('\n'); strings.TrimSpace(k) != "" {
			os.Setenv("GROQ_API_KEY", strings.TrimSpace(k))
		}
		providers["groq"] = config.ProviderConfig{URL: "https://api.groq.com/openai/v1/chat/completions", Headers: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer {{GROQ_API_KEY}}"}, BodyTemplate: `{"model":"llama-3.3-70b-versatile","messages":[{"role":"user","content":"{{prompt}}"}]}`, ResponsePath: "choices.0.message.content"}
	case "7":
		active = "custom"
		fmt.Print("Name: ")
		n, _ := reader.ReadString('\n')
		active = strings.TrimSpace(n)
		fmt.Print("URL: ")
		u, _ := reader.ReadString('\n')
		fmt.Print("Body Template: ")
		b, _ := reader.ReadString('\n')
		fmt.Print("Response Path: ")
		r, _ := reader.ReadString('\n')
		providers[active] = config.ProviderConfig{URL: strings.TrimSpace(u), Headers: map[string]string{"Content-Type": "application/json"}, BodyTemplate: strings.TrimSpace(b), ResponsePath: strings.TrimSpace(r)}
	}
	cfg := &config.AppConfig{ActiveProvider: active, Providers: providers}
	os.MkdirAll(filepath.Dir(configPath), 0755)
	config.SaveConfig(configPath, cfg)
	fmt.Printf("Saved: %s\n", configPath)
}

// ── MCP ───────────────────────────────────────────────────────────────────


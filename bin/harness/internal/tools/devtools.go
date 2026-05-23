package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ChromeDevToolsTool implements browser/devtools operations with fallback.
// Supports both attribute XML format and JSON body format.
type ChromeDevToolsTool struct {
	Root string
}

func (c *ChromeDevToolsTool) Execute(name string, args map[string]any) ToolResult {
	// Try to extract params from both attribute format and JSON body format
	action := getStr(args, "action")
	url := getStr(args, "url")
	// The expression parameter is available for future use

	// If action was passed as a JSON body under a "body" or raw "action" key, handle it
	if action == "" {
		if body, ok := args["body"].(string); ok {
			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(body), &parsed); err == nil {
				action, _ = parsed["action"].(string)
				url, _ = parsed["url"].(string)
			}
		}
	}

	if action == "" {
		return ToolResult{Error: fmt.Errorf("missing 'action' parameter. Use: action=screenshot|html|fetch|list-tabs")}
	}

	switch action {
	case "screenshot":
		return c.screenshot(url)
	case "html":
		return c.html(url)
	case "fetch":
		return c.fetch(url)
	case "list-tabs":
		return c.listTabs()
	default:
		return ToolResult{Error: fmt.Errorf("unknown action '%s'. Use: screenshot, html, fetch, list-tabs", action)}
	}
}

// fetch does a simple HTTP GET and returns the page content (always works, no Chrome needed).
func (c *ChromeDevToolsTool) fetch(url string) ToolResult {
	if url == "" {
		return ToolResult{Error: fmt.Errorf("url is required")}
	}

	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("request error: %w", err)}
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; HarnessAgent/1.0)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	resp, err := client.Do(req)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("HTTP error: %w", err)}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 2*1024*1024))
	if err != nil {
		return ToolResult{Error: fmt.Errorf("read error: %w", err)}
	}

	content := string(body)

	// Try to extract title
	title := ""
	if s, e := extractTitle(content); e == nil {
		title = s
	}

	// Strip HTML tags for cleaner output
	text := stripHTML(content)

	// Limit output size
	if len(text) > 5000 {
		text = text[:5000] + "\n\n... (truncado)"
	}

	result := fmt.Sprintf("📄 %s\n%s\n\nConteudo (%d bytes):\n%s", url, title, len(body), text)
	return ToolResult{Output: result}
}

// html tries Chrome DevTools first, falls back to HTTP fetch.
func (c *ChromeDevToolsTool) html(url string) ToolResult {
	if url == "" {
		return ToolResult{Error: fmt.Errorf("url is required")}
	}

	// Try Chrome first
	if chromePath := findChrome(); chromePath != "" {
		// Try headless dump
		cmd := exec.Command(chromePath, "--headless", "--disable-gpu", "--no-sandbox",
			"--dump-dom", url)
		output, err := cmd.Output()
		if err == nil && len(output) > 100 {
			text := string(output)
			if len(text) > 10000 {
				text = text[:10000] + "\n\n... (truncado)"
			}
			return ToolResult{Output: fmt.Sprintf("📄 %s\n\n%s", url, text)}
		}
	}

	// Fallback to HTTP fetch
	return c.fetch(url)
}

// screenshot captures a screenshot using Chrome headless.
func (c *ChromeDevToolsTool) screenshot(url string) ToolResult {
	if url == "" {
		url = "about:blank"
	}

	chromePath := findChrome()
	if chromePath == "" {
		return ToolResult{Error: fmt.Errorf("Chrome not found. Install google-chrome, chromium, or chromium-browser.\nAlternatively, use action=fetch to get the page content as text.")}
	}

	tmpDir, _ := os.MkdirTemp("", "harness-shot-*")
	defer os.RemoveAll(tmpDir)

	screenshotPath := filepath.Join(tmpDir, "shot.png")
	cmd := exec.Command(chromePath,
		"--headless", "--disable-gpu", "--no-sandbox",
		fmt.Sprintf("--screenshot=%s", screenshotPath),
		"--window-size=1280,720", url,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return ToolResult{Error: fmt.Errorf("screenshot failed: %w\n%s", err, string(output))}
	}

	data, err := os.ReadFile(screenshotPath)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("failed to read screenshot: %w", err)}
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	result := fmt.Sprintf("📸 Screenshot de %s (%d KB)\nBase64 (primeiros 100 chars): %.100s...",
		url, len(data)/1024, encoded)

	return ToolResult{Output: result}
}

// listTabs lists Chrome tabs (requires Chrome with --remote-debugging-port=9222).
func (c *ChromeDevToolsTool) listTabs() ToolResult {
	tabs, err := getChromeTabs("http://localhost:9222")
	if err != nil {
		return ToolResult{Error: fmt.Errorf("Chrome DevTools nao disponivel: %w\nCertifique-se de iniciar Chrome com: google-chrome --remote-debugging-port=9222", err)}
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("🔗 %d abas encontradas:\n\n", len(tabs)))
	for _, t := range tabs {
		title := t["title"]
		if len(title) > 80 {
			title = title[:77] + "..."
		}
		result.WriteString(fmt.Sprintf("  • %s\n    URL: %s\n", title, t["url"]))
	}
	return ToolResult{Output: result.String()}
}

// ── Image Tools ────────────────────────────────────────────────────────────

// PasteClipboardImageExecutor reads an image from clipboard and returns base64.
type PasteClipboardImageExecutor struct{}

func (e *PasteClipboardImageExecutor) Execute(name string, args map[string]any) ToolResult {
	data, mime, err := PasteImageFromClipboard()
	if err != nil {
		return ToolResult{Error: fmt.Errorf("clipboard: %w", err)}
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	previewLen := 80
	if len(encoded) < previewLen {
		previewLen = len(encoded)
	}

	result := fmt.Sprintf("📋 Imagem copiada da área de transferência\n\n")
	result += fmt.Sprintf("**Tipo:** %s\n", mime)
	result += fmt.Sprintf("**Tamanho:** %d bytes\n", len(data))
	result += fmt.Sprintf("**Base64 (primeiros %d chars):** %.80s...\n\n", previewLen, encoded)
	result += "Para processar esta imagem, o provedor LLM precisa de suporte multimodal (visão).\n"

	return ToolResult{Output: result}
}

// ImageInfoExecutor reads image metadata from a file path.
type ImageInfoExecutor struct{}

func (e *ImageInfoExecutor) Execute(name string, args map[string]any) ToolResult {
	path, _ := args["path"].(string)
	if path == "" {
		return ToolResult{Error: fmt.Errorf("missing required argument: path")}
	}

	fi, err := os.Stat(path)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("file not found: %s", path)}
	}

	sizeKB := fi.Size() / 1024
	var sizeStr string
	if sizeKB > 1024 {
		sizeStr = fmt.Sprintf("%.1f MB", float64(sizeKB)/1024)
	} else {
		sizeStr = fmt.Sprintf("%d KB", sizeKB)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("📷 **Imagem:** %s\n", filepath.Base(path)))
	b.WriteString(fmt.Sprintf("**Tamanho:** %s\n", sizeStr))

	// file command
	if out, err := exec.Command("file", path).Output(); err == nil {
		meta := strings.TrimSpace(string(out))
		if idx := strings.Index(meta, ":"); idx >= 0 {
			meta = strings.TrimSpace(meta[idx+1:])
		}
		b.WriteString(fmt.Sprintf("**Tipo:** %s\n", meta))
	}

	// ImageMagick identify for dimensions
	if out, err := exec.Command("identify", "-format", "%wx%h", path).Output(); err == nil {
		dims := strings.TrimSpace(string(out))
		b.WriteString(fmt.Sprintf("**Dimensões:** %s\n", dims))
	}

	return ToolResult{Output: b.String()}
}

// ─── Helpers ──────────────────────────────────────────────────────────────

// PasteImageFromClipboard reads an image from clipboard using xclip/wl-paste.
func PasteImageFromClipboard() ([]byte, string, error) {
	if path, err := exec.LookPath("xclip"); err == nil {
		cmd := exec.Command(path, "-selection", "clipboard", "-t", "image/png", "-o")
		data, err := cmd.Output()
		if err == nil && len(data) > 100 {
			return data, "image/png", nil
		}
	}
	if path, err := exec.LookPath("wl-paste"); err == nil {
		cmd := exec.Command(path, "--type", "image/png")
		data, err := cmd.Output()
		if err == nil && len(data) > 100 {
			return data, "image/png", nil
		}
	}
	return nil, "", fmt.Errorf("no image in clipboard or clipboard tool not found")
}

func findChrome() string {
	for _, name := range []string{"google-chrome", "google-chrome-stable", "chromium", "chromium-browser"} {
		if p, err := exec.LookPath(name); err == nil {
			return p
		}
	}
	return ""
}

func getChromeTabs(devtoolsURL string) ([]map[string]string, error) {
	resp, err := http.Get(devtoolsURL + "/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var tabs []map[string]interface{}
	if err := json.Unmarshal(body, &tabs); err != nil {
		return nil, err
	}

	var result []map[string]string
	for _, t := range tabs {
		result = append(result, map[string]string{
			"id":    fmt.Sprintf("%v", t["id"]),
			"title": fmt.Sprintf("%v", t["title"]),
			"url":   fmt.Sprintf("%v", t["url"]),
		})
	}
	return result, nil
}

func getStr(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func extractTitle(html string) (string, error) {
	// Simple title extraction
	idx := strings.Index(strings.ToLower(html), "<title")
	if idx < 0 {
		return "", fmt.Errorf("no title")
	}
	start := strings.Index(html[idx:], ">")
	if start < 0 {
		return "", fmt.Errorf("no title close")
	}
	start += idx + 1
	end := strings.Index(html[start:], "</title")
	if end < 0 {
		return "", fmt.Errorf("no title end")
	}
	return strings.TrimSpace(html[start : start+end]), nil
}

func stripHTML(html string) string {
	var result strings.Builder
	inTag := false
	inScript := false
	for i := 0; i < len(html); i++ {
		if inScript {
			if strings.HasPrefix(html[i:], "</script") {
				inScript = false
			}
			continue
		}
		if strings.HasPrefix(strings.ToLower(html[i:]), "<script") {
			inScript = true
			continue
		}
		if html[i] == '<' {
			inTag = true
			continue
		}
		if html[i] == '>' {
			inTag = false
			continue
		}
		if !inTag {
			result.WriteByte(html[i])
		}
	}
	// Clean up whitespace
	text := result.String()
	text = strings.ReplaceAll(text, "\n\n\n", "\n\n")
	text = strings.ReplaceAll(text, "  ", " ")
	return strings.TrimSpace(text)
}

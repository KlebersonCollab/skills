package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// FetchContentExecutor implements the fetch_content tool for URLs, YouTube, GitHub, and local videos.
type FetchContentExecutor struct{}

func (e *FetchContentExecutor) Execute(name string, args map[string]any) ToolResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Single URL
	if url, ok := args["url"].(string); ok && url != "" {
		return e.fetchURL(ctx, url, args)
	}

	// Multiple URLs
	if urlsRaw, ok := args["urls"].([]interface{}); ok && len(urlsRaw) > 0 {
		var urls []string
		for _, u := range urlsRaw {
			if s, ok := u.(string); ok {
				urls = append(urls, s)
			}
		}
		return e.fetchURLs(ctx, urls, args)
	}

	// Local video file
	if path, ok := args["path"].(string); ok && path != "" {
		return e.fetchLocalVideo(ctx, path, args)
	}

	return ToolResult{Error: fmt.Errorf("missing required argument: url, urls, or path")}
}

func (e *FetchContentExecutor) fetchURL(ctx context.Context, urlStr string, args map[string]any) ToolResult {
	// YouTube detection
	if isYouTubeURL(urlStr) {
		prompt, _ := args["prompt"].(string)
		return ToolResult{Output: e.fetchYouTube(ctx, urlStr, prompt)}
	}

	// GitHub repo detection
	if isGitHubRepoURL(urlStr) {
		forceClone := false
		if fc, ok := args["forceClone"].(bool); ok {
			forceClone = fc
		}
		return e.fetchGitHub(ctx, urlStr, forceClone)
	}

	// Generic URL
	return e.fetchGenericURL(ctx, urlStr)
}

func (e *FetchContentExecutor) fetchURLs(ctx context.Context, urls []string, args map[string]any) ToolResult {
	var results []string
	for _, u := range urls {
		res := e.fetchURL(ctx, u, args)
		if res.Error != nil {
			results = append(results, fmt.Sprintf("## ❌ Erro ao buscar %s\n%s", u, res.Error.Error()))
		} else {
			results = append(results, res.Output)
		}
	}
	return ToolResult{Output: strings.Join(results, "\n\n---\n\n")}
}

func (e *FetchContentExecutor) fetchGenericURL(ctx context.Context, urlStr string) ToolResult {
	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("request creation: %w", err)}
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; HarnessAgent/1.0)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("http fetch: %w", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ToolResult{Error: fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)}
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 5*1024*1024))
	if err != nil {
		return ToolResult{Error: fmt.Errorf("read body: %w", err)}
	}

	contentType := resp.Header.Get("Content-Type")
	content := string(body)

	// Simple markdown extraction for HTML
	if strings.Contains(contentType, "text/html") {
		content = htmlToMarkdown(content)
	} else if strings.Contains(contentType, "application/json") {
		// Pretty-print JSON
		var pretty bytes.Buffer
		if err := json.Indent(&pretty, body, "", "  "); err == nil {
			content = "```json\n" + pretty.String() + "\n```"
		}
	}

	// Build result with metadata
	result := fmt.Sprintf("## 📄 Conteúdo de: %s\n\n**Tipo:** %s\n**Tamanho:** %d bytes\n\n%s",
		urlStr, contentType, len(body), content)

	return ToolResult{Output: result}
}

func (e *FetchContentExecutor) fetchYouTube(ctx context.Context, urlStr string, prompt string) string {
	// Try yt-dlp for transcript extraction
	ytCmd := exec.CommandContext(ctx, "yt-dlp", "--skip-download", "--print", "%(title)s\n%(duration)s", urlStr)
	metaOut, err := ytCmd.Output()
	
	var result strings.Builder
	result.WriteString(fmt.Sprintf("## 🎬 YouTube: %s\n\n", urlStr))

	if err == nil {
		metaLines := strings.SplitN(strings.TrimSpace(string(metaOut)), "\n", 2)
		if len(metaLines) >= 1 {
			result.WriteString(fmt.Sprintf("**Título:** %s\n", metaLines[0]))
		}
		if len(metaLines) >= 2 {
			durationSec := strings.TrimSpace(metaLines[1])
			result.WriteString(fmt.Sprintf("**Duração:** %s segundos\n\n", durationSec))
		}
	}

	// Get transcript via yt-dlp
	transcriptCmd := exec.CommandContext(ctx, "yt-dlp", "--skip-download", "--print", "%(description)s", "--write-auto-subs", "--sub-langs", "en,pt", "--sub-format", "vtt/txt", "--convert-subs", "txt", urlStr)
	transcriptOut, transcriptErr := transcriptCmd.Output()
	
	if transcriptErr == nil && len(transcriptOut) > 0 {
		result.WriteString("**📝 Transcrição/Descrição:**\n\n")
		result.WriteString(string(transcriptOut))
	} else {
		result.WriteString("*Transcrição não disponível (yt-dlp não encontrado ou sem legendas).*\n")
		result.WriteString("\nPara análise completa, configure yt-dlp e ffmpeg.\n")
	}

	if prompt != "" {
		result.WriteString(fmt.Sprintf("\n\n**🎯 Pergunta do usuário sobre o vídeo:**\n%s\n", prompt))
	}

	return result.String()
}

func (e *FetchContentExecutor) fetchGitHub(ctx context.Context, urlStr string, forceClone bool) ToolResult {
	// Parse GitHub URL
	parts := strings.Split(strings.TrimPrefix(urlStr, "https://github.com/"), "/")
	if len(parts) < 2 {
		return ToolResult{Error: fmt.Errorf("invalid GitHub URL: %s", urlStr)}
	}
	owner, repo := parts[0], strings.TrimSuffix(parts[1], ".git")

	// Use GitHub API to get repo contents
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", owner, repo)
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("github api request: %w", err)}
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "harness-agent")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("github api: %w", err)}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	result := fmt.Sprintf("## 📦 GitHub: %s/%s\n\n", owner, repo)
	
	// Try to parse README as well
	readmeURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/readme", owner, repo)
	readmeReq, _ := http.NewRequestWithContext(ctx, "GET", readmeURL, nil)
	readmeReq.Header.Set("Accept", "application/vnd.github.v3.raw")
	readmeReq.Header.Set("User-Agent", "harness-agent")
	
	if readmeResp, readmeErr := http.DefaultClient.Do(readmeReq); readmeErr == nil {
		defer readmeResp.Body.Close()
		if readmeResp.StatusCode == 200 {
			readmeBody, _ := io.ReadAll(io.LimitReader(readmeResp.Body, 100*1024))
			result += "**📖 README:**\n\n```markdown\n" + string(readmeBody) + "\n```\n\n"
		}
	}

	// Parse directory listing
	var contents []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	if err := json.Unmarshal(body, &contents); err == nil {
		result += fmt.Sprintf("**📁 Conteúdo da raiz (%d itens):**\n\n", len(contents))
		for _, item := range contents {
			icon := "📄"
			if item.Type == "dir" {
				icon = "📁"
			}
			result += fmt.Sprintf("%s %s\n", icon, item.Name)
		}
	}

	return ToolResult{Output: result}
}

func (e *FetchContentExecutor) fetchLocalVideo(ctx context.Context, path string, args map[string]any) ToolResult {
	// Verify file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ToolResult{Error: fmt.Errorf("file not found: %s", path)}
	}

	prompt, _ := args["prompt"].(string)
	timestamp, _ := args["timestamp"].(string)
	frames := 6
	if f, ok := args["frames"].(float64); ok {
		frames = int(f)
	}

	result := fmt.Sprintf("## 🎬 Vídeo Local: %s\n\n**Caminho:** %s\n", filepath.Base(path), path)

	if prompt != "" {
		result += fmt.Sprintf("**Pergunta:** %s\n\n", prompt)
	}

	// Extract frame(s) with ffmpeg if timestamp provided
	if timestamp != "" || frames > 0 {
		ffmpegPath, err := exec.LookPath("ffmpeg")
		if err != nil {
			result += "*ffmpeg não encontrado. Instale ffmpeg para extração de frames.*\n"
		} else {
			frameDir, _ := os.MkdirTemp("", "harness-frames-*")
			defer os.RemoveAll(frameDir)

			if timestamp != "" {
				// Single frame at timestamp
				frameFile := filepath.Join(frameDir, "frame_%03d.jpg")
				args := []string{
					"-ss", timestamp,
					"-i", path,
					"-vframes", fmt.Sprintf("%d", frames),
					"-q:v", "2",
					frameFile,
				}
				if out, err := exec.Command(ffmpegPath, args...).CombinedOutput(); err == nil {
					result += fmt.Sprintf("✅ Frame(s) extraído(s) em: %s\n", frameDir)
					_ = out
				} else {
					result += fmt.Sprintf("⚠️ Erro ao extrair frame: %s\n", string(out))
				}
			}
		}
	}

	// Get file info
	if fi, err := os.Stat(path); err == nil {
		sizeMB := float64(fi.Size()) / (1024 * 1024)
		result += fmt.Sprintf("**Tamanho:** %.1f MB\n", sizeMB)
	}

	return ToolResult{Output: result}
}

// --- Helpers ---

func isYouTubeURL(u string) bool {
	patterns := []string{
		"youtube.com/watch",
		"youtu.be/",
		"youtube.com/shorts/",
		"m.youtube.com/watch",
	}
	for _, p := range patterns {
		if strings.Contains(u, p) {
			return true
		}
	}
	return false
}

func isGitHubRepoURL(u string) bool {
	return strings.Contains(u, "github.com/") && strings.Count(u, "/") >= 2
}

var (
	htmlTagRE      = regexp.MustCompile(`<[^>]*>`)
	htmlScriptRE   = regexp.MustCompile(`(?si)<script[^>]*>.*?</script>`)
	htmlStyleRE    = regexp.MustCompile(`(?si)<style[^>]*>.*?</style>`)
	htmlNavRE      = regexp.MustCompile(`(?si)<nav[^>]*>.*?</nav>`)
	htmlNewlineRE  = regexp.MustCompile(`\n{3,}`)
	htmlSpaceRE    = regexp.MustCompile(`[ \t]{2,}`)
	htmlHeaderRE   = regexp.MustCompile(`(?si)<h([1-6])[^>]*>(.*?)</h[1-6]>`)
	htmlAnchorRE   = regexp.MustCompile(`(?si)<a[^>]*href="([^"]*)"[^>]*>(.*?)</a>`)
	htmlCodeRE     = regexp.MustCompile(`(?si)<code[^>]*>(.*?)</code>`)
	htmlPreRE      = regexp.MustCompile(`(?si)<pre[^>]*>(.*?)</pre>`)
	htmlLiRE       = regexp.MustCompile(`(?si)<li[^>]*>(.*?)</li>`)
	htmlBrRE       = regexp.MustCompile(`(?si)<br\s*/?>`)
	htmlTitleRE    = regexp.MustCompile(`(?si)<title[^>]*>(.*?)</title>`)
	htmlPEndRE     = regexp.MustCompile(`</p>`)
)

func htmlToMarkdown(html string) string {
	// Remove scripts, styles, nav
	html = htmlScriptRE.ReplaceAllString(html, "")
	html = htmlStyleRE.ReplaceAllString(html, "")
	html = htmlNavRE.ReplaceAllString(html, "")

	// Extract title
	title := ""
	if m := htmlTitleRE.FindStringSubmatch(html); len(m) > 1 {
		title = strings.TrimSpace(m[1])
	}

	// Convert headers
	html = htmlHeaderRE.ReplaceAllString(html, "\n\n### $2\n\n")

	// Convert links
	html = htmlAnchorRE.ReplaceAllString(html, "$2 ($1)")

	// Convert code blocks
	html = htmlPreRE.ReplaceAllString(html, "\n```\n$1\n```\n")
	html = htmlCodeRE.ReplaceAllString(html, "`$1`")

	// Convert lists
	html = htmlLiRE.ReplaceAllString(html, "- $1\n")

	// Convert br and p
	html = htmlBrRE.ReplaceAllString(html, "\n")
	html = htmlPEndRE.ReplaceAllString(html, "\n\n")

	// Strip remaining tags
	html = htmlTagRE.ReplaceAllString(html, "")

	// Clean whitespace
	html = htmlSpaceRE.ReplaceAllString(html, " ")
	html = htmlNewlineRE.ReplaceAllString(html, "\n\n")

	html = strings.TrimSpace(html)

	if title != "" {
		html = fmt.Sprintf("# %s\n\n%s", title, html)
	}

	return html
}

// Package tools provides code search via GitHub and Stack Overflow APIs.
package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CodeSearchResult holds the outcome of a code search.
type CodeSearchResult struct {
	Output string
	Error  error
}

// CodeSearchExecutor implements the code_search tool.
type CodeSearchExecutor struct {
	GitHubToken string // Optional: GITHUB_TOKEN for higher rate limits
}

func (e *CodeSearchExecutor) Execute(name string, args map[string]any) ToolResult {
	query, _ := args["query"].(string)
	if query == "" {
		return ToolResult{Error: fmt.Errorf("missing required argument: query")}
	}

	maxTokens := 10000
	if mt, ok := args["maxTokens"].(float64); ok {
		maxTokens = int(mt)
	}

	result := e.search(query, maxTokens)
	if result.Error != nil {
		return ToolResult{Error: result.Error}
	}
	return ToolResult{Output: result.Output}
}

func (e *CodeSearchExecutor) search(query string, maxTokens int) CodeSearchResult {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var allParts []string

	// 1. Search GitHub
	ghResults, err := e.searchGitHub(ctx, query, 5)
	if err == nil && len(ghResults) > 0 {
		allParts = append(allParts, "## GitHub Results\n")
		for _, r := range ghResults {
			allParts = append(allParts, r)
		}
	}

	// 2. Search Stack Overflow
	soResults, err := e.searchStackOverflow(ctx, query, 5)
	if err == nil && len(soResults) > 0 {
		allParts = append(allParts, "\n## Stack Overflow Results\n")
		for _, r := range soResults {
			allParts = append(allParts, r)
		}
	}

	// 3. If both failed, return error
	if len(ghResults) == 0 && len(soResults) == 0 {
		return CodeSearchResult{Error: fmt.Errorf("no results found for: %s", query)}
	}

	// Truncate to maxTokens estimate (1 token ≈ 4 chars)
	full := strings.Join(allParts, "\n")
	if len(full) > maxTokens*4 {
		full = full[:maxTokens*4] + "\n\n... (truncado por limite de tokens)"
	}

	return CodeSearchResult{Output: full}
}

func (e *CodeSearchExecutor) searchGitHub(ctx context.Context, query string, maxResults int) ([]string, error) {
	apiURL := fmt.Sprintf("https://api.github.com/search/code?q=%s&per_page=%d", url.QueryEscape(query), maxResults)

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("github request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "harness-agent")
	if e.GitHubToken != "" {
		req.Header.Set("Authorization", "Bearer "+e.GitHubToken)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("github api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("github api: status %d", resp.StatusCode)
	}

	var ghResp struct {
		Items []struct {
			Name       string `json:"name"`
			Path       string `json:"path"`
			HTMLURL    string `json:"html_url"`
			Repository struct {
				FullName string `json:"full_name"`
			} `json:"repository"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&ghResp); err != nil {
		return nil, fmt.Errorf("github parse: %w", err)
	}

	var results []string
	for _, item := range ghResp.Items {
		results = append(results, fmt.Sprintf("📁 %s/%s\n   📄 %s\n   🔗 %s\n",
			item.Repository.FullName, item.Path, item.Name, item.HTMLURL))
	}

	return results, nil
}

func (e *CodeSearchExecutor) searchStackOverflow(ctx context.Context, query string, maxResults int) ([]string, error) {
	apiURL := fmt.Sprintf("https://api.stackexchange.com/2.3/search?order=desc&sort=relevance&q=%s&pagesize=%d&site=stackoverflow",
		url.QueryEscape(query), maxResults)

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("stackoverflow request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("stackoverflow api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("stackoverflow api: status %d", resp.StatusCode)
	}

	var soResp struct {
		Items []struct {
			Title string `json:"title"`
			Link  string `json:"link"`
			Score int    `json:"score"`
			Tags  []string `json:"tags"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&soResp); err != nil {
		return nil, fmt.Errorf("stackoverflow parse: %w", err)
	}

	var results []string
	for _, item := range soResp.Items {
		tags := strings.Join(item.Tags, ", ")
		results = append(results, fmt.Sprintf("❓ %s\n   👍 Score: %d | 🏷️ %s\n   🔗 %s\n",
			item.Title, item.Score, tags, item.Link))
	}

	return results, nil
}

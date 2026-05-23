package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// GeminiWebProvider implements WebSearchProvider using Gemini's search grounding.
type GeminiWebProvider struct {
	apiKey string
}

// NewGeminiWebProvider creates a new Gemini web search provider.
func NewGeminiWebProvider(apiKey string) *GeminiWebProvider {
	return &GeminiWebProvider{apiKey: apiKey}
}

func (g *GeminiWebProvider) Name() string { return "gemini" }

func (g *GeminiWebProvider) Search(ctx context.Context, query string, opts WebSearchOpts) (*WebSearchResult, error) {
	model := os.Getenv("GEMINI_WEB_MODEL")
	if model == "" {
		model = os.Getenv("GEMINI_MODEL")
	}
	if model == "" {
		model = "gemini-2.0-flash-exp"
	}
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, g.apiKey)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// Gemini search grounding via Google Search
	bodyMap := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": query},
				},
			},
		},
		"tools": []map[string]interface{}{
			{
				"google_search": map[string]interface{}{},
			},
		},
	}

	bodyBytes, _ := json.Marshal(bodyMap)

	respBody, err := DoRequest(ctx, url, headers, bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("gemini-web: %w", err)
	}

	answer := ExtractToken(string(respBody), "candidates.0.content.parts.0.text")
	if answer == "" {
		return nil, fmt.Errorf("gemini-web: empty response")
	}

	// Extract grounding chunks (sources) from the response
	var sources []string
	var rawResp map[string]interface{}
	if err := json.Unmarshal(respBody, &rawResp); err == nil {
		if candidates, ok := rawResp["candidates"].([]interface{}); ok && len(candidates) > 0 {
			if candidate, ok := candidates[0].(map[string]interface{}); ok {
				if grounding, ok := candidate["groundingChunks"].([]interface{}); ok {
					for _, chunk := range grounding {
						if c, ok := chunk.(map[string]interface{}); ok {
							if web, ok := c["web"].(map[string]interface{}); ok {
								if uri, ok := web["uri"].(string); ok {
									sources = append(sources, uri)
								}
							}
						}
					}
				}
			}
		}
	}

	if len(sources) > 0 {
		answer += "\n\n**Fontes consultadas:**\n"
		for i, src := range sources {
			answer += fmt.Sprintf("%d. %s\n", i+1, src)
		}
	}

	return &WebSearchResult{
		Answer:    answer,
		Sources:   sources,
		Provider:  "gemini",
		QueryUsed: query,
	}, nil
}

func (g *GeminiWebProvider) SearchParallel(ctx context.Context, queries []string, opts WebSearchOpts) (*WebSearchResult, error) {
	mergedQuery := "Responda cada uma das seguintes questões de forma separada:\n"
	for i, q := range queries {
		mergedQuery += fmt.Sprintf("\nQuestão %d: %s", i+1, q)
	}

	return g.Search(ctx, mergedQuery, opts)
}

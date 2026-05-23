package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// PerplexityProvider implements WebSearchProvider for Perplexity AI.
type PerplexityProvider struct {
	apiKey string
}

// NewPerplexityProvider creates a new Perplexity web search provider.
func NewPerplexityProvider(apiKey string) *PerplexityProvider {
	return &PerplexityProvider{apiKey: apiKey}
}

func (p *PerplexityProvider) Name() string { return "perplexity" }

func (p *PerplexityProvider) Search(ctx context.Context, query string, opts WebSearchOpts) (*WebSearchResult, error) {
	url := "https://api.perplexity.ai/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + p.apiKey,
	}

	searchCtx := "Responda à pergunta do usuário de forma completa e precisa. Inclua citações das fontes utilizadas no formato [1], [2], etc."
	
	bodyMap := map[string]interface{}{
		"model": "sonar-pro",
		"messages": []map[string]string{
			{"role": "system", "content": searchCtx},
			{"role": "user", "content": query},
		},
		"max_tokens": 2000,
	}

	if opts.NumResults > 0 {
		bodyMap["max_tokens"] = opts.NumResults * 400
	}

	bodyBytes, _ := json.Marshal(bodyMap)

	respBody, err := DoRequest(ctx, url, headers, bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("perplexity: %w", err)
	}

	answer := ExtractToken(string(respBody), "choices.0.message.content")
	if answer == "" {
		return nil, fmt.Errorf("perplexity: empty response")
	}

	// Extract citations if present
	var sources []string
	var rawResp map[string]interface{}
	if err := json.Unmarshal(respBody, &rawResp); err == nil {
		if choices, ok := rawResp["choices"].([]interface{}); ok && len(choices) > 0 {
			if choice, ok := choices[0].(map[string]interface{}); ok {
				if message, ok := choice["message"].(map[string]interface{}); ok {
					if citations, ok := message["citations"].([]interface{}); ok {
						for _, c := range citations {
							if cs, ok := c.(string); ok {
								sources = append(sources, cs)
							}
						}
					}
				}
			}
		}
	}

	return &WebSearchResult{
		Answer:    answer,
		Sources:   sources,
		Provider:  "perplexity",
		QueryUsed: query,
	}, nil
}

func (p *PerplexityProvider) SearchParallel(ctx context.Context, queries []string, opts WebSearchOpts) (*WebSearchResult, error) {
	// For Perplexity, we merge multiple queries into a single enhanced query
	mergedQuery := "Pesquise e responda às seguintes questões:\n"
	for i, q := range queries {
		mergedQuery += fmt.Sprintf("%d. %s\n", i+1, q)
	}

	return p.Search(ctx, mergedQuery, opts)
}

// ExaProvider implements WebSearchProvider for Exa.
type ExaProvider struct {
	apiKey string
}

func NewExaProvider(apiKey string) *ExaProvider {
	return &ExaProvider{apiKey: apiKey}
}

func (e *ExaProvider) Name() string { return "exa" }

func (e *ExaProvider) Search(ctx context.Context, query string, opts WebSearchOpts) (*WebSearchResult, error) {
	url := "https://api.exa.ai/search"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"x-api-key":     e.apiKey,
		"Authorization": "Bearer " + e.apiKey,
	}

	numResults := opts.NumResults
	if numResults <= 0 {
		numResults = 5
	}

	bodyMap := map[string]interface{}{
		"query": query,
		"numResults": numResults,
		"useWebSource": true,
		"contents": map[string]interface{}{
			"text": true,
		},
	}

	if opts.IncludeContent {
		bodyMap["contents"] = map[string]interface{}{
			"text": map[string]interface{}{
				"maxCharacters": 2000,
			},
		}
	}

	if opts.RecencyFilter != "" {
		bodyMap["recency"] = opts.RecencyFilter
	}

	if len(opts.DomainFilter) > 0 {
		var include, exclude []string
		for _, d := range opts.DomainFilter {
			if strings.HasPrefix(d, "-") {
				exclude = append(exclude, strings.TrimPrefix(d, "-"))
			} else {
				include = append(include, d)
			}
		}
		if len(include) > 0 {
			bodyMap["includeDomains"] = include
		}
		if len(exclude) > 0 {
			bodyMap["excludeDomains"] = exclude
		}
	}

	bodyBytes, _ := json.Marshal(bodyMap)

	respBody, err := DoRequest(ctx, url, headers, bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("exa: %w", err)
	}

	// Parse Exa response
	var exaResp struct {
		Results []struct {
			Title       string `json:"title"`
			URL         string `json:"url"`
			Text        string `json:"text"`
			PublishedDate string `json:"publishedDate"`
			Score       float64 `json:"score"`
		} `json:"results"`
	}

	if err := json.Unmarshal(respBody, &exaResp); err != nil {
		return nil, fmt.Errorf("exa: failed to parse response: %w", err)
	}

	if len(exaResp.Results) == 0 {
		return &WebSearchResult{
			Answer:    "Nenhum resultado encontrado para: " + query,
			Provider:  "exa",
			QueryUsed: query,
		}, nil
	}

	// Build answer
	var answer strings.Builder
	answer.WriteString(fmt.Sprintf("🔍 Resultados da busca por: %s\n\n", query))
	answer.WriteString(fmt.Sprintf("Encontrados %d resultados:\n\n", len(exaResp.Results)))

	var sources []string
	for i, r := range exaResp.Results {
		sources = append(sources, r.URL)
		answer.WriteString(fmt.Sprintf("**[%d] %s**\n", i+1, r.Title))
		answer.WriteString(fmt.Sprintf("   Fonte: %s\n", r.URL))
		if r.PublishedDate != "" {
			answer.WriteString(fmt.Sprintf("   Data: %s\n", r.PublishedDate))
		}
		if r.Text != "" {
			preview := r.Text
			if len(preview) > 300 {
				preview = preview[:297] + "..."
			}
			answer.WriteString(fmt.Sprintf("   » %s\n", preview))
		}
		answer.WriteString("\n")
	}

	return &WebSearchResult{
		Answer:    answer.String(),
		Sources:   sources,
		Provider:  "exa",
		QueryUsed: query,
	}, nil
}

func (e *ExaProvider) SearchParallel(ctx context.Context, queries []string, opts WebSearchOpts) (*WebSearchResult, error) {
	// For Exa, run first query (most relevant)
	return e.Search(ctx, queries[0], opts)
}

// Package llm provides web search capabilities via multiple providers.
package llm

import (
	"context"
	"fmt"
	"strings"
)

// WebSearchProvider defines the interface for web search backends.
type WebSearchProvider interface {
	// Name returns the provider identifier.
	Name() string
	
	// Search performs a single query and returns synthesized results with sources.
	Search(ctx context.Context, query string, opts WebSearchOpts) (*WebSearchResult, error)

	// SearchParallel performs multiple queries in parallel and merges results.
	SearchParallel(ctx context.Context, queries []string, opts WebSearchOpts) (*WebSearchResult, error)
}

// WebSearchOpts defines options for web search.
type WebSearchOpts struct {
	NumResults    int      `json:"numResults,omitempty"`
	RecencyFilter string   `json:"recencyFilter,omitempty"` // day, week, month, year
	DomainFilter  []string `json:"domainFilter,omitempty"`  // "stackoverflow.com", "-reddit.com"
	IncludeContent bool    `json:"includeContent,omitempty"`
	MaxTokens     int      `json:"maxTokens,omitempty"`
}

// WebSearchResult holds the search response.
type WebSearchResult struct {
	Answer    string   `json:"answer"`
	Sources   []string `json:"sources"`
	Provider  string   `json:"provider"`
	QueryUsed string   `json:"queryUsed"`
}

// WebSearchRegistry manages web search providers with fallback.
type WebSearchRegistry struct {
	providers []WebSearchProvider
}

// NewWebSearchRegistry creates a registry with the default providers.
func NewWebSearchRegistry(perplexityKey, exaKey, geminiKey string) *WebSearchRegistry {
	r := &WebSearchRegistry{}

	if perplexityKey != "" {
		r.providers = append(r.providers, NewPerplexityProvider(perplexityKey))
	}
	if exaKey != "" {
		r.providers = append(r.providers, NewExaProvider(exaKey))
	}
	// Gemini web provider always available if GEMINI_API_KEY is set (checked at call time)
	if geminiKey != "" {
		r.providers = append(r.providers, NewGeminiWebProvider(geminiKey))
	}

	return r
}

// Search tries each provider in order until one succeeds (fallback chain).
func (r *WebSearchRegistry) Search(ctx context.Context, query string, opts WebSearchOpts) (*WebSearchResult, error) {
	if len(r.providers) == 0 {
		return r.fallbackSearch(ctx, query, opts)
	}

	var lastErr error
	for _, p := range r.providers {
		result, err := p.Search(ctx, query, opts)
		if err == nil {
			return result, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("all web search providers failed: %w", lastErr)
}

// SearchParallel performs multiple queries via the first available provider.
func (r *WebSearchRegistry) SearchParallel(ctx context.Context, queries []string, opts WebSearchOpts) (*WebSearchResult, error) {
	if len(r.providers) == 0 {
		return r.fallbackSearch(ctx, strings.Join(queries, " "), opts)
	}

	provider := r.providers[0]
	return provider.SearchParallel(ctx, queries, opts)
}

// fallbackSearch is a last resort when no API keys are configured.
func (r *WebSearchRegistry) fallbackSearch(ctx context.Context, query string, opts WebSearchOpts) (*WebSearchResult, error) {
	return &WebSearchResult{
		Answer:    fmt.Sprintf("Web search não disponível. Nenhuma chave de API configurada.\n\nPara ativar, configure alguma destas variáveis de ambiente:\n  PERPLEXITY_API_KEY  (recomendado)\n  EXA_API_KEY\n  GEMINI_API_KEY"),
		Sources:   []string{},
		Provider:  "none",
		QueryUsed: query,
	}, nil
}

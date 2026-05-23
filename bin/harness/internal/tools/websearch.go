package tools

import (
	"context"
	"fmt"
	"strings"
	"time"

	"harness/internal/llm"
)

// WebSearchExecutor implements the web_search tool.
type WebSearchExecutor struct {
	Registry *llm.WebSearchRegistry
}

func (e *WebSearchExecutor) Execute(name string, args map[string]any) ToolResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Parse options
	opts := llm.WebSearchOpts{}

	if nr, ok := args["numResults"].(float64); ok {
		opts.NumResults = int(nr)
	}
	if rf, ok := args["recencyFilter"].(string); ok {
		opts.RecencyFilter = rf
	}
	if df, ok := args["domainFilter"].([]interface{}); ok {
		for _, d := range df {
			if s, ok := d.(string); ok {
				opts.DomainFilter = append(opts.DomainFilter, s)
			}
		}
	}
	if ic, ok := args["includeContent"].(bool); ok {
		opts.IncludeContent = ic
	}

	// Single query
	if query, ok := args["query"].(string); ok && query != "" {
		return e.singleSearch(ctx, query, opts)
	}

	// Multiple queries
	if queriesRaw, ok := args["queries"].([]interface{}); ok && len(queriesRaw) > 0 {
		var queries []string
		for _, q := range queriesRaw {
			if s, ok := q.(string); ok {
				queries = append(queries, s)
			}
		}
		return e.multiSearch(ctx, queries, opts)
	}

	return ToolResult{Error: fmt.Errorf("missing required argument: query or queries")}
}

func (e *WebSearchExecutor) singleSearch(ctx context.Context, query string, opts llm.WebSearchOpts) ToolResult {
	if e.Registry == nil {
		return ToolResult{Output: "Web search not configured. Run 'harness init' to set up API keys."}
	}

	result, err := e.Registry.Search(ctx, query, opts)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("web search failed: %w", err)}
	}

	output := formatResult(result)
	return ToolResult{Output: output}
}

func (e *WebSearchExecutor) multiSearch(ctx context.Context, queries []string, opts llm.WebSearchOpts) ToolResult {
	if e.Registry == nil {
		return ToolResult{Output: "Web search not configured. Run 'harness init' to set up API keys."}
	}

	// Try parallel search first
	result, err := e.Registry.SearchParallel(ctx, queries, opts)
	if err == nil && result != nil {
		return ToolResult{Output: formatResult(result)}
	}

	// Fallback: individual searches
	var parts []string
	for i, q := range queries {
		res, err := e.Registry.Search(ctx, q, opts)
		if err != nil {
			parts = append(parts, fmt.Sprintf("## Query %d: %s\n❌ Erro: %v", i+1, q, err))
		} else {
			parts = append(parts, formatResult(res))
		}
	}

	return ToolResult{Output: strings.Join(parts, "\n\n---\n\n")}
}

func formatResult(r *llm.WebSearchResult) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("🔍 **Fonte:** %s\n", strings.ToUpper(r.Provider)))

	if r.QueryUsed != "" {
		b.WriteString(fmt.Sprintf("**Consulta:** %s\n\n", r.QueryUsed))
	}

	b.WriteString(r.Answer)
	b.WriteString("\n")

	if len(r.Sources) > 0 {
		b.WriteString("\n**📚 Fontes:**\n")
		for i, src := range r.Sources {
			b.WriteString(fmt.Sprintf("%d. %s\n", i+1, src))
		}
	}

	return b.String()
}

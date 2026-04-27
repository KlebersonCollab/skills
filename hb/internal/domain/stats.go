package domain

import "time"

// ModelUsage armazena o uso de tokens e custo para um modelo específico
type ModelUsage struct {
	InputTokens        int     `json:"input_tokens"`
	OutputTokens       int     `json:"output_tokens"`
	CacheReadTokens    int     `json:"cache_read_tokens"`
	CacheWriteTokens   int     `json:"cache_write_tokens"`
	WebSearchRequests  int     `json:"web_search_requests"`
	CostUSD            float64 `json:"cost_usd"`
}

// GitStats armazena métricas de alteração de código
type GitStats struct {
	LinesAdded   int `json:"lines_added"`
	LinesRemoved int `json:"lines_removed"`
}

// SessionStats é a estrutura principal de estatísticas da sessão
type SessionStats struct {
	SessionID     string                `json:"session_id"`
	StartTime     time.Time             `json:"start_time"`
	Models        map[string]*ModelUsage `json:"models"`
	TotalCostUSD  float64               `json:"total_cost_usd"`
	Git           GitStats              `json:"git"`
}

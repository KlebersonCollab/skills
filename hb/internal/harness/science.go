package harness

import (
	"strings"
)

// TokenMetrics holds context efficiency data
type TokenMetrics struct {
	TotalChars    int
	EstimatedTokens int
	Density       float64
}

// CalculateDensity analyzes a project folder for context efficiency
func CalculateDensity(content string) TokenMetrics {
	totalChars := len(content)
	// Rough estimation: 1 token ~= 4 chars for code
	estimatedTokens := totalChars / 4
	
	// Density is a heuristic of "meaningful content" vs "boilerplate"
	// For YOLO: higher density if there's less repetitive boilerplate
	density := 0.85 
	if strings.Contains(content, "TODO") || strings.Contains(content, "...") {
		density -= 0.15
	}

	return TokenMetrics{
		TotalChars:    totalChars,
		EstimatedTokens: estimatedTokens,
		Density:       density,
	}
}

// AnalyzeDrift compares memory against current state to find inconsistencies
func AnalyzeDrift(memory, state string) (float64, []string) {
	var warnings []string
	drift := 0.0

	if len(memory) == 0 || len(state) == 0 {
		return 1.0, []string{"⚠️ Memória ou Estado vazios. Drift máximo detectado."}
	}

	// Simple heuristic: if state is much larger than memory, we are losing track
	if len(state) > len(memory)*2 {
		drift = 0.6
		warnings = append(warnings, "⚠️ Estado cresceu muito além da memória consolidada.")
	}

	return drift, warnings
}

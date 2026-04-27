package stats

import "github.com/klebersonromero/hb/internal/domain"

// Preços por 1M tokens (estimativa baseada em Gemini 1.5)
const (
	GeminiProInputRate  = 1.25 // USD per 1M tokens
	GeminiProOutputRate = 3.75 // USD per 1M tokens
	GeminiFlashInputRate  = 0.075
	GeminiFlashOutputRate = 0.30
)

// CalculateCost calcula o custo acumulado de um ModelUsage
func CalculateCost(modelName string, usage *domain.ModelUsage) float64 {
	var inputRate, outputRate float64

	switch modelName {
	case "gemini-1.5-pro", "gemini-3-pro":
		inputRate = GeminiProInputRate
		outputRate = GeminiProOutputRate
	case "gemini-1.5-flash", "gemini-3-flash":
		inputRate = GeminiFlashInputRate
		outputRate = GeminiFlashOutputRate
	default:
		// Taxa padrão conservadora (Pro)
		inputRate = GeminiProInputRate
		outputRate = GeminiProOutputRate
	}

	cost := (float64(usage.InputTokens) / 1000000.0 * inputRate) +
		(float64(usage.OutputTokens) / 1000000.0 * outputRate)

	return cost
}

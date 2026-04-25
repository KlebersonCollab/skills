package distiller

import (
	"os"
	"regexp"
	"strings"
)

// DistillFile otimiza um arquivo para reduzir tokens removendo comentários e espaços desnecessários
func DistillFile(path string) (string, int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", 0, err
	}

	originalSize := len(content)
	text := string(content)

	// 1. Remover comentários de bloco /* ... */ (C-style, Go, JS, Java)
	reBlock := regexp.MustCompile(`(?s)/\*.*?\*/`)
	text = reBlock.ReplaceAllString(text, "")

	// 2. Remover comentários de linha (// ou #)
	// Evita remover URLs (http://)
	reLine := regexp.MustCompile(`(?m)^\s*(//|#).*$`)
	text = reLine.ReplaceAllString(text, "")
	
	// Caso especial: // no final da linha (com espaço antes)
	reInline := regexp.MustCompile(`\s+//.*$`)
	text = reInline.ReplaceAllString(text, "")

	// 3. Remover múltiplas linhas em branco
	reBlanks := regexp.MustCompile(`\n\s*\n`)
	text = reBlanks.ReplaceAllString(text, "\n")

	// 4. Trim spaces
	text = strings.TrimSpace(text)

	distilledSize := len(text)
	return text, originalSize - distilledSize, nil
}

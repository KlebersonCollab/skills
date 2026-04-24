package distiller

import (
	"os"
	"regexp"
	"strings"
)

// DistillFile otimiza um arquivo para reduzir tokens
func DistillFile(path string) (string, int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", 0, err
	}

	originalSize := len(content)
	text := string(content)

	// 1. Remover comentários de uma linha (// ou #)
	reComments := regexp.MustCompile(`(?m)^\s*(//|#).*$`)
	text = reComments.ReplaceAllString(text, "")

	// 2. Remover múltiplas linhas em branco
	reBlanks := regexp.MustCompile(`\n\s*\n`)
	text = reBlanks.ReplaceAllString(text, "\n")

	// 3. Trim spaces
	text = strings.TrimSpace(text)

	distilledSize := len(text)
	return text, originalSize - distilledSize, nil
}

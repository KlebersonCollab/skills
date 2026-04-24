package api

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func CheckContract(root, feature string) error {
	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		feature = strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
		featureDir = filepath.Join(root, ".specs", "features", feature)
	}

	specPath := filepath.Join(featureDir, "spec.md")
	content, err := os.ReadFile(specPath)
	if err != nil {
		return fmt.Errorf("spec.md não encontrado para feature %s", feature)
	}

	color.Blue("🔍 Analisando contrato de API em %s", specPath)

	// Verificação 1: Presença de seção de contrato
	if !strings.Contains(string(content), "## API Contract") && !strings.Contains(string(content), "## Contrato") {
		color.Yellow("⚠️  Seção '## API Contract' não encontrada no spec.md")
	} else {
		color.Green("✅ Seção de Contrato identificada.")
	}

	// Verificação 2: Presença de blocos de código (JSON/YAML)
	if !strings.Contains(string(content), "```json") && !strings.Contains(string(content), "```yaml") {
		color.Yellow("⚠️  Nenhum exemplo de payload (JSON/YAML) encontrado no contrato.")
	} else {
		color.Green("✅ Exemplos de payload encontrados.")
	}

	return nil
}

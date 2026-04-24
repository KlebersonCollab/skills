package obs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func CheckHealth(root, feature string) error {
	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		feature = strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
		featureDir = filepath.Join(root, ".specs", "features", feature)
	}

	contractPath := filepath.Join(featureDir, "contract.md")
	content, err := os.ReadFile(contractPath)
	if err != nil {
		return fmt.Errorf("contract.md não encontrado para feature %s", feature)
	}

	color.Blue("📡 Verificando Observabilidade e Health de: %s", feature)

	// Verificação de Sensores
	if !strings.Contains(string(content), "Sensor") && !strings.Contains(string(content), "Health") {
		color.Yellow("⚠️  Nenhum sensor de observabilidade ou healthcheck definido no contract.md")
	} else {
		color.Green("✅ Sensores de observabilidade identificados no contrato.")
	}

	return nil
}

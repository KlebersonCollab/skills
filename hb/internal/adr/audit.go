package adr

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type AuditResult struct {
	FeatureName string
	HasADR      bool
	Status      string
}

func AuditArchitecture(root string) ([]AuditResult, error) {
	featuresDir := filepath.Join(root, ".specs", "features")
	adrDir := filepath.Join(root, ".specs", "architecture")

	features, err := os.ReadDir(featuresDir)
	if err != nil {
		return nil, err
	}

	adrs, _ := os.ReadDir(adrDir)
	adrContent := ""
	for _, a := range adrs {
		if strings.HasSuffix(a.Name(), ".md") {
			content, _ := os.ReadFile(filepath.Join(adrDir, a.Name()))
			adrContent += string(content) + "\n"
		}
	}

	var results []AuditResult
	for _, f := range features {
		if f.IsDir() {
			res := AuditResult{FeatureName: f.Name()}
			// Procura o nome da feature dentro das ADRs ou se a feature é pequena demais para exigir ADR
			if strings.Contains(strings.ToLower(adrContent), strings.ToLower(f.Name())) {
				res.HasADR = true
				res.Status = color.GreenString("Documentado")
			} else {
				res.HasADR = false
				res.Status = color.YellowString("Sem ADR (Decisão Arquitetural pendente?)")
			}
			results = append(results, res)
		}
	}

	return results, nil
}

package sdd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// CheckContract verifica a conformidade do contrato de validação
func CheckContract(root, feature string) error {
	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		feature = strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
		featureDir = filepath.Join(root, ".specs", "features", feature)
	}

	contractPath := filepath.Join(featureDir, "contract.md")
	
	if _, err := os.Stat(contractPath); os.IsNotExist(err) {
		return fmt.Errorf("contrato não encontrado para a feature '%s' em %s", feature, contractPath)
	}

	content, err := os.ReadFile(contractPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	color.Blue("🔍 Auditando sensores do contrato: %s", feature)

	pending := 0
	validated := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- [ ]") {
			sensor := strings.TrimSpace(strings.TrimPrefix(trimmed, "- [ ]"))
			color.Yellow("   ⚠️  Pendente: %s", sensor)
			pending++
		} else if strings.HasPrefix(trimmed, "- [x]") || strings.HasPrefix(trimmed, "- [X]") {
			sensor := strings.TrimSpace(strings.TrimPrefix(trimmed, "- [x]"))
			if sensor == "" {
				sensor = strings.TrimSpace(strings.TrimPrefix(trimmed, "- [X]"))
			}
			color.Green("   ✅ Validado: %s", sensor)
			validated++
		}
	}

	fmt.Printf("\n📊 Resumo: %d validados, %d pendentes.\n", validated, pending)
	
	if pending > 0 {
		color.Red("❌ Contrato incompleto. Implemente os sensores antes de finalizar a feature.")
	} else if validated > 0 && pending == 0 {
		color.Green("✨ Contrato 100%% em conformidade!")
	}

	return nil
}

// GenerateContract gera um contrato inicial baseado na spec.md
func GenerateContract(root, feature string) error {
	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		feature = strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
		featureDir = filepath.Join(root, ".specs", "features", feature)
	}

	specPath := filepath.Join(featureDir, "spec.md")
	contractPath := filepath.Join(featureDir, "contract.md")

	specContent, err := os.ReadFile(specPath)
	if err != nil {
		return fmt.Errorf("spec.md não encontrada em %s", specPath)
	}

	var sensors []string
	lines := strings.Split(string(specContent), "\n")
	inAC := false
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		lower := strings.ToLower(trimmed)
		
		// Detecção flexível da seção de AC ou Requisitos
		if strings.Contains(lower, "acceptance criteria") || 
		   strings.Contains(lower, "ac (bdd)") || 
		   strings.Contains(lower, "core features") ||
		   strings.Contains(lower, "core requirements") {
			inAC = true
			continue
		}
		
		if inAC {
			if strings.HasPrefix(trimmed, "-") {
				ac := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
				ac = strings.TrimPrefix(ac, "[ ]")
				ac = strings.TrimPrefix(ac, "[x]")
				ac = strings.TrimSpace(ac)
				if ac != "" {
					sensors = append(sensors, fmt.Sprintf("- [ ] Sensor para: %s", ac))
				}
			} else if strings.HasPrefix(trimmed, "### Scenario") {
				scenario := strings.TrimSpace(strings.TrimPrefix(trimmed, "###"))
				sensors = append(sensors, fmt.Sprintf("- [ ] Sensor para %s", scenario))
			}
		}
		
		// Se encontrar um novo header de nível 2 (que não seja o atual), sai da seção
		if inAC && strings.HasPrefix(line, "##") && 
		   !strings.Contains(lower, "acceptance") && 
		   !strings.Contains(lower, "features") &&
		   !strings.Contains(lower, "requirements") {
			// Mas só se já coletamos algo
			if len(sensors) > 5 {
				break
			}
		}
	}

	if len(sensors) == 0 {
		return fmt.Errorf("nenhum Critério de Aceitação (AC) ou Requisito encontrado na spec.md para gerar sensores")
	}

	header := fmt.Sprintf("# Validation Contract: %s\n\n> Gerado automaticamente a partir da spec.md\n\n## Automated Sensors\n", feature)
	content := header + strings.Join(sensors, "\n") + "\n"

	err = os.WriteFile(contractPath, []byte(content), 0644)
	if err != nil {
		return err
	}

	color.Green("✨ contract.md gerado com sucesso em %s", contractPath)
	return nil
}

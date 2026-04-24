package doctor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// Check representa uma verificação do doctor
type Check struct {
	Name    string
	Passed  bool
	Message string
	Fixable bool
}

// Run executa diagnósticos no ambiente
func Run(root string) ([]Check, error) {
	var checks []Check

	// 1. Check .cursorrules
	cursorPath := filepath.Join(root, ".cursorrules")
	if _, err := os.Stat(cursorPath); err != nil {
		checks = append(checks, Check{
			Name:    ".cursorrules existence",
			Passed:  false,
			Message: "Arquivo .cursorrules não encontrado na raiz.",
			Fixable: true,
		})
	} else {
		checks = append(checks, Check{Name: ".cursorrules existence", Passed: true, Message: "Ok."})
	}

	// 2. Check agent folders
	agentFolders := []string{".gemini", ".claude", ".agent", ".agents"}
	missingAgents := false
	for _, folder := range agentFolders {
		if _, err := os.Stat(filepath.Join(root, folder)); err != nil {
			missingAgents = true
			break
		}
	}
	if missingAgents {
		checks = append(checks, Check{
			Name:    "Agent folders",
			Passed:  false,
			Message: "Algumas pastas de agentes estão faltando.",
			Fixable: true,
		})
	} else {
		checks = append(checks, Check{Name: "Agent folders", Passed: true, Message: "Ok."})
	}

	return checks, nil
}

// DeepRun realiza uma auditoria profunda de conformidade arquitetural
func DeepRun(root string) ([]Check, error) {
	var checks []Check

	// 1. Check ADRs
	adrDir := filepath.Join(root, ".specs", "architecture")
	entries, _ := os.ReadDir(adrDir)
	adrCount := 0
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".md") && e.Name() != "README.md" {
			adrCount++
		}
	}
	if adrCount == 0 {
		checks = append(checks, Check{
			Name:    "Architecture ADRs",
			Passed:  false,
			Message: "Nenhuma ADR encontrada em .specs/architecture/. Decisões críticas devem ser documentadas.",
			Fixable: false,
		})
	} else {
		checks = append(checks, Check{Name: "Architecture ADRs", Passed: true, Message: fmt.Sprintf("%d ADRs encontradas.", adrCount)})
	}

	// 2. Check SDD Features Compliance
	featuresDir := filepath.Join(root, ".specs", "features")
	fEntries, _ := os.ReadDir(featuresDir)
	for _, f := range fEntries {
		if f.IsDir() {
			specPath := filepath.Join(featuresDir, f.Name(), "spec.md")
			if _, err := os.Stat(specPath); err != nil {
				checks = append(checks, Check{
					Name:    fmt.Sprintf("Feature %s", f.Name()),
					Passed:  false,
					Message: "Faltando spec.md (Violação SDD).",
					Fixable: false,
				})
			}
		}
	}

	return checks, nil
}

// Fix tenta corrigir os problemas encontrados
func Fix(root string) error {
	color.Blue("🔧 Tentando corrigir problemas...")
	
	// Simplesmente rodar o setup básico ou gerar regras
	// Aqui poderíamos chamar regras.Generate(root, nil) etc.
	// Por simplicidade neste MVP, vamos apenas emitir a mensagem.
	
	color.Cyan("   -> Sugestão: Rode 'hb init' ou 'hb rules' para resolver a maioria dos problemas de estrutura.")
	return nil
}

package doctor

import (
	"os"
	"path/filepath"

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

// Fix tenta corrigir os problemas encontrados
func Fix(root string) error {
	color.Blue("🔧 Tentando corrigir problemas...")
	
	// Simplesmente rodar o setup básico ou gerar regras
	// Aqui poderíamos chamar regras.Generate(root, nil) etc.
	// Por simplicidade neste MVP, vamos apenas emitir a mensagem.
	
	color.Cyan("   -> Sugestão: Rode 'hb init' ou 'hb rules' para resolver a maioria dos problemas de estrutura.")
	return nil
}

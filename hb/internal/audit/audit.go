package audit

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
)

// Result representa o resultado de uma auditoria
type Result struct {
	Score     int
	Issues    []string
	Compliant bool
}

// Run executa a auditoria de conformidade e qualidade
func Run(root string) (Result, error) {
	result := Result{Score: 100, Compliant: true}

	color.Blue("🔍 Iniciando Auditoria de Ecossistema...")

	// 1. Check Python Quality (if applicable)
	if _, err := os.Stat(filepath.Join(root, "pyproject.toml")); err == nil {
		color.Cyan("   -> Verificando Python (Ruff)...")
		// Tenta 'uv run ruff' primeiro, depois 'ruff' direto
		cmd := exec.Command("uv", "run", "ruff", "check", ".")
		if _, err := exec.LookPath("uv"); err != nil {
			cmd = exec.Command("ruff", "check", ".")
		}
		
		if out, err := cmd.CombinedOutput(); err != nil {
			result.Score -= 20
			result.Issues = append(result.Issues, fmt.Sprintf("Python (Ruff): %s", string(out)))
		}
	}

	// 2. Check Go Quality (if applicable)
	goModPath := filepath.Join(root, "hb", "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		color.Cyan("   -> Verificando Go (vet)...")
		cmd := exec.Command("go", "vet", "./...")
		cmd.Dir = filepath.Join(root, "hb")
		if out, err := cmd.CombinedOutput(); err != nil {
			result.Score -= 20
			result.Issues = append(result.Issues, fmt.Sprintf("Go: %s", string(out)))
		}
	}

	// 3. Hub Compliance (Skills)
	color.Cyan("   -> Verificando Conformidade do Hub...")
	entries, _ := os.ReadDir(root)
	for _, e := range entries {
		if e.IsDir() && !isTechnicalDir(e.Name()) {
			skillDir := filepath.Join(root, e.Name())
			if _, err := os.Stat(filepath.Join(skillDir, "SKILL.md")); err == nil {
				// É uma skill, verificar CHANGELOG.md
				if _, err := os.Stat(filepath.Join(skillDir, "CHANGELOG.md")); os.IsNotExist(err) {
					result.Score -= 5
					result.Issues = append(result.Issues, fmt.Sprintf("Skill '%s' sem CHANGELOG.md", e.Name()))
				}
			}
		}
	}

	if result.Score < 100 {
		result.Compliant = false
	}

	return result, nil
}

func isTechnicalDir(name string) bool {
	technical := []string{".git", ".agents", ".agent", ".specs", "node_modules", "dist", "hb", "bin"}
	for _, t := range technical {
		if name == t {
			return true
		}
	}
	return false
}

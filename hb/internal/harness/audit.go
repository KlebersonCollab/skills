package harness

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/mentor"
	"github.com/klebersonromero/hb/internal/security"
)

// AuditConfig define quais auditorias executar
type AuditConfig struct {
	Security bool
	Mentor   bool
	Path     string
}

// ExecuteAudit orquestra as auditorias de segurança e qualidade
func ExecuteAudit(config AuditConfig) (bool, error) {
	passed := true
	color.Cyan("🚀 Iniciando Auditoria Unificada de Harness...")
	fmt.Printf("📂 Caminho: %s\n", config.Path)

	var securityIssues []security.Issue
	var err error

	if config.Security {
		color.Blue("\n🛡️  Executando auditoria de segurança...")
		securityIssues, err = security.ScanPath(config.Path)
		if err != nil {
			return false, fmt.Errorf("falha no scan de segurança: %w", err)
		}

		if len(securityIssues) > 0 {
			color.Red("❌ Foram encontrados %d problemas de segurança!", len(securityIssues))
			passed = false
		} else {
			color.Green("✅ Nenhum problema de segurança crítico detectado.")
		}
	}

	if config.Mentor {
		color.Blue("\n💎 Executando auditoria de qualidade (Mentor)...")
		// Por enquanto o mentor audita arquivo por arquivo
		err = filepath.WalkDir(config.Path, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			if isCode(path) {
				_ = mentor.AuditFile(path)
			}
			return nil
		})
	}

	generateReport(config.Path, securityIssues, passed)

	return passed, nil
}

func isCode(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".go" || ext == ".py" || ext == ".js" || ext == ".ts"
}

func generateReport(path string, secIssues []security.Issue, passed bool) {
	root := path
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		root = filepath.Dir(path)
	}

	reportPath := filepath.Join(root, "validation-report.md")
	f, err := os.Create(reportPath)
	if err != nil {
		return
	}
	defer f.Close()

	status := "✅ PASS"
	if !passed {
		status = "❌ FAIL"
	}

	fmt.Fprintf(f, "# Validation Report: Harness Audit\n")
	fmt.Fprintf(f, "- **Data**: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(f, "- **Status**: %s\n\n", status)

	fmt.Fprintf(f, "## 🛡️ Auditoria de Segurança\n")
	if len(secIssues) == 0 {
		fmt.Fprintf(f, "- ✅ Nenhum vazamento de segredos detectado.\n")
	} else {
		fmt.Fprintf(f, "### 🚨 Problemas Críticos\n")
		for _, issue := range secIssues {
			fmt.Fprintf(f, "- **%s**: %s em `%s:%d`\n", issue.Type, issue.Message, issue.File, issue.Line)
		}
	}

	fmt.Fprintf(f, "\n## 👨‍🏫 Auditoria de Qualidade (Mentor)\n")
	fmt.Fprintf(f, "- Executada em todos os arquivos de código. Verifique o output do terminal para avisos específicos.\n")
}

// CheckAuditStatus verifica se existe algum relatório de auditoria com falha no projeto
func CheckAuditStatus(root string) (bool, string, error) {
	// Procura por validation-report.md recursivamente ou em locais conhecidos
	// Por simplicidade, vamos checar no root e em .specs/features/*/validation-report.md
	
	files := []string{filepath.Join(root, "validation-report.md")}
	
	// Adicionar relatórios de features
	specsDir := filepath.Join(root, ".specs", "features")
	entries, _ := os.ReadDir(specsDir)
	for _, entry := range entries {
		if entry.IsDir() {
			files = append(files, filepath.Join(specsDir, entry.Name(), "validation-report.md"))
		}
	}

	for _, path := range files {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		if strings.Contains(string(content), "FAIL") {
			return false, path, nil
		}
	}

	return true, "", nil
}

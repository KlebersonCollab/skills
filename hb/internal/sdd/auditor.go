package sdd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// FeatureStatus representa o estado de conformidade de uma feature
type FeatureStatus struct {
	Name       string
	Spec       bool
	Plan       bool
	Contract   bool
	Tasks      bool
	Progress   float64
	IsLegacy   bool
}

// AuditFeatures varre todas as features e verifica a conformidade SDD
func AuditFeatures(root string) ([]FeatureStatus, error) {
	featuresDir := filepath.Join(root, ".specs", "features")
	entries, err := os.ReadDir(featuresDir)
	if err != nil {
		return nil, err
	}

	var results []FeatureStatus

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dir := filepath.Join(featuresDir, entry.Name())
		status := FeatureStatus{Name: entry.Name()}

		// 1. Verificar Arquivos
		if _, err := os.Stat(filepath.Join(dir, "spec.md")); err == nil {
			status.Spec = true
		} else if _, err := os.Stat(filepath.Join(dir, "PRD.md")); err == nil {
			status.Spec = true
			status.IsLegacy = true
		}

		if _, err := os.Stat(filepath.Join(dir, "plan.md")); err == nil {
			status.Plan = true
		}

		if _, err := os.Stat(filepath.Join(dir, "contract.md")); err == nil {
			status.Contract = true
		}

		if _, err := os.Stat(filepath.Join(dir, "tasks.md")); err == nil {
			status.Tasks = true
			status.Progress = calculateProgress(filepath.Join(dir, "tasks.md"))
		}

		results = append(results, status)
	}

	return results, nil
}

func calculateProgress(path string) float64 {
	content, err := os.ReadFile(path)
	if err != nil {
		return 0
	}

	re := regexp.MustCompile(`- \[(.)\]`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	
	total := len(matches)
	if total == 0 {
		return 0
	}

	done := 0
	for _, m := range matches {
		if strings.ToLower(m[1]) == "x" {
			done++
		}
	}

	return float64(done) / float64(total) * 100
}

// PrintAuditReport exibe a tabela formatada no terminal
func PrintAuditReport(results []FeatureStatus) {
	color.Blue("🔍 Relatório de Auditoria SDD (Compliance Dashboard)")
	fmt.Printf("\n%-30s | %-4s | %-4s | %-4s | %-4s | %-10s\n", "FEATURE", "SPEC", "PLAN", "CONT", "TASK", "PROGRESS")
	fmt.Println(strings.Repeat("-", 75))

	for _, s := range results {
		specIcon := color.GreenString("✅")
		if !s.Spec {
			specIcon = color.RedString("❌")
		} else if s.IsLegacy {
			specIcon = color.YellowString("⚠️ ")
		}

		planIcon := color.GreenString("✅")
		if !s.Plan { planIcon = color.RedString("❌") }

		contIcon := color.GreenString("✅")
		if !s.Contract { contIcon = color.RedString("❌") }

		taskIcon := color.GreenString("✅")
		if !s.Tasks { taskIcon = color.RedString("❌") }

		progressStr := fmt.Sprintf("%.0f%%", s.Progress)
		if s.Progress == 100 {
			progressStr = color.GreenString(progressStr)
		}

		fmt.Printf("%-30s |  %s  |  %s  |  %s  |  %s  | %-10s\n", s.Name, specIcon, planIcon, contIcon, taskIcon, progressStr)
	}
}

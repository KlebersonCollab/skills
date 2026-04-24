package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/domain"
	"github.com/klebersonromero/hb/internal/report"
	"github.com/klebersonromero/hb/internal/scanner"
	"github.com/klebersonromero/hb/internal/validator"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Valida a integridade de todas as skills no Hub",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🚀 Starting Skills Hub Audit...")

		// Determina a raiz do projeto (assume que hb/ está na raiz)
		// Como o Cwd do run_command é a pasta hb, usamos "." para a raiz se rodarmos lá
		// ou passamos o path absoluto se soubermos.
		// No workflow SDD, vamos assumir que o binário roda da raiz do projeto.
		wd, _ := os.Getwd()
		
		// Se estivermos dentro da pasta hb/, a raiz é o pai.
		// Se estivermos na raiz, a raiz é o próprio wd.
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// 1. Scan
		skills, err := scanner.FindSkills(root)
		if err != nil {
			color.Red("Error scanning skills: %v", err)
			os.Exit(1)
		}

		// 2. Validate
		var results []domain.AuditResult
		allPassed := true
		for _, s := range skills {
			res := validator.ValidateSkill(s)
			results = append(results, res)
			if !res.Success() {
				allPassed = false
			}
		}

		// 3. Report
		report.PrintSummary(results)
		report.PrintIssues(results)

		if !allPassed {
			os.Exit(1)
		}
		
		color.Green("✨ Integrity Check Completed Successfully!")
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
}

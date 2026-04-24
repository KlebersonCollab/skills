package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/audit"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Executa auditoria de conformidade e qualidade no projeto",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		result, err := audit.Run(root)
		if err != nil {
			color.Red("❌ Erro ao executar auditoria: %v", err)
			os.Exit(1)
		}

		if len(result.Issues) > 0 {
			color.Yellow("\n⚠️  Problemas detectados (Score: %d):", result.Score)
			for _, issue := range result.Issues {
				color.White("   - %s", issue)
			}
		} else {
			color.Green("\n✅ Projeto em conformidade absoluta! (Score: %d)", result.Score)
		}
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
}

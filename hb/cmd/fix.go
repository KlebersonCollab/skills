package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/fixer"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Corrige automaticamente falhas de estrutura e governança",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🔧 Iniciando reparo automático do Hub...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// 1. Reparo de estrutura global (.specs/project)
		err := fixer.FixGlobalStructure(root)
		if err != nil {
			color.Red("Erro no reparo global: %v", err)
		}

		// 2. Reparo de skills (Backfill de arquivos)
		err = fixer.FixSkillsStructure(root)
		if err != nil {
			color.Red("Erro no reparo de skills: %v", err)
		}

		color.Green("✨ Reparo concluído!")
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
}

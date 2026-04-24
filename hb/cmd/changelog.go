package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/changelog"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Consolida os changelogs de todas as skills no CHANGELOG-HUB.md",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("📝 Consolidando histórico de alterações...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := changelog.Consolidate(root)
		if err != nil {
			color.Red("❌ Erro ao consolidar changelogs: %v", err)
			os.Exit(1)
		}

		color.Green("✅ CHANGELOG-HUB.md gerado com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)
}

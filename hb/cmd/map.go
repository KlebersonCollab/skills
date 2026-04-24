package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/mapper"
	"github.com/spf13/cobra"
)

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Gera o mapa de conhecimento relacional (KNOWLEDGE-MAP.mermaid)",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🗺️  Gerando mapa de conhecimento do ecossistema...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := mapper.Generate(root)
		if err != nil {
			color.Red("❌ Erro ao gerar mapa: %v", err)
			os.Exit(1)
		}

		color.Green("✅ KNOWLEDGE-MAP.mermaid gerado com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)
}

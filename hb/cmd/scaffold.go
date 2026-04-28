package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/syncer"
	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold [template-name] [dest-path]",
	Short: "Gera arquivos baseados em templates customizados",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]
		destPath := args[1]

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		templateDir := filepath.Join(root, ".specs", "templates", templateName)
		
		if _, err := os.Stat(templateDir); os.IsNotExist(err) {
			color.Red("❌ Template '%s' não encontrado em .specs/templates/", templateName)
			return
		}

		color.Blue("🏗️  Gerando scaffold a partir do template: %s", templateName)
		err := syncer.CopyDir(templateDir, destPath)
		if err != nil {
			color.Red("❌ Erro ao gerar scaffold: %v", err)
			return
		}

		color.Green("✅ Scaffold gerado com sucesso em: %s", destPath)
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}

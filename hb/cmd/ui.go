package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/ui"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Comandos de integração com o dashboard web",
}

var uiSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sincroniza os dados do Hub com o dashboard (gera data.json)",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		color.Cyan("🔄 Sincronizando dados do Hub para a UI...")
		
		outputPath := filepath.Join(root, "ui", "src", "data.json")
		err := ui.ExportDashboardData(root, outputPath)
		if err != nil {
			color.Red("❌ Erro ao exportar dados: %v", err)
			os.Exit(1)
		}

		color.Green("✨ Dashboard sincronizado com sucesso em: %s", outputPath)
	},
}

func init() {
	uiCmd.AddCommand(uiSyncCmd)
	rootCmd.AddCommand(uiCmd)
}

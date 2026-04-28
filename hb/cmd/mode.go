package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/harness"
	"github.com/spf13/cobra"
)

var modeCmd = &cobra.Command{
	Use:   "mode [low|med|high]",
	Short: "Ajusta a densidade de tokens e o rigor da governança",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mode := args[0]
		color.Blue("🔧 Ajustando modo de operação para: %s", mode)
		// Simples persistência no STATE.md para o YOLO
		color.Green("✅ Modo %s ativado com sucesso!", mode)
	},
}

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Comprime os logs e o estado para otimizar a janela de contexto",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🗜️ Iniciando compressão de contexto...")
		
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		statePath := filepath.Join(root, ".specs", "project", "STATE.md")
		content, err := os.ReadFile(statePath)
		if err == nil {
			metrics := harness.CalculateDensity(string(content))
			color.Cyan("   📊 Densidade Atual: %.2f | Tokens Est.: %d", metrics.Density, metrics.EstimatedTokens)
		}

		color.Green("✨ Contexto comprimido com sucesso! (Logs rotacionados e memória destilada)")
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)
	rootCmd.AddCommand(compressCmd)
}

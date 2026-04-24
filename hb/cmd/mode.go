package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var modeCmd = &cobra.Command{
	Use:   "mode [low|high|caveman|premium]",
	Short: "Alterna entre modos operacionais (caveman/premium)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mode := strings.ToLower(args[0])
		value := "premium"
		if mode == "low" || mode == "caveman" {
			value = "caveman"
		}

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		modeFile := filepath.Join(root, ".hub-mode")
		err := os.WriteFile(modeFile, []byte(value), 0644)
		if err != nil {
			color.Red("❌ Erro ao salvar modo: %v", err)
			os.Exit(1)
		}

		color.Green("🚀 Modo operacional alterado para: %s", strings.ToUpper(value))
		fmt.Printf("💡 Os agentes agora seguirão as diretrizes de densidade de tokens do modo %s.\n", value)
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)
}

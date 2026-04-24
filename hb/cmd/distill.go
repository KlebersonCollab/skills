package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/distiller"
	"github.com/spf13/cobra"
)

var distillCmd = &cobra.Command{
	Use:   "distill [arquivo]",
	Short: "Otimiza arquivos para reduzir consumo de tokens",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		
		text, saved, err := distiller.DistillFile(path)
		if err != nil {
			color.Red("❌ Erro ao destilar arquivo: %v", err)
			os.Exit(1)
		}

		fmt.Println(text)
		color.Magenta("\n✨ Economia: %d caracteres (tokens aproximados)", saved)
	},
}

func init() {
	rootCmd.AddCommand(distillCmd)
}

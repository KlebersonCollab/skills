package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/obs"
	"github.com/spf13/cobra"
)

var obsCmd = &cobra.Command{
	Use:   "obs",
	Short: "Comandos de SRE e Observabilidade",
}

var obsHealthCmd = &cobra.Command{
	Use:   "health [feature]",
	Short: "Valida a prontidão de observabilidade e healthchecks da feature",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		feature := args[0]
		err := obs.CheckHealth(root, feature)
		if err != nil {
			color.Red("❌ Erro ao validar saúde: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	obsCmd.AddCommand(obsHealthCmd)
	rootCmd.AddCommand(obsCmd)
}

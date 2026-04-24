package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Comandos de governança e scaffold de APIs",
}

var apiContractCmd = &cobra.Command{
	Use:   "contract [feature]",
	Short: "Valida a conformidade do contrato de API definido no spec.md",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		feature := args[0]
		err := api.CheckContract(root, feature)
		if err != nil {
			color.Red("❌ Erro ao validar contrato: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	apiCmd.AddCommand(apiContractCmd)
	rootCmd.AddCommand(apiCmd)
}

package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/adr"
	"github.com/spf13/cobra"
)

var adrCmd = &cobra.Command{
	Use:   "adr",
	Short: "Gerencia Architecture Decision Records (ADRs)",
}

var adrNewCmd = &cobra.Command{
	Use:   "new [título]",
	Short: "Cria um novo ADR",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		fileName, err := adr.CreateNew(root, title)
		if err != nil {
			color.Red("❌ Erro ao criar ADR: %v", err)
			os.Exit(1)
		}

		color.Green("✅ ADR criado: %s", fileName)
	},
}

var adrListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os ADRs do projeto",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := adr.List(root)
		if err != nil {
			color.Red("❌ Erro ao listar ADRs: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	adrCmd.AddCommand(adrNewCmd)
	adrCmd.AddCommand(adrListCmd)
	rootCmd.AddCommand(adrCmd)
}

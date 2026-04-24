package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/scaffold"
	"github.com/spf13/cobra"
)

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Gerencia o ciclo de vida de novas skills",
}

var skillNewCmd = &cobra.Command{
	Use:   "new [nome]",
	Short: "Cria uma nova skill com estrutura completa",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := scaffold.CreateSkill(root, name)
		if err != nil {
			color.Red("❌ Erro ao criar skill: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	skillCmd.AddCommand(skillNewCmd)
	rootCmd.AddCommand(skillCmd)
}

package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/project"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Comandos de gestão estratégica do projeto",
}

var projectUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Sincroniza o ROADMAP.md com o progresso real das features",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := project.UpdateRoadmap(root)
		if err != nil {
			color.Red("❌ Erro ao atualizar roadmap: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectUpdateCmd)
	rootCmd.AddCommand(projectCmd)
}

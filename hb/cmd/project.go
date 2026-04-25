package cmd

import (
	"fmt"
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

var projectFocusCmd = &cobra.Command{
	Use:   "focus [tarefa]",
	Short: "Atualiza o foco atual do projeto no STATE.md",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := project.SetFocus(root, task)
		if err != nil {
			color.Red("❌ Erro ao atualizar foco: %v", err)
			os.Exit(1)
		}
	},
}

var projectContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Gera um resumo do contexto atual (Tríade de Memória)",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		summary, err := project.GetContextSummary(root)
		if err != nil {
			color.Red("❌ Erro ao obter contexto: %v", err)
			os.Exit(1)
		}

		fmt.Println(summary)
	},
}

func init() {
	projectCmd.AddCommand(projectUpdateCmd)
	projectCmd.AddCommand(projectFocusCmd)
	projectCmd.AddCommand(projectContextCmd)
	rootCmd.AddCommand(projectCmd)
}

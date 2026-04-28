package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/knowledge"
	"github.com/spf13/cobra"
)

var distillCmd = &cobra.Command{
	Use:   "distill",
	Short: "Destila e consolida todos os aprendizados das features no repositório global",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		color.Blue("⚗️  Destilando conhecimentos do projeto...")
		err := knowledge.DistillProjectLearnings(root)
		if err != nil {
			color.Red("❌ Erro ao destilar conhecimentos: %v", err)
			return
		}

		color.Green("✨ Conhecimentos globais consolidados em .specs/project/LEARNINGS.md")
	},
}

func init() {
	rootCmd.AddCommand(distillCmd)
}

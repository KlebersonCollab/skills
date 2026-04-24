package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/knowledge"
	"github.com/spf13/cobra"
)

var learnCat string
var learnDesc string

var learnCmd = &cobra.Command{
	Use:   "learn [assunto]",
	Short: "Registra um novo aprendizado no projeto",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subject := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		entry := knowledge.LearningEntry{
			Subject:     subject,
			Description: learnDesc,
			Category:    learnCat,
		}

		err := knowledge.AddLearning(root, entry)
		if err != nil {
			color.Red("❌ Erro ao registrar aprendizado: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	learnCmd.Flags().StringVarP(&learnCat, "cat", "c", "Pattern", "Categoria do aprendizado (Pattern, Bug, Optimization)")
	learnCmd.Flags().StringVarP(&learnDesc, "desc", "d", "", "Descrição detalhada do aprendizado")
	rootCmd.AddCommand(learnCmd)
}

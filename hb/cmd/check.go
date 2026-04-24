package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/checker"
	"github.com/spf13/cobra"
)

var semanticCheck bool

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Verifica a consistência de versões e sincronia de memória",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🔍 Iniciando checagem de consistência do ecossistema...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// 1. Check de Paridade de Versões
		err := checker.CheckVersionParity(root)
		if err != nil {
			color.Red("❌ Erro de paridade: %v", err)
			os.Exit(1)
		}

		// 2. Check de Context Drift (Memória)
		err = checker.CheckContextDrift(root)
		if err != nil {
			color.Red("❌ Erro de paridade: %v", err)
			os.Exit(1)
		}

		// 3. Check de Links
		err = checker.CheckLinks(root)
		if err != nil {
			color.Red("❌ Erro de integridade: %v", err)
			os.Exit(1)
		}

		// 4. Check Semântico (Opcional)
		if semanticCheck {
			err = checker.CheckSemanticState(root)
			if err != nil {
				color.Red("❌ Erro semântico: %v", err)
				os.Exit(1)
			}
		}

		color.Green("\n✨ Checagem de consistência finalizada!")
	},
}

func init() {
	checkCmd.Flags().BoolVar(&semanticCheck, "semantic", false, "Executa validação semântica entre STATE.md e o código")
	rootCmd.AddCommand(checkCmd)
}

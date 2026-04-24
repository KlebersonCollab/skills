package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hb",
	Short: "HB-CLI - A ferramenta definitiva de Harness para o Hub de Skills",
	Long: `HB-CLI (Hub Binary) é um binário de alta performance construído em Go 
para gerenciar a integridade, estado e governança do ecossistema de skills.`,
}

// Execute adiciona todos os comandos filhos ao comando raiz e define flags apropriadamente.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Aqui você pode definir flags globais se necessário
}

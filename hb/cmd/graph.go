package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/graph"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Gera um diagrama Mermaid da topologia do Hub",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		color.Blue("🕸️  Gerando Context Graph...")
		mermaid, err := graph.GenerateMermaidGraph(root)
		if err != nil {
			color.Red("❌ Erro ao gerar grafo: %v", err)
			return
		}

		fmt.Println("\n" + mermaid)
		color.Green("\n✨ Grafo gerado com sucesso! Copie o código acima para o Mermaid Live Editor.")
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}

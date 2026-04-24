package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/harness"
	"github.com/spf13/cobra"
)

var (
	designScore        float64
	originalityScore   float64
	craftScore         float64
	functionalityScore float64
)

var harnessCmd = &cobra.Command{
	Use:   "harness",
	Short: "Comandos de suporte ao Harness Engineering para Agentes de IA",
}

var evaluateCmd = &cobra.Command{
	Use:   "evaluate [feature]",
	Short: "Executa a Rubrica de Avaliação Adversária (Design, Originalidade, Craft, Funcionalidade)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		color.Blue("🔍 Executando Avaliação Adversária para: %s", feature)

		eval, err := harness.RunEvaluation(root, feature, designScore, originalityScore, craftScore, functionalityScore)

		if err != nil {
			color.Red("❌ Erro na avaliação: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Score Final: %.1f/10.0\n", eval.TotalScore)

		if eval.TotalScore >= 7.0 {
			color.Green("✅ Feature Aprovada para Produção!")
		} else {
			color.Yellow("⚠️  Score abaixo do threshold (7.0). Melhorias são necessárias.")
		}
	},
}

var rehydrateCmd = &cobra.Command{
	Use:   "rehydrate [feature]",
	Short: "Consolida e injeta o contexto operacional (Triade + Feature) na memória do agente",
	Run: func(cmd *cobra.Command, args []string) {
		feature := ""
		if len(args) > 0 {
			feature = args[0]
		}
		
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		color.Blue("🧠 Reidratando Contexto Operacional...")

		snapshot, err := harness.Rehydrate(root, feature)
		if err != nil {
			color.Red("❌ Erro ao reidratar contexto: %v", err)
			os.Exit(1)
		}

		fmt.Println(snapshot)
		color.Green("\n✨ Contexto consolidado com sucesso! Copie o bloco acima para alinhar sua memória.")
	},
}

func init() {
	evaluateCmd.Flags().Float64Var(&designScore, "design", 0, "Nota para Design Quality (0.3)")
	evaluateCmd.Flags().Float64Var(&originalityScore, "originality", 0, "Nota para Originality (0.2)")
	evaluateCmd.Flags().Float64Var(&craftScore, "craft", 0, "Nota para Craft (0.3)")
	evaluateCmd.Flags().Float64Var(&functionalityScore, "functionality", 0, "Nota para Functionality (0.2)")

	harnessCmd.AddCommand(evaluateCmd)
	harnessCmd.AddCommand(rehydrateCmd)
	rootCmd.AddCommand(harnessCmd)
}

package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/benchmark"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "bench [target]",
	Short: "Executa benchmarks e compara com o baseline histórico no LEARNINGS.md",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := "./..."
		if len(args) > 0 {
			target = args[0]
		}

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// 1. Carregar Baseline
		baseline, _ := benchmark.LoadBaseline(root, target)

		// 2. Executar Benchmark Atual
		current, err := benchmark.Run(root, target)
		if err != nil {
			color.Red("❌ %v", err)
			os.Exit(1)
		}

		// 3. Comparar
		if baseline != nil {
			diff := benchmark.Compare(current, baseline)
			color.Cyan("📊 Baseline (%s): %.2f ns/op", baseline.Date, baseline.NsPerOp)
			color.Cyan("📊 Atual: %.2f ns/op", current.NsPerOp)

			if diff > 10 {
				color.Red("🚨 REGRESSÃO DETECTADA: +%.2f%% de latência!", diff)
				color.Yellow("O limite aceitável é 10%%. Por favor, otimize o código.")
			} else if diff < -5 {
				color.Green("🚀 MELHORIA DETECTADA: %.2f%% mais rápido!", -diff)
			} else {
				color.Green("✅ Performance estável (variação de %.2f%%).", diff)
			}
		} else {
			color.Yellow("⚠️  Nenhum baseline encontrado para '%s'. Definindo este como o primeiro registro.", target)
		}

		// 4. Salvar novo baseline
		err = benchmark.SaveBaseline(root, target, current)
		if err != nil {
			color.Red("⚠️  Erro ao salvar baseline no LEARNINGS.md: %v", err)
		} else {
			color.Blue("💾 Resultado persistido em LEARNINGS.md")
		}
	},
}

func init() {
	rootCmd.AddCommand(benchCmd)
}

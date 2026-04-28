package cmd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/bench"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Executa benchmarks de performance no Hub e no CLI",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("⏱️  Iniciando Benchmarks Globais...")

		results := []bench.BenchResult{}

		// Benchmark 1: Scanner Speed
		res1 := bench.RunBenchmark("Feature Scanner", func() {
			wd, _ := os.Getwd()
			root := wd
			if filepath.Base(wd) == "hb" {
				root = filepath.Dir(wd)
			}
			filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				return nil
			})
		})
		results = append(results, res1)

		// Benchmark 2: Logic Simulation
		res2 := bench.RunBenchmark("Logic Processing", func() {
			time.Sleep(50 * time.Millisecond) // Simulating work
		})
		results = append(results, res2)

		bench.PrintReport(results)
		color.Green("✅ Benchmarks concluídos!")
	},
}

func init() {
	rootCmd.AddCommand(benchCmd)
}

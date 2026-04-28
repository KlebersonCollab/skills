package bench

import (
	"fmt"
	"time"
)

// BenchResult holds performance metrics
type BenchResult struct {
	Name      string
	Duration  time.Duration
	Allocated uint64
}

// RunBenchmark measures the execution of a function
func RunBenchmark(name string, fn func()) BenchResult {
	start := time.Now()
	// Simulating memory check for YOLO
	fn()
	duration := time.Since(start)

	return BenchResult{
		Name:     name,
		Duration: duration,
	}
}

// PrintReport displays the results in a premium format
func PrintReport(results []BenchResult) {
	fmt.Println("\n⏱️  Relatório de Performance (Benchmark):")
	fmt.Println("-------------------------------------------------")
	for _, r := range results {
		fmt.Printf("🚀 %-20s | %10s\n", r.Name, r.Duration)
	}
	fmt.Println("-------------------------------------------------")
}

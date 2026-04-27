package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/domain"
	"github.com/klebersonromero/hb/internal/git"
	"github.com/klebersonromero/hb/internal/stats"
	"text/tabwriter"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Exibe estatísticas de uso e custo da sessão atual",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		s, err := stats.LoadStats(root)
		if err != nil || s == nil {
			color.Red("❌ Nenhuma sessão ativa com estatísticas encontrada.")
			return
		}

		// Atualiza métricas de Git em tempo real
		added, removed, _ := git.GetDiffStats()
		s.Git.LinesAdded = added
		s.Git.LinesRemoved = removed

		color.Cyan("\n📊 ESTATÍSTICAS DA SESSÃO: %s", s.SessionID)
		fmt.Printf("Iniciada em: %s\n", s.StartTime.Format("15:04:05"))
		fmt.Printf("Duração: %s\n", s.StartTime.Format("15:04:05"))

		fmt.Println("\n--- Uso por Modelo ---")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "MODELO\tINPUT\tOUTPUT\tCACHE R\tCUSTO (USD)")

		for name, usage := range s.Models {
			fmt.Fprintf(w, "%s\t%d\t%d\t%d\t$%.4f\n",
				name,
				usage.InputTokens,
				usage.OutputTokens,
				usage.CacheReadTokens,
				usage.CostUSD,
			)
		}
		w.Flush()

		color.Green("\n💰 Custo Total Estimado: $%.4f", s.TotalCostUSD)
		color.Yellow("📝 Alterações de Código (Git): +%d -%d linhas", s.Git.LinesAdded, s.Git.LinesRemoved)
		fmt.Println()
	},
}

var (
	modelName string
	inputT    int
	outputT   int
	cacheRT   int
)

var statsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uso de tokens manualmente à sessão",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		s, err := stats.LoadStats(root)
		if err != nil || s == nil {
			color.Red("❌ Erro: Sessão não encontrada.")
			return
		}

		if s.Models == nil {
			s.Models = make(map[string]*domain.ModelUsage)
		}

		usage, ok := s.Models[modelName]
		if !ok {
			usage = &domain.ModelUsage{}
			s.Models[modelName] = usage
		}

		usage.InputTokens += inputT
		usage.OutputTokens += outputT
		usage.CacheReadTokens += cacheRT

		// Recalcula custo
		cost := stats.CalculateCost(modelName, usage)
		usage.CostUSD = cost

		// Totaliza
		s.TotalCostUSD = 0
		for _, m := range s.Models {
			s.TotalCostUSD += m.CostUSD
		}

		err = stats.SaveStats(root, s)
		if err != nil {
			color.Red("❌ Erro ao salvar: %v", err)
			return
		}

		color.Green("✅ Uso registrado para %s. Novo custo total: $%.4f", modelName, s.TotalCostUSD)
	},
}

func init() {
	statsAddCmd.Flags().StringVarP(&modelName, "model", "m", "gemini-1.5-flash", "Nome do modelo")
	statsAddCmd.Flags().IntVarP(&inputT, "input", "i", 0, "Tokens de entrada")
	statsAddCmd.Flags().IntVarP(&outputT, "output", "o", 0, "Tokens de saída")
	statsAddCmd.Flags().IntVarP(&cacheRT, "cache", "c", 0, "Tokens de cache read")

	statsCmd.AddCommand(statsAddCmd)
	rootCmd.AddCommand(statsCmd)
}

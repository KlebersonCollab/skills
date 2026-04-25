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

var auditSec      bool
var auditMentor   bool
var auditAll      bool

var harnessAuditCmd = &cobra.Command{
	Use:   "audit [path]",
	Short: "Executa a Auditoria Unificada (Segurança e Qualidade)",
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		if !auditSec && !auditMentor {
			auditAll = true
		}

		config := harness.AuditConfig{
			Security: auditSec || auditAll,
			Mentor:   auditMentor || auditAll,
			Path:     path,
		}

		passed, err := harness.ExecuteAudit(config)
		if err != nil {
			color.Red("❌ Erro fatal na auditoria: %v", err)
			os.Exit(1)
		}

		if !passed {
			os.Exit(1)
		}
	},
}

var distillState   bool
var distillCode    bool
var distillFeature string
var distillOverwrite bool

var harnessDistillCmd = &cobra.Command{
	Use:   "distill [path]",
	Short: "Otimiza o contexto (Memória ou Código) para reduzir tokens",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		if distillState {
			err := harness.ExecuteDistillState(root, distillFeature)
			if err != nil {
				color.Red("❌ Erro ao destilar estado: %v", err)
				os.Exit(1)
			}
			return
		}

		if distillCode {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}
			err := harness.ExecuteDistillCode(path, distillOverwrite)
			if err != nil {
				color.Red("❌ Erro ao destilar código: %v", err)
				os.Exit(1)
			}
			return
		}

		color.Yellow("⚠️  Especifique --state ou --code")
	},
}

var contractCmd = &cobra.Command{
	Use:   "contract [feature]",
	Short: "Valida se a implementação respeita o contrato (contract.md) da feature",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		passed, err := harness.ExecuteContractCheck(root, feature)
		if err != nil {
			color.Red("❌ Erro ao validar contrato: %v", err)
			os.Exit(1)
		}

		if !passed {
			os.Exit(1)
		}
	},
}

var harnessMapCmd = &cobra.Command{
	Use:   "map",
	Short: "Sincroniza o mapa de conhecimento com as mudanças recentes",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		err := harness.ExecuteKnowledgeSync(root, applyMap) // Usa applyMap definido em map.go se compartilhado, ou local
		if err != nil {
			color.Red("❌ Erro ao sincronizar mapa: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	evaluateCmd.Flags().Float64Var(&designScore, "design", 0, "Nota para Design Quality (0.3)")
	evaluateCmd.Flags().Float64Var(&originalityScore, "originality", 0, "Nota para Originality (0.2)")
	evaluateCmd.Flags().Float64Var(&craftScore, "craft", 0, "Nota para Craft (0.3)")
	evaluateCmd.Flags().Float64Var(&functionalityScore, "functionality", 0, "Nota para Functionality (0.2)")

	harnessAuditCmd.Flags().BoolVar(&auditSec, "sec", false, "Executa apenas auditoria de segurança")
	harnessAuditCmd.Flags().BoolVar(&auditMentor, "mentor", false, "Executa apenas auditoria de qualidade (mentor)")
	harnessAuditCmd.Flags().BoolVar(&auditAll, "all", false, "Executa todas as auditorias (padrão)")

	harnessDistillCmd.Flags().BoolVar(&distillState, "state", false, "Destila o estado da memória da feature")
	harnessDistillCmd.Flags().BoolVar(&distillCode, "code", false, "Destila o código para reduzir tokens")
	harnessDistillCmd.Flags().StringVar(&distillFeature, "feature", "", "Nome da feature para destilar estado")
	harnessDistillCmd.Flags().BoolVar(&distillOverwrite, "overwrite", false, "Sobrescreve o arquivo com a versão destilada")

	harnessMapCmd.Flags().BoolVar(&applyMap, "apply", false, "Aplica as sugestões de mapeamento automaticamente")

	harnessCmd.AddCommand(evaluateCmd)
	harnessCmd.AddCommand(rehydrateCmd)
	harnessCmd.AddCommand(harnessAuditCmd)
	harnessCmd.AddCommand(harnessDistillCmd)
	harnessCmd.AddCommand(contractCmd)
	harnessCmd.AddCommand(harnessMapCmd)
	rootCmd.AddCommand(harnessCmd)
}

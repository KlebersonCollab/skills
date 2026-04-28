package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/git"
	"github.com/klebersonromero/hb/internal/review"
	"github.com/klebersonromero/hb/internal/sdd"
	"github.com/klebersonromero/hb/internal/security"
	"github.com/klebersonromero/hb/internal/ui"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Realiza uma auditoria semântica das mudanças no stage contra o contexto do SDD",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🔍 Iniciando Structured Review...")

		wd, _ := os.Getwd()
		
		// 1. Load SDD Context
		ctx, err := sdd.LoadActiveContext(wd)
		if err != nil {
			color.Red("❌ Erro ao carregar contexto SDD: %v", err)
			os.Exit(1)
		}

		fmt.Printf("📍 Feature: %s\n", color.CyanString(ctx.Name))
		fmt.Printf("📝 Task Ativa: %s [%s]\n", color.YellowString(ctx.ActiveTask.Description), color.WhiteString(ctx.ActiveTask.ID))

		// 2. Capture & Parse Diff
		uiManager := ui.NewUIManager(os.Stdout)
		uiManager.Start()
		uiManager.SetState(ui.UIState{
			FeatureName: ctx.Name,
			CurrentTask: "Capturando mudanças do stage...",
			Status:      "AUDITANDO",
		})

		diffStr, err := git.StagedDiff()
		if err != nil {
			uiManager.Stop()
			color.Red("❌ Erro ao capturar diff: %v", err)
			os.Exit(1)
		}

		if diffStr == "" {
			uiManager.Stop()
			color.Yellow("⚠️ Nenhum mudança encontrada no stage. Use 'git add' primeiro.")
			return
		}

		hunks, _ := git.ParseDiff(diffStr)
		uiManager.SetState(ui.UIState{
			FeatureName: ctx.Name,
			CurrentTask: fmt.Sprintf("Analisando %d mudanças semânticas...", len(hunks)),
			Status:      "PROCESSANDO",
			Progress:    0.5,
		})

		// 3. Perform Audit
		securityResult := security.ScanHunks(hunks)
		if len(securityResult.Violations) > 0 {
			uiManager.Stop()
			color.Red("\n🚨 FALHA DE SEGURANÇA DETECTADA!")
			for _, v := range securityResult.Violations {
				fmt.Println(v)
			}
			color.Yellow("\nSane os segredos/PII antes de prosseguir.")
			os.Exit(1)
		}

		result := review.AuditStagedChanges(ctx, hunks)
		uiManager.Stop()

		// 4. Render Premium Report
		renderAuditReport(result)

		if !result.Pass {
			color.Red("\n❌ AUDITORIA FALHOU (Score: %d/100)", result.Score)
			os.Exit(1)
		} else {
			color.Green("\n✅ AUDITORIA APROVADA (Score: %d/100)", result.Score)
		}
	},
}

func renderAuditReport(res *review.AuditResult) {
	fmt.Println("\n" + color.New(color.Bold).Sprint("Relatório de Auditoria:"))
	
	if len(res.Warnings) == 0 {
		fmt.Println(color.GreenString("  ✨ Nenhuma irregularidade encontrada."))
		return
	}

	for _, warn := range res.Warnings {
		fmt.Printf("  ⚠️  %s\n", color.YellowString(warn))
	}
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}

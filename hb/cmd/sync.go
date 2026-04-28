package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/harness"
	"github.com/klebersonromero/hb/internal/rules"
	"github.com/klebersonromero/hb/internal/syncer"
	"github.com/spf13/cobra"
)

var remoteSync bool

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sincroniza mandatos globais e diretórios de skills com os agentes",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🔄 Iniciando sincronização do ecossistema...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// 0. Harness Gate (Evaluation & Audit Check)
		score, feature, err := harness.GetLatestScore(root)
		if err == nil && score < 7.0 {
			color.Red("❌ SYNC BLOQUEADO: A feature '%s' possui um score de %.1f (abaixo do threshold de 7.0).", feature, score)
			color.Yellow("Por favor, melhore a qualidade da entrega e execute 'hb harness evaluate' novamente.")
			os.Exit(1)
		}

		auditPassed, auditFile, _ := harness.CheckAuditStatus(root)
		if !auditPassed {
			color.Red("❌ SYNC BLOQUEADO: Auditoria de Segurança/Qualidade falhou!")
			color.Yellow("Relatório com falha: %s", auditFile)
			color.Yellow("Por favor, corrija os problemas críticos e execute 'hb harness audit' novamente.")
			os.Exit(1)
		}

		// 1. Download de Mandatos (opcional)
		if remoteSync {
			err = syncer.DownloadMandates(root)
			if err != nil {
				color.Red("⚠️  Falha ao baixar mandatos: %v", err)
			}
		}

		// 2. Injeção de Mandatos Globais
		err = syncer.SyncMandates(root)
		if err != nil {
			color.Red("Erro ao sincronizar mandatos: %v", err)
			os.Exit(1)
		}

		// 3. Sincronização de Skills (Paralela)
		err = syncer.SyncSkills(root)
		if err != nil {
			color.Red("Erro ao sincronizar skills: %v", err)
			os.Exit(1)
		}

		rulesFlag, _ := cmd.Flags().GetBool("rules")
		if rulesFlag {
			color.Blue("🔨 Compilando regras do projeto (.cursorrules)...")
			if err := rules.CompileRules(root); err != nil {
				color.Red("❌ Erro ao compilar regras: %v", err)
			} else {
				color.Green("✨ .cursorrules atualizado com sucesso!")
			}
		}

		color.Green("🚀 Sincronização concluída!")
	},
}

func init() {
	syncCmd.Flags().BoolVarP(&remoteSync, "remote", "r", false, "Baixa os mandatos globais do repositório remoto antes de sincronizar")
	syncCmd.Flags().Bool("rules", false, "Compila as regras do projeto (.cursorrules)")
	rootCmd.AddCommand(syncCmd)
}

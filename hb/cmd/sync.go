package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
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

		// 1. Download de Mandatos (opcional)
		if remoteSync {
			err := syncer.DownloadMandates(root)
			if err != nil {
				color.Red("⚠️  Falha ao baixar mandatos: %v", err)
			}
		}

		// 2. Injeção de Mandatos Globais
		err := syncer.SyncMandates(root)
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

		color.Green("✨ Sincronização concluída com sucesso!")
	},
}

func init() {
	syncCmd.Flags().BoolVarP(&remoteSync, "remote", "r", false, "Baixa os mandatos globais do repositório remoto antes de sincronizar")
	rootCmd.AddCommand(syncCmd)
}

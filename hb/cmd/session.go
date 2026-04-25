package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/session"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Gerencia sessões de desenvolvimento e captura de estado",
}

var sessionStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia uma nova sessão de desenvolvimento",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := session.Start(root)
		if err != nil {
			color.Red("❌ Erro ao iniciar sessão: %v", err)
			os.Exit(1)
		}
	},
}

var sessionStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Encerra a sessão atual e resume as atividades",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := session.Stop(root)
		if err != nil {
			color.Red("❌ Erro ao encerrar sessão: %v", err)
			os.Exit(1)
		}
	},
}

var sessionHandoffCmd = &cobra.Command{
	Use:   "handoff",
	Aliases: []string{"end"},
	Short: "Gera um relatório de handoff para a próxima sessão",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		err := session.ExecuteHandoff(root)
		if err != nil {
			color.Red("❌ Erro ao realizar handoff: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	sessionCmd.AddCommand(sessionStartCmd)
	sessionCmd.AddCommand(sessionStopCmd)
	sessionCmd.AddCommand(sessionHandoffCmd)
	rootCmd.AddCommand(sessionCmd)
}

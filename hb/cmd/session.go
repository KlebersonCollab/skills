package cmd

import (
	"os"
	"path/filepath"

	"github.com/klebersonromero/hb/internal/session"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Gerencia sessões atômicas de desenvolvimento",
}

var sessionStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia uma nova sessão atômica",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		session.Start(root)
	},
}

var sessionStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Encerra a sessão atual e exibe o sumário",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		session.Stop(root)
	},
}

func init() {
	sessionCmd.AddCommand(sessionStartCmd)
	sessionCmd.AddCommand(sessionStopCmd)
	rootCmd.AddCommand(sessionCmd)
}

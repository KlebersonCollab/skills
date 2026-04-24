package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/git"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Gerencia o fluxo de trabalho Git (Mandato 9)",
}

var gitStartCmd = &cobra.Command{
	Use:   "start [feature]",
	Short: "Cria uma nova branch de feature",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := git.Start(args[0])
		if err != nil {
			color.Red("❌ %v", err)
			os.Exit(1)
		}
	},
}

var gitCommitType string
var gitCommitScope string

var gitCommitCmd = &cobra.Command{
	Use:   "commit [mensagem]",
	Short: "Realiza um commit seguindo o padrão Conventional Commits",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := git.Commit(gitCommitType, gitCommitScope, args[0])
		if err != nil {
			color.Red("❌ %v", err)
			os.Exit(1)
		}
	},
}

var gitSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sincroniza a branch atual com a main via rebase",
	Run: func(cmd *cobra.Command, args []string) {
		err := git.Sync()
		if err != nil {
			color.Red("❌ %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	gitCommitCmd.Flags().StringVarP(&gitCommitType, "type", "t", "feat", "Tipo do commit (feat, fix, docs, chore)")
	gitCommitCmd.Flags().StringVarP(&gitCommitScope, "scope", "s", "core", "Escopo do commit")

	gitCmd.AddCommand(gitStartCmd)
	gitCmd.AddCommand(gitCommitCmd)
	gitCmd.AddCommand(gitSyncCmd)
	rootCmd.AddCommand(gitCmd)
}

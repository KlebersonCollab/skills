package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/mentor"
	"github.com/spf13/cobra"
)

var mentorCmd = &cobra.Command{
	Use:   "mentor",
	Short: "Comandos de mentoria e revisão de Clean Code",
}

var mentorAuditCmd = &cobra.Command{
	Use:   "audit [arquivo]",
	Short: "Analisa um arquivo em busca de violações de princípios de Clean Code",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		err := mentor.AuditFile(path)
		if err != nil {
			color.Red("❌ Erro na auditoria: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	mentorCmd.AddCommand(mentorAuditCmd)
	rootCmd.AddCommand(mentorCmd)
}

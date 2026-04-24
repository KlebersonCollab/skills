package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/doctor"
	"github.com/spf13/cobra"
)

var fixFlag bool

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnostica e corrige problemas no ecossistema Harness",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		checks, err := doctor.Run(root)
		if err != nil {
			color.Red("❌ Erro ao rodar diagnostico: %v", err)
			os.Exit(1)
		}

		allPassed := true
		for _, c := range checks {
			if c.Passed {
				color.Green("✅ %s: %s", c.Name, c.Message)
			} else {
				color.Red("❌ %s: %s", c.Name, c.Message)
				allPassed = false
			}
		}

		if !allPassed && fixFlag {
			doctor.Fix(root)
		}
	},
}

func init() {
	doctorCmd.Flags().BoolVar(&fixFlag, "fix", false, "Tenta corrigir os problemas encontrados")
	rootCmd.AddCommand(doctorCmd)
}

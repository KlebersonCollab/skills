package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/report"
	"github.com/klebersonromero/hb/internal/scanner"
	"github.com/spf13/cobra"
)

var jsonOutput bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as skills registradas no Hub",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		skills, err := scanner.FindSkills(root)
		if err != nil {
			color.Red("❌ Erro ao buscar skills: %v", err)
			os.Exit(1)
		}

		if jsonOutput {
			data, _ := json.MarshalIndent(skills, "", "  ")
			fmt.Println(string(data))
			return
		}

		color.Blue("📋 Skills Registradas no Hub:")
		report.PrintSkillList(skills)
	},
}

func init() {
	listCmd.Flags().BoolVar(&jsonOutput, "json", false, "Exibe o resultado em formato JSON")
	rootCmd.AddCommand(listCmd)
}

package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/rules"
	"github.com/spf13/cobra"
)

var skillsList string

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Gera o arquivo .cursorrules com mandatos e definições de skills",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🔧 Gerando .cursorrules...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		var targetSkills []string
		if skillsList != "" {
			targetSkills = strings.Split(skillsList, ",")
			for i := range targetSkills {
				targetSkills[i] = strings.TrimSpace(targetSkills[i])
			}
		}

		err := rules.Generate(root, targetSkills)
		if err != nil {
			color.Red("❌ Erro ao gerar regras: %v", err)
			os.Exit(1)
		}

		color.Green("✅ .cursorrules gerado com sucesso!")
	},
}

func init() {
	rulesCmd.Flags().StringVarP(&skillsList, "skills", "s", "", "Lista de skills separadas por vírgula")
	rootCmd.AddCommand(rulesCmd)
}

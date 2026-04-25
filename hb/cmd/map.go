package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/harness"
	"github.com/klebersonromero/hb/internal/mapper"
	"github.com/spf13/cobra"
)

var impactFile string
var autoMap    bool
var applyMap   bool

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Gera o mapa de conhecimento relacional (KNOWLEDGE-MAP.mermaid)",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		if autoMap || applyMap {
			err := harness.ExecuteKnowledgeSync(root, applyMap)
			if err != nil {
				color.Red("❌ Erro no Knowledge Sync: %v", err)
				os.Exit(1)
			}
			return
		}

		if impactFile != "" {
			color.Blue("🔍 Analisando impacto de: %s", impactFile)
			impacted, err := mapper.AnalyzeImpact(root, impactFile)
			if err != nil {
				color.Red("❌ Erro na análise: %v", err)
				os.Exit(1)
			}

			if len(impacted) == 0 {
				color.Yellow("Nenhuma skill ou documento diretamente dependente encontrado.")
			} else {
				color.Cyan("\nSkills/Documentos potencialmente impactados:")
				for _, s := range impacted {
					fmt.Printf(" - %s\n", s)
				}
			}
			return
		}

		color.Blue("🗺️  Gerando mapa de conhecimento do ecossistema...")
		err := mapper.Generate(root)
		if err != nil {
			color.Red("❌ Erro ao gerar mapa: %v", err)
			os.Exit(1)
		}

		color.Green("✅ KNOWLEDGE-MAP.mermaid gerado com sucesso!")
	},
}

func init() {
	mapCmd.Flags().StringVarP(&impactFile, "impact", "i", "", "Analisa o impacto de alterar um arquivo específico")
	mapCmd.Flags().BoolVar(&autoMap, "auto", false, "Sincroniza o mapa com as mudanças recentes do Git")
	mapCmd.Flags().BoolVar(&applyMap, "apply", false, "Aplica as sugestões de mapeamento automaticamente")
	rootCmd.AddCommand(mapCmd)
}

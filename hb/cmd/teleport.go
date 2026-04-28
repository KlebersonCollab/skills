package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/knowledge"
	"github.com/klebersonromero/hb/internal/sdd"
	"github.com/spf13/cobra"
)

var teleportCmd = &cobra.Command{
	Use:   "teleport",
	Short: "Gerencia o teleporte de aprendizados entre features e projetos",
}

var teleportExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exporta os aprendizados da feature atual para o hub global",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		ctx, err := sdd.LoadActiveContext(wd)
		if err != nil {
			color.Red("❌ Erro: Nenhuma feature ativa encontrada para exportar.")
			return
		}

		color.Blue("🚀 Teleportando conhecimentos de '%s'...", ctx.Name)

		bundle, err := knowledge.ExportFeature(wd, ctx.Name)
		if err != nil {
			color.Red("❌ Erro ao empacotar conhecimentos: %v", err)
			return
		}

		err = knowledge.SaveBundle(wd, bundle)
		if err != nil {
			color.Red("❌ Erro ao salvar no hub global: %v", err)
			return
		}

		color.Green("✨ Conhecimento teleportado com sucesso! (ID: %s)", bundle.ID)
		fmt.Printf("📂 Local: .specs/knowledge/global/%s_%s.json\n", bundle.Feature, bundle.ID)
	},
}

var teleportImportCmd = &cobra.Command{
	Use:   "import [bundle-filename]",
	Short: "Importa aprendizados de um bundle global para a feature atual",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bundleName := args[0]
		wd, _ := os.Getwd()
		ctx, err := sdd.LoadActiveContext(wd)
		if err != nil {
			color.Red("❌ Erro: Nenhuma feature ativa encontrada para importar.")
			return
		}

		color.Blue("📥 Importando conhecimentos de '%s' para '%s'...", bundleName, ctx.Name)

		bundle, err := knowledge.ImportBundle(wd, bundleName)
		if err != nil {
			color.Red("❌ Erro ao ler bundle: %v", err)
			return
		}

		err = knowledge.MergeKnowledge(wd, ctx.Name, bundle)
		if err != nil {
			color.Red("❌ Erro ao injetar conhecimento: %v", err)
			return
		}

		color.Green("✨ Conhecimento herdado com sucesso de %s!", bundle.Feature)
	},
}

func init() {
	teleportCmd.AddCommand(teleportExportCmd)
	teleportCmd.AddCommand(teleportImportCmd)
	rootCmd.AddCommand(teleportCmd)
}

package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/installer"
	"github.com/spf13/cobra"
)

var (
	updateAll bool
)

var updateCmd = &cobra.Command{
	Use:   "update [skill-name]",
	Short: "Atualiza uma skill ou todas as skills instaladas",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		opts := installer.Options{
			Target: targetDir,
			Remote: remote,
			Root:   root,
		}

		skillName := ""
		if len(args) > 0 {
			skillName = args[0]
			opts.Name = skillName
		} else if !updateAll {
			color.Red("❌ Erro: Especifique o nome da skill ou use --all")
			os.Exit(1)
		}

		err := installer.Update(opts, updateAll)
		if err != nil {
			color.Red("❌ Erro na atualização: %v", err)
			os.Exit(1)
		}

		if updateAll {
			color.Green("✅ Todas as skills foram atualizadas!")
		} else {
			color.Green("✅ Skill '%s' atualizada com sucesso!", skillName)
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&targetDir, "target", "t", "./.agent/skills", "Diretório de destino")
	updateCmd.Flags().BoolVarP(&remote, "remote", "r", false, "Baixar do repositório remoto via Git")
	updateCmd.Flags().BoolVarP(&updateAll, "all", "a", false, "Atualizar todas as skills instaladas")
	rootCmd.AddCommand(updateCmd)
}

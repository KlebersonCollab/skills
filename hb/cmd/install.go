package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/installer"
	"github.com/spf13/cobra"
)

var (
	targetDir string
	force     bool
	remote    bool
)

var installCmd = &cobra.Command{
	Use:   "install [skill-name]",
	Short: "Instala uma skill no diretório de destino",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		skillName := args[0]
		color.Blue("🚀 Preparando instalação da skill: %s", skillName)

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		opts := installer.Options{
			Name:   skillName,
			Target: targetDir,
			Force:  force,
			Remote: remote,
			Root:   root,
		}

		err := installer.Install(opts)
		if err != nil {
			color.Red("❌ Erro na instalação: %v", err)
			os.Exit(1)
		}

		color.Green("✅ Skill '%s' instalada com sucesso!", skillName)
	},
}

func init() {
	installCmd.Flags().StringVarP(&targetDir, "target", "t", "./.agent/skills", "Diretório de destino")
	installCmd.Flags().BoolVarP(&force, "force", "f", false, "Sobrescrever se a skill já existir")
	installCmd.Flags().BoolVarP(&remote, "remote", "r", false, "Baixar do repositório remoto via Git")
	rootCmd.AddCommand(installCmd)
}

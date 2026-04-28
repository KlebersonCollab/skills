package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Verifica a saúde do ambiente de desenvolvimento e do Hub",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🩺 Iniciando diagnóstico do sistema...")

		checkCommand("go", "Go Compiler")
		checkCommand("git", "Git Version Control")
		
		wd, _ := os.Getwd()
		checkDir(filepath.Join(wd, ".specs"), "SDD Specs Directory")
		checkDir(filepath.Join(wd, ".specs", "features"), "Features Directory")
		checkDir(filepath.Join(wd, "hb"), "CLI Source Directory")

		color.Green("\n✅ Diagnóstico concluído! Se não houver erros acima, seu ambiente está pronto.")
	},
}

func checkCommand(name, label string) {
	_, err := exec.LookPath(name)
	if err != nil {
		color.Red("❌ %s: NÃO ENCONTRADO", label)
	} else {
		color.Green("✅ %s: Instalado", label)
	}
}

func checkDir(path, label string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		color.Red("❌ %s: FALTANDO (%s)", label, path)
	} else {
		color.Green("✅ %s: Presente", label)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listSkillsCmd = &cobra.Command{
	Use:   "listskills",
	Short: "Lista todas as skills disponíveis no repositório remoto (GitHub)",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("🔍 Buscando skills no repositório remoto...")
		
		repoURL := "https://github.com/KlebersonCollab/skills.git"
		
		// Criar diretório temporário para clone otimizado
		tmpDir, err := os.MkdirTemp("", "hb-listskills-*")
		if err != nil {
			color.Red("❌ Erro ao criar diretório temporário: %v", err)
			os.Exit(1)
		}
		defer os.RemoveAll(tmpDir)

		// Clone extremamente raso apenas para metadados
		cloneCmd := exec.Command("git", "clone", "--depth", "1", "--filter=blob:none", "--sparse", repoURL, tmpDir)
		if out, err := cloneCmd.CombinedOutput(); err != nil {
			color.Red("❌ Erro ao acessar repositório remoto: %v\n%s", err, string(out))
			os.Exit(1)
		}

		// Listar apenas caminhos que contenham SKILL.md
		lsTree := exec.Command("git", "-C", tmpDir, "ls-tree", "-r", "--name-only", "HEAD")
		out, err := lsTree.Output()
		if err != nil {
			color.Red("❌ Erro ao listar árvore de arquivos: %v", err)
			os.Exit(1)
		}

		lines := strings.Split(string(out), "\n")
		var skills []string
		seen := make(map[string]bool)

		for _, line := range lines {
			if strings.HasSuffix(line, "/SKILL.md") {
				// Extrair o nome da skill (penúltimo segmento do path)
				parts := strings.Split(line, "/")
				if len(parts) >= 2 {
					skillName := parts[len(parts)-2]
					
					// Ignorar diretórios internos ou temporários
					if skillName == ".agents" || skillName == ".agent" || strings.Contains(line, "temp-skills") {
						continue
					}
					
					if !seen[skillName] {
						skills = append(skills, skillName)
						seen[skillName] = true
					}
				}
			}
		}

		if len(skills) == 0 {
			color.Yellow("⚠️  Nenhuma skill encontrada no repositório remoto.")
			return
		}

		color.Green("✅ %d Skills encontradas no GitHub:", len(skills))
		for _, s := range skills {
			fmt.Printf("  📦 %-20s  (hb install %s --remote)\n", s, s)
		}
		
		color.Cyan("\nDica: Use 'hb install <name> --remote' para baixar e instalar automaticamente.")
	},
}

func init() {
	rootCmd.AddCommand(listSkillsCmd)
}

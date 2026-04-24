package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/installer"
	"github.com/klebersonromero/hb/internal/rules"
	"github.com/klebersonromero/hb/internal/syncer"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa um novo projeto com o ecossistema Harness completo",
	Long: `O comando init realiza o setup completo do projeto:
1. Cria a estrutura global de governança (.specs/)
2. Baixa os mandatos globais do repositório remoto
3. Instala as skills essenciais (sdd, harness-expert, git-workflow)
4. Sincroniza os agentes (.gemini, .agent, etc.)
5. Gera o arquivo .cursorrules inicial`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Magenta("🚀 Iniciando Setup de Governança Harness...")

		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// 1. Bootstrap (Pastas e arquivos base)
		color.Blue("\n📂 1. Criando estrutura global de governança...")
		// Chamamos a lógica do bootstrapCmd manualmente ou replicamos aqui
		dirs := []string{
			filepath.Join(root, ".specs", "project"),
			filepath.Join(root, ".specs", "architecture"),
			filepath.Join(root, ".specs", "codebase"),
			filepath.Join(root, ".specs", "features"),
		}
		for _, d := range dirs {
			os.MkdirAll(d, 0755)
		}
		color.Green("   ✅ Estrutura criada.")

		// 2. Mandatos
		color.Blue("\n📜 2. Sincronizando Mandatos Globais...")
		err := syncer.DownloadMandates(root)
		if err != nil {
			color.Red("   ⚠️  Falha ao baixar mandatos: %v", err)
		}

		// 3. Skills Essenciais
		color.Blue("\n🛠️  3. Instalando Skills Core (SDD, Harness, Git)...")
		coreSkills := []string{"sdd", "harness-expert", "git-workflow"}
		for _, skill := range coreSkills {
			color.Cyan("   -> Instalando %s...", skill)
			opts := installer.Options{
				Name:   skill,
				Target: filepath.Join(root, ".agents", "skills"), // Local padrão
				Remote: true,
				Force:  true,
				Root:   root,
			}
			err := installer.Install(opts)
			if err != nil {
				color.Red("   ❌ Erro ao instalar %s: %v", skill, err)
			}
		}

		// 4. Sync Agentes
		color.Blue("\n🔄 4. Sincronizando agentes...")
		err = syncer.SyncMandates(root)
		if err == nil {
			err = syncer.SyncSkills(root)
		}
		if err != nil {
			color.Red("   ❌ Erro na sincronização: %v", err)
		}

		// 5. Rules
		color.Blue("\n🔧 5. Compilando .cursorrules...")
		err = rules.Generate(root, nil)
		if err != nil {
			color.Red("   ❌ Erro ao gerar regras: %v", err)
		}

		color.Magenta("\n✨ Projeto inicializado com sucesso! O Harness está pronto para operar.")
		color.Cyan("Dica: Use 'hb sdd init <nome-da-feature>' para começar uma nova tarefa.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

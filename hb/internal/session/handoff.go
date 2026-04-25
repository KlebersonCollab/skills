package session

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ExecuteHandoff gera um relatório de handoff para a próxima sessão
func ExecuteHandoff(root string) error {
	color.Blue("📝 Iniciando protocolo de handoff de sessão...")

	// 1. Coletar commits recentes (últimas 4 horas)
	cmd := exec.Command("git", "log", "--since='4 hours ago'", "--oneline", "--decorate")
	cmd.Dir = root
	out, _ := cmd.Output()
	commits := string(out)

	// 2. Coletar tarefas marcadas recentemente (opcional, complexo sem rastreamento de tempo)
	
	// 3. Gerar o arquivo HANDOFF.md
	date := time.Now().Format("2006-01-02 15:04")
	handoffPath := filepath.Join(root, "HANDOFF.md")
	
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# 🏁 Session Handoff - %s\n\n", date))
	sb.WriteString("> Este arquivo contém o progresso da última sessão para facilitar a reidratação do próximo agente.\n\n")
	
	sb.WriteString("## 🚀 Progresso Recente (Commits)\n")
	if commits != "" {
		sb.WriteString("```bash\n")
		sb.WriteString(commits)
		sb.WriteString("```\n")
	} else {
		sb.WriteString("_Nenhum commit realizado nas últimas 4 horas._\n")
	}
	
	sb.WriteString("\n## 🎯 Objetivos Concluídos\n")
	sb.WriteString("- (O agente deve preencher este campo manualmente se necessário, ou o CLI pode extrair do Git)\n")
	
	sb.WriteString("\n## 📌 Próximos Passos Imediatos\n")
	sb.WriteString("- [ ] Retomar a partir do estado atual definido no `STATE.md`.\n")
	sb.WriteString("- [ ] Validar integridade do ambiente com `hb doctor`.\n")

	sb.WriteString("\n## 💡 Aprendizados Registrados\n")
	sb.WriteString("_Use `hb learn` para registrar padrões importantes descobertos nesta sessão._\n")

	err := os.WriteFile(handoffPath, []byte(sb.String()), 0644)
	if err != nil {
		return err
	}

	// 4. Sugerir atualização do STATE.md
	color.Green("✅ HANDOFF.md gerado com sucesso!")
	color.Yellow("👉 Lembre-se de usar 'hb focus [tarefa]' para definir o próximo objetivo no STATE.md.")
	
	return nil
}

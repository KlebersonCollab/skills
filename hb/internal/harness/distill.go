package harness

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/distiller"
)

// ExecuteDistillCode executa a minificação de um arquivo para economizar tokens
func ExecuteDistillCode(path string, overwrite bool) error {
	color.Blue("🔍 Destilando código em: %s", path)
	
	text, saved, err := distiller.DistillFile(path)
	if err != nil {
		return err
	}

	if overwrite {
		err = os.WriteFile(path, []byte(text), 0644)
		if err != nil {
			return err
		}
		color.Green("✅ Arquivo sobrescrito com sucesso.")
	} else {
		fmt.Println(text)
	}

	color.Magenta("✨ Economia: %d caracteres", saved)
	return nil
}

// ExecuteDistillState consolida o estado operacional da feature
func ExecuteDistillState(root, feature string) error {
	if feature == "" {
		return fmt.Errorf("nome da feature é obrigatório para --state")
	}

	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		return fmt.Errorf("feature '%s' não encontrada em .specs/features/", feature)
	}

	color.Blue("🧠 Destilando memória e estado para: %s", feature)

	// 1. Coletar Diffs do Git
	diff, _ := getGitDiff(root)
	
	// 2. Coletar Logs Recentes
	logs, _ := getGitLogs(root)

	// 3. Gerar resumo do estado
	statePath := filepath.Join(featureDir, "STATE.md")
	
	content := fmt.Sprintf(`# Operational State: %s
- **Snapshot**: %s

## 🛠️ Atividade Recente (Git)
### Arquivos Alterados:
%s

### Commits Recentes:
%s

## 📌 Contexto Atual
- [O agente deve complementar este bloco com os aprendizados da sessão atual]

## 📝 Próximos Passos
- Verifique o tasks.md da feature.
`, feature, time.Now().Format("2006-01-02 15:04:05"), diff, logs)

	err := os.WriteFile(statePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	color.Green("✅ STATE.md atualizado com sucesso em %s", statePath)
	return nil
}

func getGitDiff(root string) (string, error) {
	cmd := exec.Command("git", "diff", "--stat")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return "Não foi possível coletar diffs.", nil
	}
	return string(out), nil
}

func getGitLogs(root string) (string, error) {
	cmd := exec.Command("git", "log", "-n", "3", "--oneline")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return "Não foi possível coletar logs.", nil
	}
	return string(out), nil
}

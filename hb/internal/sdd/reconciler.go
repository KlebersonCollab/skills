package sdd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// ReconcileFeature sincroniza os artefatos de uma feature com a implementação real
func ReconcileFeature(root, feature string) error {
	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		return fmt.Errorf("feature '%s' não encontrada", feature)
	}

	color.Blue("🔄 Reconciliando SDD para: %s", feature)

	// 1. Reconciliar Tasks
	taskFile := filepath.Join(featureDir, "tasks.md")
	if _, err := os.Stat(taskFile); err == nil {
		updated, count, err := reconcileTasksFile(root, taskFile)
		if err != nil {
			return err
		}
		if updated {
			color.Green("   ✅ %d tarefas marcadas como concluídas via preenchimento reverso.", count)
		} else {
			color.Yellow("   ℹ️  Nenhuma nova tarefa concluída detectada.")
		}
	}

	// 2. Gerar Contrato se faltar
	contractPath := filepath.Join(featureDir, "contract.md")
	if _, err := os.Stat(contractPath); os.IsNotExist(err) {
		err := generateReverseContract(root, feature, contractPath)
		if err != nil {
			color.Red("   ❌ Erro ao gerar contrato reverso: %v", err)
		} else {
			color.Green("   ✅ contract.md gerado automaticamente baseado na implementação.")
		}
	}

	return nil
}

func reconcileTasksFile(root, path string) (bool, int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return false, 0, err
	}

	lines := strings.Split(string(content), "\n")
	updated := false
	count := 0

	for i, line := range lines {
		// Só processa linhas que são tarefas pendentes
		if strings.HasPrefix(line, "- [ ]") {
			description := strings.TrimPrefix(line, "- [ ]")
			if hasEvidence(root, description) {
				lines[i] = strings.Replace(line, "[ ]", "[x]", 1)
				updated = true
				count++
			}
		}
	}

	if updated {
		err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
		return true, count, err
	}

	return false, 0, nil
}

func hasEvidence(root, description string) bool {
	// Heurística 1: Nome de arquivo
	reFile := regexp.MustCompile(`([a-zA-Z0-9_\-.]+\.(go|py|js|md|yml|yaml|sh))`)
	if matches := reFile.FindStringSubmatch(description); len(matches) > 1 {
		filename := matches[1]
		found := false
		_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() { return err }
			if d.Name() == filename {
				found = true
				return filepath.SkipDir
			}
			return nil
		})
		if found { return true }
	}

	// Heurística 2: Comandos CLI
	if strings.Contains(strings.ToLower(description), "command") || strings.Contains(strings.ToLower(description), "comando") {
		reCmd := regexp.MustCompile("`([^`]+)`")
		if matches := reCmd.FindStringSubmatch(description); len(matches) > 1 {
			cmdName := matches[1]
			if strings.HasPrefix(cmdName, "hb ") {
				cmdName = strings.TrimPrefix(cmdName, "hb ")
				parts := strings.Fields(cmdName)
				if len(parts) > 0 {
					cmdName = parts[0]
				}
			}
			
			// Busca registro do comando
			grep := exec.Command("grep", "-r", `Use:.*"`+cmdName+`"`, filepath.Join(root, "hb", "cmd"))
			if err := grep.Run(); err == nil { return true }
		}
	}

	return false
}

func generateReverseContract(root, feature, path string) error {
	// Tenta encontrar comandos relacionados ao nome da feature
	// Simplificação: busca por termos da feature no hb/cmd
	featureTerm := feature
	if strings.Contains(feature, "-") {
		parts := strings.Split(feature, "-")
		featureTerm = parts[len(parts)-1]
	}

	// Busca interfaces
	cmd := exec.Command("grep", "-rE", `Use:\s*"[a-zA-Z0-9-]+"`, filepath.Join(root, "hb", "cmd"))
	out, _ := cmd.Output()
	lines := strings.Split(string(out), "\n")
	
	var discoveredCmds []string
	for _, l := range lines {
		if strings.Contains(l, featureTerm) {
			re := regexp.MustCompile(`Use:\s*"([^"]+)"`)
			if m := re.FindStringSubmatch(l); len(m) > 1 {
				discoveredCmds = append(discoveredCmds, m[1])
			}
		}
	}

	content := fmt.Sprintf("# Contract: %s (Auto-Generated)\n\n## Command Interface\n| Command | Description |\n| :--- | :--- |\n", feature)
	for _, c := range discoveredCmds {
		content += fmt.Sprintf("| `hb %s` | Detectado via preenchimento reverso. |\n", c)
	}

	if len(discoveredCmds) == 0 {
		content += "| `(pendente)` | Nenhuma interface detectada automaticamente. |\n"
	}

	return os.WriteFile(path, []byte(content), 0644)
}

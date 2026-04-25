package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// SetFocus atualiza o STATE.md com uma nova tarefa de foco
func SetFocus(root, task string) error {
	statePath := filepath.Join(root, ".specs", "project", "STATE.md")
	content, err := os.ReadFile(statePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Se não existir, tenta criar a estrutura básica
			err = os.MkdirAll(filepath.Dir(statePath), 0755)
			if err != nil {
				return err
			}
			content = []byte("# Project State\n\n## Current Status\n- **Focus**: Initializing...\n")
		} else {
			return err
		}
	}

	lines := strings.Split(string(content), "\n")
	date := time.Now().Format("2006-01-02")
	focusLine := fmt.Sprintf("- **Focus**: %s (%s) 🎯", task, date)

	updated := false
	for i, line := range lines {
		if strings.Contains(line, "**Focus**:") {
			lines[i] = focusLine
			updated = true
			break
		}
	}

	if !updated {
		// Procura por ## Current Status para inserir abaixo
		for i, line := range lines {
			if strings.Contains(line, "## Current Status") {
				newLines := append(lines[:i+1], append([]string{focusLine}, lines[i+1:]...)...)
				lines = newLines
				updated = true
				break
			}
		}
	}

	// Se ainda não atualizou, adiciona no fim
	if !updated {
		lines = append(lines, "## Current Status", focusLine)
	}

	err = os.WriteFile(statePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return err
	}

	color.Green("🎯 Foco atualizado em STATE.md: %s", task)
	return nil
}

// GetContextSummary gera um resumo conciso da tríade de memória
func GetContextSummary(root string) (string, error) {
	statePath := filepath.Join(root, ".specs", "project", "STATE.md")
	memoryPath := filepath.Join(root, ".specs", "project", "MEMORY.md")
	learningsPath := filepath.Join(root, ".specs", "project", "LEARNINGS.md")

	var summary strings.Builder

	summary.WriteString("🧠 **HB Project Context Summary**\n")
	summary.WriteString(strings.Repeat("-", 40) + "\n")

	// 1. STATE (Current Status)
	summary.WriteString("📍 **STATE (Status Atual):**\n")
	stateContent, err := os.ReadFile(statePath)
	if err == nil {
		lines := strings.Split(string(stateContent), "\n")
		count := 0
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "- **Focus**:") || 
			   strings.HasPrefix(trimmed, "- **Feature**:") || 
			   strings.HasPrefix(trimmed, "- **Status**:") ||
			   strings.HasPrefix(trimmed, "- **Session**:") {
				summary.WriteString("  " + trimmed + "\n")
				count++
			}
		}
		if count == 0 {
			summary.WriteString("  Nenhum foco ativo definido.\n")
		}
	} else {
		summary.WriteString("  STATE.md não encontrado.\n")
	}

	// 2. MEMORY (Conventions)
	summary.WriteString("\n📜 **MEMORY (Convenções Recentes):**\n")
	memoryContent, err := os.ReadFile(memoryPath)
	if err == nil {
		lines := strings.Split(string(memoryContent), "\n")
		count := 0
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "- **") && count < 3 {
				summary.WriteString("  " + trimmed + "\n")
				count++
			}
		}
	} else {
		summary.WriteString("  MEMORY.md não encontrado.\n")
	}

	// 3. LEARNINGS (Padrões e Bugs)
	summary.WriteString("\n💡 **LEARNINGS (Últimos Aprendizados):**\n")
	learningsContent, err := os.ReadFile(learningsPath)
	if err == nil {
		lines := strings.Split(string(learningsContent), "\n")
		count := 0
		// Pega as últimas 2 entradas ## [Data]
		for i := len(lines) - 1; i >= 0; i-- {
			if strings.HasPrefix(lines[i], "## [") || strings.HasPrefix(lines[i], "### [") {
				summary.WriteString("  " + strings.TrimSpace(lines[i]) + "\n")
				if i+1 < len(lines) {
					summary.WriteString("    " + strings.TrimSpace(lines[i+1]) + "\n")
				}
				count++
				if count >= 2 {
					break
				}
			}
		}
	} else {
		summary.WriteString("  LEARNINGS.md não encontrado.\n")
	}

	summary.WriteString(strings.Repeat("-", 40))
	return summary.String(), nil
}

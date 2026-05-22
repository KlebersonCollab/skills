
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func agentExecutionLoop(config *AppConfig, tree *SessionTree, sessionFile string, streamMode bool) {
	maxIterations := 25
	consecutiveFailures := 0
	maxConsecutiveFailures := 8

	for i := 0; i < maxIterations; i++ {
		historyText, err := tree.FormatLinearHistoryText("")
		if err != nil {
			fmt.Printf("❌ Erro ao formatar histórico: %v\n", err)
			return
		}

		root, err := FindWorkspaceRoot()
		skillsText := "Nenhuma skill ativa encontrada no Hub."
		if err == nil {
			activeSkills := DiscoverSkills(root)
			if len(activeSkills) > 0 {
				skillsText = "As seguintes Skills de Engenharia estão ativas e disponíveis no Hub (cada uma possui regras rígidas e diretrizes de desenvolvimento no arquivo SKILL.md de seu respectivo diretório):\n"
				for _, sk := range activeSkills {
					skillsText += fmt.Sprintf("  - %s: %s (caminho: %s/SKILL.md)\n", sk.Name, sk.Description, sk.Path)
				}
				skillsText += "\n💡 REQUISITO DE CONFORMIDADE OBRIGATÓRIA: Se o usuário pedir para usar uma skill ou se a tarefa envolver o domínio de uma delas (por exemplo, usar 'sdd', 'python-uv', 'clean-code-mentor', 'git-workflow', etc.), você DEVE obrigatoriamente ler o arquivo SKILL.md correspondente utilizando a ferramenta de leitura <tool:read_file path=\"nome_da_skill/SKILL.md\"/> para entender e aplicar todas as regras de qualidade, governança e engenharia nela especificadas. Não adivinhe as diretrizes; leia a Skill correspondente antes de continuar a execução.\n"
			}
		}

		agentsMandates := ""
		if err == nil {
			mandateFiles := []string{
				"AGENTS.md",
				"CLAUDE.md",
				"GEMINI.md",
				"COPILOT.md",
				"CONVENTIONS.md",
				".cursorrules",
				".windsurfrules",
				filepath.Join(".github", "copilot-instructions.md"),
			}
			loadedCount := 0
			for _, mf := range mandateFiles {
				mfPath := filepath.Join(root, mf)
				if content, readErr := os.ReadFile(mfPath); readErr == nil {
					agentsMandates += fmt.Sprintf("# MANDATOS DO WORKSPACE (%s)\nAs regras a seguir foram carregadas automaticamente do arquivo %s e são OBRIGATÓRIAS:\n\n%s\n\n---\n\n", mf, mf, string(content))
					fmt.Printf("\033[38;5;99m│\033[0m  📜 %s carregado (%d bytes) como instrução mandatória.\n", mf, len(content))
					loadedCount++
				}
			}
			if loadedCount == 0 {
				fmt.Printf("\033[38;5;99m│\033[0m  ⚠️  Nenhum arquivo de mandato encontrado (AGENTS.md, CLAUDE.md, etc.).\n")
			}
		}

		systemInstructions := agentsMandates
		systemInstructions += "Você é um Harness AI Agent. Você tem acesso às seguintes ferramentas de console em formato XML:\n"
		systemInstructions += "- Ler Arquivo: <tool:read_file path=\"caminho/relativo\"/>\n"
		systemInstructions += "- Escrever/Sobrescrever: <tool:write_file path=\"caminho/relativo\">conteudo</tool:write_file>\n"</tool:write_file>\n"
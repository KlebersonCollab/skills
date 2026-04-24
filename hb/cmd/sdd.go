package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var sddCmd = &cobra.Command{
	Use:   "sdd",
	Short: "Gerencia o workflow Spec-Driven Development (SDD)",
}

var sddInitCmd = &cobra.Command{
	Use:   "init [nome-da-feature]",
	Short: "Inicializa uma nova feature com estrutura completa SDD",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		name := args[0]
		slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
		dir := filepath.Join(root, ".specs", "features", slug)

		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			color.Red("❌ Erro: Feature '%s' já existe em %s", name, dir)
			os.Exit(1)
		}

		os.MkdirAll(dir, 0755)

		// Templates de Feature (Rigorosos)
		files := map[string]string{
			"spec.md":                fmt.Sprintf("# Specification: %s\n\n## Goal\n\n## Acceptance Criteria (BDD)\n", name),
			"plan.md":                fmt.Sprintf("# Technical Plan: %s\n\n## Architecture\n\n## Mermaid Diagrams\n", name),
			"tasks.md":               fmt.Sprintf("# Tasks: %s\n\n- [ ] 1. Define specifications\n- [ ] 2. Technical design\n- [ ] 3. Implementation\n- [ ] 4. Verification\n", name),
			"contract.md":            fmt.Sprintf("# Validation Contract: %s\n\n## Validation Sensors\n- [ ] Sensor 1\n", name),
			"validation-report.md":   fmt.Sprintf("# Validation Report: %s\n\n## Results\n\n## Evidence\n", name),
			"verification-report.md": fmt.Sprintf("# Quality Verification: %s\n\n## Checklist\n- [ ] Clean Code\n- [ ] Security\n", name),
		}

		for filename, content := range files {
			path := filepath.Join(dir, filename)
			os.WriteFile(path, []byte(content), 0644)
		}

		color.Green("✨ Feature '%s' inicializada com estrutura completa em %s", name, dir)
	},
}

var sddBootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Inicializa a estrutura global de governança do projeto (.specs/project, architecture, codebase)",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// Pastas Globais
		dirs := []string{
			filepath.Join(root, ".specs", "project"),
			filepath.Join(root, ".specs", "architecture"),
			filepath.Join(root, ".specs", "codebase"),
			filepath.Join(root, ".specs", "features"),
		}

		for _, d := range dirs {
			os.MkdirAll(d, 0755)
		}

		// Arquivos de Projeto (Tríade de Memória e Governança)
		projectFiles := map[string]string{
			".specs/project/STATE.md":       "# Project State\n\n## Current Focus\n\n## Roadmap Progress\n",
			".specs/project/MEMORY.md":      "# Project Memory\n\n## Key Contexts\n\n## Decisions History\n",
			".specs/project/LEARNINGS.md":    "# Project Learnings\n\n## Technical Gotchas\n\n## Best Practices\n",
			".specs/project/PROJECT.md":     "# Project Overview\n\n## Vision\n\n## Stakeholders\n",
			".specs/project/ROADMAP.md":     "# Project Roadmap\n\n## Q1\n\n## Q2\n",
			".specs/codebase/STACK.md":      "# Technology Stack\n\n- Language: \n- Framework: \n",
			".specs/codebase/GLOBAL_MANDATES.md": "# Global Mandates\n\n1. Use SDD\n2. Use Clean Code\n",
			".specs/architecture/README.md": "# Architecture Governance\n\nStore ADRs here.\n",
		}

		for relPath, content := range projectFiles {
			absPath := filepath.Join(root, relPath)
			if _, err := os.Stat(absPath); os.IsNotExist(err) {
				os.WriteFile(absPath, []byte(content), 0644)
				color.Cyan("   📄 Criado: %s", relPath)
			}
		}

		color.Green("🚀 Estrutura global de governança inicializada com sucesso!")
	},
}

var sddStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Mostra o progresso de todas as features em desenvolvimento",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		specsDir := filepath.Join(root, ".specs", "features")
		entries, err := os.ReadDir(specsDir)
		if err != nil {
			color.Yellow("⚠️  Nenhuma feature encontrada em %s", specsDir)
			return
		}

		color.Blue("📊 Estado do Projeto (SDD):")
		fmt.Printf("%-30s %-15s %-15s\n", "FEATURE", "PROGRESSO", "STATUS")
		fmt.Println(strings.Repeat("-", 65))

		re := regexp.MustCompile(`- \[(.)\]`)

		for _, entry := range entries {
			if entry.IsDir() {
				taskFile := filepath.Join(specsDir, entry.Name(), "tasks.md")
				content, err := os.ReadFile(taskFile)
				if err != nil {
					continue
				}

				matches := re.FindAllStringSubmatch(string(content), -1)
				total := len(matches)
				done := 0
				for _, m := range matches {
					if strings.ToLower(m[1]) == "x" {
						done++
					}
				}

				percent := 0.0
				if total > 0 {
					percent = float64(done) / float64(total) * 100
				}

				status := color.YellowString("Em progresso")
				if percent == 100 && total > 0 {
					status = color.GreenString("Concluída")
				} else if total == 0 {
					status = color.RedString("Sem tarefas")
				}

				fmt.Printf("%-30s %-15s %-15s\n", entry.Name(), fmt.Sprintf("%d/%d (%.0f%%)", done, total, percent), status)
			}
		}
	},
}

func init() {
	sddCmd.AddCommand(sddInitCmd)
	sddCmd.AddCommand(sddStatusCmd)
	sddCmd.AddCommand(sddBootstrapCmd)
	rootCmd.AddCommand(sddCmd)
}

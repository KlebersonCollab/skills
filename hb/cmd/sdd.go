package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/sdd"
	"github.com/klebersonromero/hb/internal/ui"
	"github.com/spf13/cobra"
)

var uiEnabled bool
var watchEnabled bool

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
			"spec.md": fmt.Sprintf("# Specification: %s\n\n## Goal\n\n## Acceptance Criteria (BDD)\n- **Given** \n- **When** \n- **Then** \n", name),
			"plan.md": fmt.Sprintf("# Technical Plan: %s\n\n## Architecture\n\n## Mermaid Diagrams\n```mermaid\ngraph TD\n    A --> B\n```\n", name),
			"tasks.md": fmt.Sprintf("# Tasks: %s\n\n## Phase 1: Planning\n- [ ] T1: Define specifications. `id: TASK-001`\n- [ ] T2: Technical design. `id: TASK-002`\n\n## Phase 2: Implementation\n- [ ] T3: Core logic. `id: TASK-003`\n\n## Phase 3: Verification\n- [ ] T4: Quality audit. `id: TASK-004`\n", name),
			"contract.md": fmt.Sprintf("# Validation Contract: %s\n\n## Validation Sensors\n- [ ] Sensor: Test Coverage > 80%%\n- [ ] Sensor: Lint Passing\n- [ ] Sensor: BDD Scenarios Validated\n", name),
			"validation-report.md": fmt.Sprintf("# Validation Report: %s\n\n## Results\n- **Score**: 0/100\n- **Status**: Pending\n\n## Evidence\n- [ ] Evidence 1\n", name),
			"verification-report.md": fmt.Sprintf("# Quality Verification: %s\n\n## Checklist\n- [ ] Clean Code (SOLID, DRY)\n- [ ] Security Audit\n- [ ] Performance Baseline\n", name),
		}

		for filename, content := range files {
			path := filepath.Join(dir, filename)
			os.WriteFile(path, []byte(content), 0644)
		}

		color.Green("✨ Feature '%s' inicializada com estrutura completa em %s", name, dir)
		color.Cyan("💡 Arquivos padrão gerados, pode editá-los conforme nosso padrão.")
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

		totalDoneGlobal := 0
		totalTasksGlobal := 0

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

				totalDoneGlobal += done
				totalTasksGlobal += total

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

		if uiEnabled {
			manager := ui.NewUIManager(os.Stdout)
			manager.Start()
			defer manager.Stop()

			percentGlobal := 0.0
			if totalTasksGlobal > 0 {
				percentGlobal = float64(totalDoneGlobal) / float64(totalTasksGlobal)
			}

			manager.SetState(ui.UIState{
				FeatureName: "ALL FEATURES",
				CurrentTask: fmt.Sprintf("Total Progress: %d/%d tasks", totalDoneGlobal, totalTasksGlobal),
				Progress:    percentGlobal,
				Status:      "Monitoring",
			})

			if watchEnabled {
				color.Cyan("\n👀 Watch mode enabled. Press Ctrl+C to exit.")
				for {
					time.Sleep(1 * time.Second)
				}
			}
			time.Sleep(1 * time.Second)
		}
	},
}

var sddTaskCmd = &cobra.Command{
	Use:   "task [feature] [task-id]",
	Short: "Marca uma tarefa como concluída no arquivo tasks.md da feature",
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		feature := ""
		taskID := ""

		if len(args) == 0 {
			// Interactive Mode
			var err error
			feature, err = sdd.SelectFeatureInteractively(root)
			if err != nil {
				color.Red("❌ %v", err)
				os.Exit(1)
			}
			taskID, err = sdd.SelectTaskInteractively(root, feature)
			if err != nil {
				color.Red("❌ %v", err)
				os.Exit(1)
			}
		} else {
			feature = args[0]
			if len(args) > 1 {
				taskID = args[1]
			}
		}

		specsDir := filepath.Join(root, ".specs", "features", feature)
		taskFile := filepath.Join(specsDir, "tasks.md")

		if _, err := os.Stat(taskFile); os.IsNotExist(err) {
			// Tenta buscar por slug se não encontrar direto
			feature = strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
			specsDir = filepath.Join(root, ".specs", "features", feature)
			taskFile = filepath.Join(specsDir, "tasks.md")

			if _, err := os.Stat(taskFile); os.IsNotExist(err) {
				color.Red("❌ Erro: Arquivo tasks.md não encontrado para a feature '%s'", feature)
				os.Exit(1)
			}
		}

		content, err := os.ReadFile(taskFile)
		if err != nil {
			color.Red("❌ Erro ao ler arquivo: %v", err)
			os.Exit(1)
		}

		lines := strings.Split(string(content), "\n")
		updated := false

		// Se tiver ID, procura a linha que começa com esse ID
		// Se não tiver ID, mas tiver descrição, tenta bater
		// Aqui vamos simplificar: se tiver ID numérico (ex: "1"), procura "- [ ] 1."
		re := regexp.MustCompile(`^- \[ \] ` + regexp.QuoteMeta(taskID) + `\.?`)

		for i, line := range lines {
			if re.MatchString(line) {
				lines[i] = strings.Replace(line, "[ ]", "[x]", 1)
				updated = true
				break
			}
		}

		if !updated {
			color.Yellow("⚠️  Tarefa '%s' não encontrada ou já concluída.", taskID)
			return
		}

		err = os.WriteFile(taskFile, []byte(strings.Join(lines, "\n")), 0644)
		if err != nil {
			color.Red("❌ Erro ao salvar arquivo: %v", err)
			os.Exit(1)
		}

		color.Green("✅ Tarefa '%s' marcada como concluída na feature '%s'!", taskID, feature)
	},
}

var sddAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audita a conformidade SDD e o progresso de todas as features",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		results, err := sdd.AuditFeatures(root)
		if err != nil {
			color.Red("❌ Erro ao auditar features: %v", err)
			os.Exit(1)
		}

		sdd.PrintAuditReport(results)
	},
}

var sddReconcileCmd = &cobra.Command{
	Use:   "reconcile [feature]",
	Short: "Sincroniza os artefatos SDD com a implementação real (Preenchimento Reverso)",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		if len(args) > 0 {
			feature := args[0]
			err := sdd.ReconcileFeature(root, feature)
			if err != nil {
				color.Red("❌ Erro ao reconciliar: %v", err)
				os.Exit(1)
			}
		} else {
			// Reconcilia todas as features com falhas ou progresso baixo
			results, _ := sdd.AuditFeatures(root)
			for _, s := range results {
				if !s.Contract || s.Progress < 100 {
					sdd.ReconcileFeature(root, s.Name)
				}
			}
		}
	},
}

var sddReviewCmd = &cobra.Command{
	Use:   "review [feature]",
	Short: "Verifica se a feature cumpre todos os requisitos do SDD",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		featureDir := filepath.Join(root, ".specs", "features", feature)
		artifacts := []string{"spec.md", "plan.md", "tasks.md", "contract.md", "validation-report.md"}
		
		color.Blue("🔍 Revisando Feature: %s", feature)
		allFound := true
		for _, art := range artifacts {
			path := filepath.Join(featureDir, art)
			if _, err := os.Stat(path); err == nil {
				color.Green("   ✅ %s encontrado.", art)
			} else {
				color.Red("   ❌ %s faltando!", art)
				allFound = false
			}
		}

		if allFound {
			color.Green("\n✨ Feature '%s' está em conformidade com o SDD!", feature)
		} else {
			color.Yellow("\n⚠️  Feature incompleta. Finalize os artefatos antes da implementação.")
		}
	},
}

var sddStartCmd = &cobra.Command{
	Use:   "start [nome-da-feature]",
	Short: "Inicializa uma nova feature com estrutura completa SDD",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		name := args[0]
		err := sdd.StartFeature(root, name)
		if err != nil {
			color.Red("❌ Erro ao iniciar feature: %v", err)
			os.Exit(1)
		}

		color.Green("✨ Feature '%s' inicializada com sucesso!", name)
	},
}

var sddContractCmd = &cobra.Command{
	Use:   "contract [feature]",
	Short: "Gerencia o contrato de validação da feature",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		check, _ := cmd.Flags().GetBool("check")
		gen, _ := cmd.Flags().GetBool("generate")

		if gen {
			err := sdd.GenerateContract(root, feature)
			if err != nil {
				color.Red("❌ Erro ao gerar contrato: %v", err)
				os.Exit(1)
			}
			return
		}

		if check {
			err := sdd.CheckContract(root, feature)
			if err != nil {
				color.Red("❌ Erro ao auditar contrato: %v", err)
				os.Exit(1)
			}
			return
		}

		// Comportamento padrão: apenas mostra o status do contrato
		err := sdd.CheckContract(root, feature)
		if err != nil {
			color.Red("❌ Erro: %v", err)
			os.Exit(1)
		}
	},
}

var sddUltraplanCmd = &cobra.Command{
	Use:   "ultraplan [feature]",
	Short: "Sugere um Teleporte de Exploração para tarefas de alta complexidade",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		color.Magenta("🚀 Iniciando ULTRAPLAN para: %s", feature)
		fmt.Println("Gatilho detectado: Alta ambiguidade/complexidade.")
		fmt.Println("Ação: Desacoplando exploração para sessão remota.")
		color.Cyan("🔗 Link de Exploração: https://claude.ai/new (Mode: Deep Search)")
	},
}

func init() {
	sddStatusCmd.Flags().BoolVar(&uiEnabled, "ui", false, "Habilita a interface premium do terminal")
	sddStatusCmd.Flags().BoolVar(&watchEnabled, "watch", false, "Monitora as mudanças em tempo real")

	sddCmd.AddCommand(sddUltraplanCmd)
	sddCmd.AddCommand(sddInitCmd) // Mantendo compatibilidade
	sddCmd.AddCommand(sddStartCmd)
	sddCmd.AddCommand(sddStatusCmd)
	sddCmd.AddCommand(sddBootstrapCmd)
	sddCmd.AddCommand(sddTaskCmd)
	sddCmd.AddCommand(sddReviewCmd)
	sddCmd.AddCommand(sddAuditCmd)
	sddCmd.AddCommand(sddReconcileCmd)
	
	sddContractCmd.Flags().BoolP("check", "c", false, "Audita a conformidade dos sensores")
	sddContractCmd.Flags().BoolP("generate", "g", false, "Gera o contrato baseado na spec.md")
	sddCmd.AddCommand(sddContractCmd)
	
	rootCmd.AddCommand(sddCmd)
}

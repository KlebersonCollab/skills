// Package agent provides the full SDD v2.3.0 framework with sub-skills:
// - sdd-planner (planner): STATE.md, MEMORY.md, LEARNINGS.md, context
// - sdd-orchestrator (orchestrator): spec.md, plan.md, tasks.md, contract.md
// - Phase 0: ALIGN (Grilling Session): CONTEXT.md, ADRs
package agent

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// SDDPhase represents a phase in the SDD workflow.
type SDDPhase string

const (
	PhaseAlign     SDDPhase = "ALIGN"
	PhaseDiscovery SDDPhase = "DISCOVERY"
	PhaseSpecify   SDDPhase = "SPECIFY"
	PhaseImplement SDDPhase = "IMPLEMENT"
	PhaseVerify    SDDPhase = "VERIFY"
)

// ============================================================================
// sdd-planner: Project Vision & Session Memory
// ============================================================================

// PlannerUpdate holds changes to project memory files.
type PlannerUpdate struct {
	FeatureName string
	Phase       SDDPhase
	Goal        string
	Decisions   []string
}

// UpdateProjectState updates STATE.md, MEMORY.md, LEARNINGS.md with current session state.
func UpdateProjectState(root string, update PlannerUpdate) error {
	// STATE.md
	statePath := filepath.Join(root, ".specs", "project", "STATE.md")
	stateContent := fmt.Sprintf(`# Project State

## Current Phase: [%s] — Feature: %s

## Active Goals
- %s

## Progress Tracking
- [ ] Feature: %s (In Progress).

---

<!-- @sdd-state -->
%s
`, update.Phase, update.FeatureName, update.Goal, update.FeatureName,
		fmt.Sprintf("```yaml\nversion: \"2.3.0\"\nfeature_id: \"%s\"\nphase: \"%s\"\nstatus: \"IN_PROGRESS\"\nlast_update: \"%s\"\n```",
			update.FeatureName, string(update.Phase), time.Now().UTC().Format(time.RFC3339)))

	if err := os.WriteFile(statePath, []byte(stateContent), 0644); err != nil {
		return fmt.Errorf("failed to write STATE.md: %w", err)
	}

	// MEMORY.md — append if not exists
	memoryPath := filepath.Join(root, ".specs", "project", "MEMORY.md")
	if _, err := os.Stat(memoryPath); os.IsNotExist(err) {
		memoryContent := fmt.Sprintf(`# Project Memory

## Project Identity
- **Name**: Skills Hub
- **Purpose**: Harness AI Agent

## Context Facts
- SDD v2.3.0 governance ativo.
- Feature atual: %s

---

<!-- @sdd-state -->
%s
`, update.FeatureName,
			"```yaml\nversion: \"2.3.0\"\nfeature_id: \"PROJECT-MEMORY\"\nphase: \"ALIGN\"\nstatus: \"COMPLETED\"\n```")
		os.WriteFile(memoryPath, []byte(memoryContent), 0644)
	}

	// LEARNINGS.md
	learningsPath := filepath.Join(root, ".specs", "project", "LEARNINGS.md")
	if _, err := os.Stat(learningsPath); os.IsNotExist(err) {
		learningsContent := `# Project Learnings

## Technical Patterns
- SDD v2.3.0 Integration

---

<!-- @sdd-state -->
` + "```yaml\nversion: \"2.3.0\"\nfeature_id: \"LEARNINGS\"\nphase: \"ALIGN\"\nstatus: \"COMPLETED\"\n```\n"
		os.WriteFile(learningsPath, []byte(learningsContent), 0644)
	}

	return nil
}

// ============================================================================
// Utils: Feature name extraction, directory management
// ============================================================================

// DetectDevTask checks if input appears to be a development task requiring SDD.
func DetectDevTask(input string) bool {
	lower := strings.ToLower(input)
	keywords := []string{
		"implement", "criar", "crie", "desenvolva", "desenvolver",
		"refator", "refactor", "feature", "funcionalidade",
		"adicionar", "adiciona", "add",
		"corrigir", "corrige", "fix", "bug",
		"modificar", "modifica", "change",
		"construir", "build",
	}
	for _, kw := range keywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}

// ExtractFeatureName converts task text to a kebab-case feature name.
func ExtractFeatureName(input string) string {
	cleaned := input
	prefixes := []string{
		"implementar ", "implement ", "criar ", "crie ",
		"desenvolver ", "desenvolva ", "refatorar ", "refactor ",
		"adicione ", "adicionar ", "add ", "corrigir ", "corrige ",
		"fix ", "criar funcionalidade ",
	}
	for _, p := range prefixes {
		if strings.HasPrefix(strings.ToLower(cleaned), p) {
			cleaned = strings.TrimSpace(cleaned[len(p):])
			break
		}
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9\s-]`)
	cleaned = re.ReplaceAllString(cleaned, "")
	cleaned = strings.TrimSpace(cleaned)
	cleaned = strings.ReplaceAll(cleaned, " ", "-")
	cleaned = strings.ToLower(cleaned)
	if len(cleaned) > 50 {
		cleaned = cleaned[:50]
	}
	cleaned = strings.TrimRight(cleaned, "-")
	if cleaned == "" {
		return "unnamed-feature"
	}
	return cleaned
}

// FeatureDir returns the standard feature directory path.
func FeatureDir(root string, featureName string) string {
	return filepath.Join(root, ".specs", "features", featureName)
}

// EnsureFeatureDir creates the feature directory structure.
func EnsureFeatureDir(root string, featureName string) error {
	return os.MkdirAll(FeatureDir(root, featureName), 0755)
}

// FeatureExists checks if spec.md already exists for this feature.
func FeatureExists(root string, featureName string) bool {
	_, err := os.Stat(filepath.Join(FeatureDir(root, featureName), "spec.md"))
	return err == nil
}

// ============================================================================
// Phase 0: ALIGN — Grilling Session (sabatina)
// ============================================================================

// RunGrillingSession executes the Phase 0: ALIGN process.
// Follows the Grilling Session reference protocol:
// 1. Glossary Enforcement — challenge vague terms
// 2. Scenario-Driven Stress Testing — edge-case probing
// 3. ADR creation for hard-to-reverse decisions
func RunGrillingSession(root string, featureName string, task string, reader *bufio.Reader) ([]string, error) {
	fmt.Println("\n\033[38;5;99m┌── 🔥 SDD PHASE 0: ALIGN — GRILLING SESSION ──────────────────────────────┐\033[0m")
	fmt.Println("\033[38;5;99m│\033[0m  🎯  Sabatina de alinhamento de terminologia e arquitetura.")
	fmt.Println("\033[38;5;99m│\033[0m  📝  Responda cada pergunta. Decisoes viram ADRs.")
	fmt.Printf("\033[38;5;99m│\033[0m  🔧  Feature: \033[1;36m%s\033[0m\n", featureName)
	fmt.Printf("\033[38;5;99m│\033[0m  📋  Task: \033[1;33m%s\033[0m\n", task)
	fmt.Println("\033[38;5;99m└──────────────────────────────────────────────────────────────────────┘\033[0m")

	// Load existing CONTEXT.md glossary
	contextPath := filepath.Join(root, ".specs", "project", "CONTEXT.md")
	if _, err := os.ReadFile(contextPath); err == nil {
		fmt.Printf("\033[38;5;99m│\033[0m  📖 Glossario existente carregado\n")
	}

	decisions := []string{}

	// --- Question 1: Objective ---
	fmt.Printf("\n\033[38;5;99m│\033[0m  \033[1;36m❓ 1/5 — Qual o objetivo principal? (1-2 frases)\033[0m\n\033[38;5;99m❯\033[0m ")
	answer1, _ := reader.ReadString('\n')
	answer1 = strings.TrimSpace(answer1)
	if answer1 == "" {
		answer1 = task
	}
	decisions = append(decisions, fmt.Sprintf("Objetivo: %s", answer1))

	// --- Question 2: Domain Terminology ---
	fmt.Printf("\n\033[38;5;99m│\033[0m  \033[1;36m❓ 2/5 — Que termos de dominio sao importantes? (Ex: \"cliente\" significa usuario ou consumidor?)\033[0m\n\033[38;5;99m❯\033[0m ")
	answer2, _ := reader.ReadString('\n')
	answer2 = strings.TrimSpace(answer2)
	if answer2 != "" {
		decisions = append(decisions, fmt.Sprintf("Terminologia: %s", answer2))
		// Update CONTEXT.md with new terms
		updateGlossary(contextPath, answer2)
	}

	// --- Question 3: Alternatives ---
	fmt.Printf("\n\033[38;5;99m│\033[0m  \033[1;36m❓ 3/5 — Que alternativas voce considera? (Ou \"nenhuma\")\033[0m\n\033[38;5;99m❯\033[0m ")
	answer3, _ := reader.ReadString('\n')
	answer3 = strings.TrimSpace(answer3)
	if answer3 != "" && !strings.EqualFold(answer3, "nenhuma") {
		decisions = append(decisions, fmt.Sprintf("Alternativas consideradas: %s", answer3))
	}

	// --- Question 4: Risks ---
	fmt.Printf("\n\033[38;5;99m│\033[0m  \033[1;36m❓ 4/5 — Ha riscos ou restricoes? (performance, seguranca, compatibilidade)\033[0m\n\033[38;5;99m❯\033[0m ")
	answer4, _ := reader.ReadString('\n')
	answer4 = strings.TrimSpace(answer4)
	if answer4 != "" {
		decisions = append(decisions, fmt.Sprintf("Riscos: %s", answer4))
	}

	// --- Question 5: ADR Check ---
	fmt.Printf("\n\033[38;5;99m│\033[0m  \033[1;36m❓ 5/5 — Esta decisao envolve trade-offs reais? (S/n)\033[0m\n\033[38;5;99m❯\033[0m ")
	answer5, _ := reader.ReadString('\n')
	answer5 = strings.TrimSpace(strings.ToLower(answer5))

	if answer5 == "" || answer5 == "s" || answer5 == "sim" {
		adrContent := generateADRContent(featureName, task, decisions)
		adrDir := filepath.Join(root, ".specs", "architecture")
		os.MkdirAll(adrDir, 0755)

		nextNum := 1
		entries, _ := os.ReadDir(adrDir)
		for _, e := range entries {
			var num int
			fmt.Sscanf(e.Name(), "%d-", &num)
			if num >= nextNum {
				nextNum = num + 1
			}
		}

		adrFilename := fmt.Sprintf("%04d-%s.md", nextNum, featureName)
		adrPath := filepath.Join(adrDir, adrFilename)
		os.WriteFile(adrPath, []byte(adrContent), 0644)
		decisions = append(decisions, fmt.Sprintf("[ADR] %s", adrFilename))
		fmt.Printf("\033[38;5;99m│\033[0m  🏛️  ADR criado: \033[1;32m%s\033[0m\n", adrPath)
	}

	fmt.Println("\n\033[38;5;99m┌── ✅ GRILLING CONCLUIDA ────────────────────────────────────────────────┐\033[0m")
	fmt.Printf("\033[38;5;99m│\033[0m  📋 %d decisoes registradas\n", len(decisions))
	for _, d := range decisions {
		fmt.Printf("\033[38;5;99m│\033[0m    • %s\n", truncateStr(d, 70))
	}
	fmt.Println("\033[38;5;99m└──────────────────────────────────────────────────────────────────────┘\033[0m")

	return decisions, nil
}

func generateADRContent(featureName string, task string, decisions []string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("# ADR: Decisoes Arquiteturais — %s\n\n", featureName))
	b.WriteString(fmt.Sprintf("Contexto: %s\n\n", task))
	b.WriteString("## Status\nAccepted\n\n## Decisoes\n\n")
	for _, d := range decisions {
		b.WriteString(fmt.Sprintf("- %s\n", d))
	}
	b.WriteString(fmt.Sprintf("\n## Data\n%s\n\n## Consequencias\n", time.Now().Format(time.RFC3339)))
	b.WriteString("Decisoes validadas via Grilling Session (Phase 0: ALIGN).\n")
	return b.String()
}

func updateGlossary(contextPath string, newTerms string) {
	dir := filepath.Dir(contextPath)
	os.MkdirAll(dir, 0755)

	var existing []byte
	if data, err := os.ReadFile(contextPath); err == nil {
		existing = data
	}

	// Simple append: add new terms section if not exists
	content := string(existing)
	if content == "" {
		content = fmt.Sprintf(`# Context - Domain Glossary

This context defines precise terminology for this project.

## Language

**%s**:
_Term identified during Grilling Session._
_Avoid_: use termino consistente

---

<!-- @sdd-state -->
%s
`, newTerms,
			"```yaml\nversion: \"2.3.0\"\nfeature_id: \"CONTEXT\"\nphase: \"ALIGN\"\nstatus: \"IN_PROGRESS\"\n```")
	} else {
		content += fmt.Sprintf("\n**Novo termo**: %s\n", newTerms)
	}

	os.WriteFile(contextPath, []byte(content), 0644)
}

// ============================================================================
// sdd-orchestrator: Specification, Plan, Tasks, Contract
// ============================================================================

// CreateSDDArtifacts generates all SDD artifacts following official templates.
func CreateSDDArtifacts(root string, featureName string, task string, decisions []string) error {
	dir := FeatureDir(root, featureName)
	if err := EnsureFeatureDir(root, featureName); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// spec.md — following spec-template.md
	spec := fmt.Sprintf(`# %s — Specification

## 1. Context & Goals
- **Problem Statement**: %s
- **Scope**: Implementar conforme definido na Grilling Session.
- **Success Criteria**: Testes passando, build limpo.

## 2. Requirements (BDD Scenarios)

### Feature: %s

**Scenario 1: Implementacao Basica**
- **Given** o contexto definido
- **When** a feature for implementada
- **Then** deve funcionar conforme o esperado

**Scenario 2: Tratamento de Erros**
- **Given** uma condicao de erro
- **When** o sistema encontrar o erro
- **Then** deve tratar graciosamente

## 3. Constraints & Risks
- **Dependencies**: Nenhuma alem do ecossistema existente

---

<!-- @sdd-state -->
%s
`, featureName, task, featureName,
		fmt.Sprintf("```yaml\nversion: \"2.3.0\"\nfeature_id: \"%s\"\nphase: \"DISCOVERY\"\nstatus: \"IN_PROGRESS\"\n```", featureName))

	specPath := filepath.Join(dir, "spec.md")
	os.WriteFile(specPath, []byte(spec), 0644)
	fmt.Printf("\033[38;5;99m│\033[0m  📄 \033[1;36mspec.md\033[0m → %s\n", specPath)

	// plan.md — following plan-template.md
	plan := fmt.Sprintf(`# %s — Technical Plan

## 1. Architecture Overview
Implementacao direta seguindo principios SOLID e SDD v2.3.0.

## 2. Technical Design

### Logic Flow (Mermaid)
%s

### Implementation Details
- Isolamento: Modulos especificos da feature
- Testes: Unitarios + Integracao

## 3. Implementation Strategy
- **Testing Strategy**: Testes unitarios obrigatorios
- **Migrations**: Nao aplicavel

---

<!-- @sdd-state -->
%s
`, featureName,
		"```mermaid\ngraph TD\n    A[Requisito] --> B[Implementacao]\n    B --> C[Testes]\n    C --> D[Validacao]\n    D --> E[Entrega]\n```",
		fmt.Sprintf("```yaml\nversion: \"2.3.0\"\nfeature_id: \"%s\"\nphase: \"SPECIFY\"\nstatus: \"IN_PROGRESS\"\n```", featureName))

	planPath := filepath.Join(dir, "plan.md")
	os.WriteFile(planPath, []byte(plan), 0644)
	fmt.Printf("\033[38;5;99m│\033[0m  📄 \033[1;36mplan.md\033[0m → %s\n", planPath)

	// tasks.md — following tasks-template.md (with Evidence column)
	tasks := fmt.Sprintf(`# %s — Tasks

## 📊 Phase Progress Monitoring

| Phase | Task | Status | Evidence (Commit/Log) |
| :--- | :--- | :---: | :--- |
| **1. PREP** | Setup e analise | [ ] | |
| **2. CORE** | Implementar funcionalidade principal | [ ] | |
| **2. CORE** | Testes unitarios | [ ] | |
| **3. FINAL** | Integracao e validacao | [ ] | |
| **3. FINAL** | Documentacao e cleanup | [ ] | |

---

<!-- @sdd-state -->
%s
`, featureName,
		fmt.Sprintf("```yaml\nversion: \"2.3.0\"\nfeature_id: \"%s\"\nphase: \"SPECIFY\"\nstatus: \"IN_PROGRESS\"\n```", featureName))

	tasksPath := filepath.Join(dir, "tasks.md")
	os.WriteFile(tasksPath, []byte(tasks), 0644)
	fmt.Printf("\033[38;5;99m│\033[0m  📄 \033[1;36mtasks.md\033[0m → %s\n", tasksPath)

	// contract.md — validation sensors
	contract := fmt.Sprintf(`# %s — Contract (SDC)

## Deliverables
- Feature implementada conforme spec.md
- Testes passando
- Build limpo

## Sensors
- **Build**: go build ./... → exit 0
- **Testes**: go test ./... -count=1 → all pass
- **Lint**: go vet ./... → clean

## Constraints Verification
- Arquitetura segue plan.md
- Terminologia segue CONTEXT.md
- ADRs respeitados

---

<!-- @sdd-state -->
%s
`, featureName,
		fmt.Sprintf("```yaml\nversion: \"2.3.0\"\nfeature_id: \"%s\"\nphase: \"SPECIFY\"\nstatus: \"IN_PROGRESS\"\n```", featureName))

	contractPath := filepath.Join(dir, "contract.md")
	os.WriteFile(contractPath, []byte(contract), 0644)
	fmt.Printf("\033[38;5;99m│\033[0m  📄 \033[1;36mcontract.md\033[0m → %s\n", contractPath)

	return nil
}

// ============================================================================
// Helpers
// ============================================================================

func truncateStr(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

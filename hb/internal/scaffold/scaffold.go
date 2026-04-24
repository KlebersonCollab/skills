package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

const skillMdTemplate = `---
name: %s
version: 1.0.0
description: "Brief description of the skill."
category: general
uses:
  - sdd
---

# %s (v1.0.0)

> Enter a powerful mission statement or summary of the skill here.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
1. **Context Check**: rehydrate via STATE.md, MEMORY.md, and LEARNINGS.md.
2. **Spec/Plan Check**: ensure artifacts exist.

---

## Goal
Explain the main goal of this skill.

## Workflow
1. Phase 1...
2. Phase 2...
`

const changelogTemplate = `# Changelog: %s

## [1.0.0] - %s
### Added
- Initial release of the %s skill.
`

// CreateSkill gera a estrutura de uma nova skill no padrão Ouro
func CreateSkill(root, name string) error {
	skillDir := filepath.Join(root, name)
	if _, err := os.Stat(skillDir); err == nil {
		return fmt.Errorf("skill '%s' já existe", name)
	}

	// Cria pastas obrigatórias
	dirs := []string{"examples", "scripts", "resources"}
	for _, d := range dirs {
		os.MkdirAll(filepath.Join(skillDir, d), 0755)
	}

	// 1. SKILL.md
	skillContent := fmt.Sprintf(skillMdTemplate, name, name)
	os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(skillContent), 0644)

	// 2. CHANGELOG.md
	date := time.Now().Format("2006-01-02")
	changelogContent := fmt.Sprintf(changelogTemplate, name, date, name)
	os.WriteFile(filepath.Join(skillDir, "CHANGELOG.md"), []byte(changelogContent), 0644)

	// 3. metadata.json
	metadata := fmt.Sprintf(`{
	"name": "%s",
	"version": "1.0.0",
	"description": "Expert skill for %s",
	"author": "AI Ecosystem",
	"dependencies": ["sdd"]
}`, name, name)
	os.WriteFile(filepath.Join(skillDir, "metadata.json"), []byte(metadata), 0644)

	// 4. README.md
	os.WriteFile(filepath.Join(skillDir, "README.md"), []byte("# "+name+"\n\nDocumentation for the "+name+" skill."), 0644)

	color.Green("✨ Skill '%s' scaffolded com estrutura EXPERT!", name)
	return nil
}

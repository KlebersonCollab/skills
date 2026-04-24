package fixer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/scanner"
)

// Templates para arquivos globais
var globalTemplates = map[string]string{
	".specs/project/STATE.md":     "# Current Project State\n\n## Current Status\n- [ ] Initializing project...\n",
	".specs/project/MEMORY.md":    "# Project Persistent Memory\n\n## Conventions\n- Add long-term patterns and decisions here.\n",
	".specs/project/LEARNINGS.md": "# Project Learnings\n\n## Lessons Learned\n- Log of learnings and avoided errors.\n",
	".specs/project/ROADMAP.md":   "# Project Roadmap\n\n## Next Steps\n- [ ] Define project objectives.\n",
}

// FixGlobalStructure garante que a pasta .specs/project/ e seus arquivos existam
func FixGlobalStructure(root string) error {
	for relPath, content := range globalTemplates {
		fullPath := filepath.Join(root, relPath)
		dir := filepath.Dir(fullPath)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			err = os.WriteFile(fullPath, []byte(content), 0644)
			if err != nil {
				return err
			}
			color.Cyan("   + Global repair: %s created.", relPath)
		}
	}
	return nil
}

// FixSkillsStructure garante que cada skill tenha seus arquivos de governança
func FixSkillsStructure(root string) error {
	skills, err := scanner.FindSkills(root)
	if err != nil {
		return err
	}

	requiredFiles := map[string]string{
		"tasks.md":     "# Tasks: %s\n\n## 🛠️ Fase 3: IMPLEMENT\n- [ ] Initial setup.\n",
		"MEMORY.md":    "# Memory: %s\n\n## Persistent Context\n- Store patterns here.\n",
		"LEARNINGS.md": "# Learnings: %s\n\n## Lessons Learned\n- Record technical insights.\n",
	}

	for _, s := range skills {
		for fileName, template := range requiredFiles {
			fullPath := filepath.Join(s.Path, fileName)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				content := fmt.Sprintf(template, s.Name)
				err = os.WriteFile(fullPath, []byte(content), 0644)
				if err != nil {
					color.Red("   ! Error creating %s in %s: %v", fileName, s.Name, err)
					continue
				}
				color.Cyan("   + Skill %s: %s repaired.", s.Name, fileName)
			}
		}
	}
	return nil
}

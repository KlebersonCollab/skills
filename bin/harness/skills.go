package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// SkillDef represents a loaded skill from any standard agent directory
type SkillDef struct {
	Name        string
	Description string
	Path        string
}

// DiscoverSkills scans the workspace for skills following market conventions.
// It looks in: root (as fallback), .agents/skills/, .gemini/skills/, .claude/skills/
func DiscoverSkills(root string) []SkillDef {
	var skills []SkillDef

	// Market standard skill directories
	searchDirs := []string{
		filepath.Join(root, ".agents", "skills"),
		filepath.Join(root, ".gemini", "skills"),
		filepath.Join(root, ".claude", "skills"),
		root, // Fallback for legacy / generic
	}

	seen := make(map[string]bool)

	for _, dir := range searchDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") && entry.Name() != "bin" && entry.Name() != "architecture" && entry.Name() != "docs" {
				skillName := entry.Name()
				if seen[skillName] {
					continue // Prevent duplicate loading if same skill exists in multiple folders
				}

				skillPath := filepath.Join(dir, entry.Name())
				skillMDPath := filepath.Join(skillPath, "SKILL.md")

				if _, err := os.Stat(skillMDPath); err == nil {
					description := extractSkillDescription(skillMDPath)
					skills = append(skills, SkillDef{
						Name:        skillName,
						Description: description,
						Path:        skillPath,
					})
					seen[skillName] = true
				}
			}
		}
	}

	return skills
}

// extractSkillDescription reads the YAML frontmatter of SKILL.md to extract the description
func extractSkillDescription(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "Sem descrição (falha ao ler o arquivo)"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inFrontmatter := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		if line == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				break // End of frontmatter
			}
		}

		if inFrontmatter {
			if strings.HasPrefix(line, "description:") {
				desc := strings.TrimPrefix(line, "description:")
				desc = strings.TrimSpace(desc)
				desc = strings.Trim(desc, `"'`) // Remove quotes
				return desc
			}
		}
	}

	return "Sem descrição (YAML frontmatter ausente ou sem o campo 'description')"
}

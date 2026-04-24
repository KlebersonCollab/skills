package scanner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/klebersonromero/hb/internal/domain"
	"gopkg.in/yaml.v3"
)

// Blacklist de diretórios técnicos a serem ignorados
var blacklist = map[string]bool{
	".git":         true,
	".venv":        true,
	".specs":       true,
	"node_modules": true,
	"artifacts":    true,
	".agent":       true,
	".agents":      true,
	"dist":         true,
	"hb":           true,
	"scripts":      true,
	"tests":        true,
	"dashboard":    true,
}

// FindSkills caminha pelo diretório raiz e retorna uma lista de skills detectadas
func FindSkills(root string) ([]domain.Skill, error) {
	var skills []domain.Skill

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if blacklist[name] {
			continue
		}

		path := filepath.Join(root, name)
		skillFile := filepath.Join(path, "SKILL.md")

		// Se tem SKILL.md, é uma skill válida para o hub
		if _, err := os.Stat(skillFile); err == nil {
			skill := domain.Skill{
				Name: name,
				Path: path,
			}
			// Parse metadata
			parseSkillMetadata(skillFile, &skill)
			skills = append(skills, skill)
		}
	}

	return skills, nil
}

func parseSkillMetadata(path string, skill *domain.Skill) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	// Extrair bloco entre --- e ---
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return
	}

	var meta struct {
		Name        string   `yaml:"name"`
		Version     string   `yaml:"version"`
		Category    string   `yaml:"category"`
		Description string   `yaml:"description"`
		Uses        []string `yaml:"uses"`
		References  []string `yaml:"references"`
	}

	err = yaml.Unmarshal([]byte(parts[1]), &meta)
	if err != nil {
		return
	}

	if meta.Version != "" {
		skill.Version = meta.Version
	}
	if meta.Category != "" {
		skill.Category = meta.Category
	}
	if meta.Description != "" {
		skill.Description = meta.Description
	}
	if len(meta.Uses) > 0 {
		skill.Uses = meta.Uses
	}
	if len(meta.References) > 0 {
		skill.References = meta.References
	}
	// Se o YAML tiver um nome explícito, usamos ele
	if meta.Name != "" {
		skill.Name = meta.Name
	}
}

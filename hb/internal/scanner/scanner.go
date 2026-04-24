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
	"dist":         true,
	"hb":           true,
	"scripts":      true,
	"tests":        true,
	"dashboard":    true,
}

// FindSkills caminha pelo diretório raiz e retorna uma lista de skills detectadas
func FindSkills(root string) ([]domain.Skill, error) {
	var skills []domain.Skill

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		// Ignorar diretórios na blacklist
		name := d.Name()
		if blacklist[name] {
			return filepath.SkipDir
		}

		skillFile := filepath.Join(path, "SKILL.md")
		if _, err := os.Stat(skillFile); err == nil {
			skill := domain.Skill{
				Name: name,
				Path: path,
			}
			parseSkillMetadata(skillFile, &skill)
			skills = append(skills, skill)
			
			// Se encontramos uma skill, não precisamos entrar nela (skills não são aninhadas)
			return filepath.SkipDir
		}

		return nil
	})

	return skills, err
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

package validator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/klebersonromero/hb/internal/domain"
	"gopkg.in/yaml.v3"
)

var requiredFiles = []string{"SKILL.md", "tasks.md", "MEMORY.md", "LEARNINGS.md"}

// ValidateSkill executa todas as validações para uma única skill
func ValidateSkill(skill domain.Skill) domain.AuditResult {
	result := domain.AuditResult{Skill: skill}

	// 1. Validação de Estrutura
	for _, file := range requiredFiles {
		fullPath := filepath.Join(skill.Path, file)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			result.Issues = append(result.Issues, domain.AuditIssue{
				File:    file,
				Message: "Arquivo obrigatório ausente",
				Fatal:   true,
			})
		}
	}

	// 2. Validação de Metadados (YAML Frontmatter)
	validateMetadata(&result)

	return result
}

func validateMetadata(result *domain.AuditResult) {
	skillFile := filepath.Join(result.Skill.Path, "SKILL.md")
	content, err := os.ReadFile(skillFile)
	if err != nil {
		result.Issues = append(result.Issues, domain.AuditIssue{
			File:    "SKILL.md",
			Message: "Erro ao ler arquivo para validar metadados",
			Fatal:   true,
		})
		return
	}

	// Extrair YAML entre --- e ---
	parts := strings.Split(string(content), "---")
	if len(parts) < 3 {
		result.Issues = append(result.Issues, domain.AuditIssue{
			File:    "SKILL.md",
			Message: "YAML Frontmatter não encontrado ou mal formatado",
			Fatal:   true,
		})
		return
	}

	var meta struct {
		Name        string `yaml:"name"`
		Version     string `yaml:"version"`
		Category    string `yaml:"category"`
		Description string `yaml:"description"`
	}

	err = yaml.Unmarshal([]byte(parts[1]), &meta)
	if err != nil {
		result.Issues = append(result.Issues, domain.AuditIssue{
			File:    "SKILL.md",
			Message: "Erro de sintaxe no YAML: " + err.Error(),
			Fatal:   true,
		})
		return
	}

	// Validar campos obrigatórios
	if meta.Name == "" {
		result.Issues = append(result.Issues, domain.AuditIssue{File: "SKILL.md", Message: "Campo 'name' ausente no YAML", Fatal: true})
	}
	if meta.Version == "" {
		result.Issues = append(result.Issues, domain.AuditIssue{File: "SKILL.md", Message: "Campo 'version' ausente no YAML", Fatal: true})
	}

	// Atualizar dados da skill no domínio para o relatório
	result.Skill.Name = meta.Name
	result.Skill.Version = meta.Version
	result.Skill.Category = meta.Category
}

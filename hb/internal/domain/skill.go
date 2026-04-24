package domain

// Skill representa a estrutura de uma skill no Hub
type Skill struct {
	Name        string
	Path        string
	Version     string
	Description string
	Category    string
	Uses        []string
	References  []string
}

// AuditIssue representa um erro ou aviso encontrado durante o audit
type AuditIssue struct {
	File    string
	Message string
	Fatal   bool
}

// AuditResult consolida o resultado da auditoria de uma skill
type AuditResult struct {
	Skill  Skill
	Issues []AuditIssue
}

// Success retorna verdadeiro se não houver erros fatais
func (r *AuditResult) Success() bool {
	for _, issue := range r.Issues {
		if issue.Fatal {
			return false
		}
	}
	return true
}

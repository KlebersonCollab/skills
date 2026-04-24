package mapper

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/klebersonromero/hb/internal/domain"
	"github.com/klebersonromero/hb/internal/scanner"
)

var categoryNames = map[string]string{
	"api-design":                   "Architecture & Design",
	"architecture":                 "Architecture & Design",
	"devops-automation":            "Ecosystems & DevOps",
	"site-reliability-engineering": "Ecosystems & DevOps",
	"development-python":           "Ecosystems & DevOps",
	"development":                  "Ecosystems & DevOps",
	"development-workflow":         "Core Frameworks",
	"automation":                   "Automation & Utils",
	"orchestration":                "Navigation & Orchestration",
	"design-thinking":              "Core Frameworks",
	"code-quality":                 "Architecture & Design",
	"project-management":           "Core Frameworks",
	"skill-management":             "Core Frameworks",
	"knowledge-management":          "Core Frameworks",
}

var styleMap = map[string]string{
	"Architecture & Design":      "fill:#f3e5f5,stroke:#4a148c", // Purple
	"Ecosystems & DevOps":        "fill:#e8f5e9,stroke:#1b5e20", // Green
	"Core Frameworks":            "fill:#e1f5fe,stroke:#01579b", // Blue
	"Automation & Utils":         "fill:#fff3e0,stroke:#e65100", // Orange
	"Navigation & Orchestration": "fill:#fce4ec,stroke:#880e4f", // Pink
}

func Generate(root string) error {
	skills, err := scanner.FindSkills(root)
	if err != nil {
		return err
	}

	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Name < skills[j].Name
	})

	var builder strings.Builder
	builder.WriteString("graph TD\n")
	builder.WriteString("    %% Nodes and Hierarchical Grouping\n")

	subgraphs := make(map[string][]domain.Skill)
	relations := make(map[string]bool)

	// 1. Processamento e Agrupamento
	for _, s := range skills {
		catName, exists := categoryNames[s.Category]
		if !exists {
			catName = "Miscellaneous"
		}
		subgraphs[catName] = append(subgraphs[catName], s)

		// Relações Explícitas (Uses)
		for _, dep := range s.Uses {
			depID := strings.ReplaceAll(dep, "-", "_")
			skillID := strings.ReplaceAll(filepath.Base(s.Path), "-", "_")
			relations[fmt.Sprintf("%s --(uses)--> %s", skillID, depID)] = true
		}

		// Relações Explícitas (References)
		for _, ref := range s.References {
			refID := strings.ReplaceAll(ref, "-", "_")
			skillID := strings.ReplaceAll(filepath.Base(s.Path), "-", "_")
			relations[fmt.Sprintf("%s --(references)--> %s", skillID, refID)] = true
		}
		
		// Relações Implícitas (Menções no SKILL.md)
		skillMD := filepath.Join(s.Path, "SKILL.md")
		content, _ := os.ReadFile(skillMD)
		contentStr := string(content)
		
		for _, other := range skills {
			otherBase := filepath.Base(other.Path)
			if otherBase != filepath.Base(s.Path) {
				re := regexp.MustCompile(`\b` + regexp.QuoteMeta(otherBase) + `\b`)
				if re.MatchString(contentStr) {
					skillID := strings.ReplaceAll(filepath.Base(s.Path), "-", "_")
					otherID := strings.ReplaceAll(otherBase, "-", "_")
					rel := fmt.Sprintf("%s --(references)--> %s", skillID, otherID)
					useRel := fmt.Sprintf("%s --(uses)--> %s", skillID, otherID)
					if !relations[useRel] {
						relations[rel] = true
					}
				}
			}
		}
	}

	// 2. Construção dos Subgrafos
	for catName, catSkills := range subgraphs {
		builder.WriteString(fmt.Sprintf("\n    subgraph \"%s\"\n", catName))
		for _, s := range catSkills {
			skillID := strings.ReplaceAll(filepath.Base(s.Path), "-", "_")
			
			// Badge SDD
			content, _ := os.ReadFile(filepath.Join(s.Path, "SKILL.md"))
			badge := ""
			if strings.Contains(string(content), "🔒 Prerequisites (Mandatory)") {
				badge = " 🛡️"
			}
			
			builder.WriteString(fmt.Sprintf("        %s[\"%s (%s)%s\"]\n", skillID, s.Name, s.Version, badge))
		}
		builder.WriteString("    end\n")
	}

	// 3. Relações
	builder.WriteString("\n    %% Relations\n")
	var relList []string
	for rel := range relations {
		// Se existe (uses), ignora (references) redundante
		if strings.Contains(rel, "--(references)-->") {
			base := strings.ReplaceAll(rel, "--(references)-->", "--(uses)-->")
			if relations[base] {
				continue
			}
		}
		relList = append(relList, rel)
	}
	sort.Strings(relList)
	for _, rel := range relList {
		builder.WriteString(fmt.Sprintf("    %s\n", rel))
	}

	// 4. Estilização
	builder.WriteString("\n    %% Styling\n")
	for catName, style := range styleMap {
		if skills, ok := subgraphs[catName]; ok {
			for _, s := range skills {
				skillID := strings.ReplaceAll(filepath.Base(s.Path), "-", "_")
				builder.WriteString(fmt.Sprintf("    style %s %s\n", skillID, style))
			}
		}
	}

	// 5. Legenda
	builder.WriteString("\n    %% Legend\n")
	builder.WriteString("    subgraph \"Legend\"\n")
	builder.WriteString("        L1[\"🛡️ SDD Compliant\"]\n")
	builder.WriteString("        L2[\"Uses: Direct Dependency\"]\n")
	builder.WriteString("        L3[\"References: Documentation Mention\"]\n")
	builder.WriteString("    end\n")

	outputPath := filepath.Join(root, "KNOWLEDGE-MAP.mermaid")
	return os.WriteFile(outputPath, []byte(builder.String()), 0644)
}

func AnalyzeImpact(root, target string) ([]string, error) {
	skills, err := scanner.FindSkills(root)
	if err != nil {
		return nil, err
	}

	impactedMap := make(map[string]bool)
	targetBase := filepath.Base(target)
	targetBase = strings.TrimSuffix(targetBase, filepath.Ext(targetBase))

	for _, s := range skills {
		skillMD := filepath.Join(s.Path, "SKILL.md")
		content, err := os.ReadFile(skillMD)
		if err != nil {
			continue
		}

		// Procura menções exatas ao nome da skill ou arquivo (case-insensitive)
		re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(targetBase) + `\b`)
		if re.Match(content) {
			impactedMap[s.Name] = true
		}
	}

	impacted := []string{}
	for name := range impactedMap {
		impacted = append(impacted, name)
	}
	sort.Strings(impacted)

	return impacted, nil
}

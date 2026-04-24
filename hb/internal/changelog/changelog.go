package changelog

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/klebersonromero/hb/internal/scanner"
)

// Consolidate varre todas as skills e gera o CHANGELOG-HUB.md
func Consolidate(root string) error {
	skills, err := scanner.FindSkills(root)
	if err != nil {
		return err
	}

	// Ordenar skills por nome
	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Name < skills[j].Name
	})

	var builder strings.Builder
	builder.WriteString("# Hub Global Changelog\n\n")
	builder.WriteString(fmt.Sprintf("*Last updated: %s*\n\n", time.Now().Format("2006-01-02 15:04:05")))
	builder.WriteString("This file consolidates the most recent updates from all skills in the Hub.\n\n")

	// Regex para encontrar seções que começam com ## [
	// Usamos lookahead (?=## \[) para quebrar o texto sem perder o delimitador
	re := regexp.MustCompile(`(?m)^## \[`)

	for _, s := range skills {
		changelogPath := filepath.Join(s.Path, "CHANGELOG.md")
		content, err := os.ReadFile(changelogPath)
		if err != nil {
			continue // Pular se não tiver changelog
		}

		// Quebrar em seções
		text := string(content)
		indices := re.FindAllStringIndex(text, -1)
		
		if len(indices) == 0 {
			continue
		}

		builder.WriteString(fmt.Sprintf("## 🧩 %s\n", s.Name))

		// Pegar até as 3 entradas mais recentes
		count := 0
		for i := 0; i < len(indices) && count < 3; i++ {
			start := indices[i][0]
			end := len(text)
			if i+1 < len(indices) {
				end = indices[i+1][0]
			}

			entry := strings.TrimSpace(text[start:end])
			builder.WriteString(entry)
			builder.WriteString("\n\n")
			count++
		}

		builder.WriteString("---\n\n")
	}

	outputPath := filepath.Join(root, "CHANGELOG-HUB.md")
	return os.WriteFile(outputPath, []byte(builder.String()), 0644)
}

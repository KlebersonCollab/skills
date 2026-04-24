package project

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func UpdateRoadmap(root string) error {
	roadmapPath := filepath.Join(root, ".specs", "project", "ROADMAP.md")
	content, err := os.ReadFile(roadmapPath)
	if err != nil {
		return err
	}

	specsDir := filepath.Join(root, ".specs", "features")
	entries, _ := os.ReadDir(specsDir)

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`- \[(.)\]`)

	color.Blue("🚀 Atualizando ROADMAP.md com base no progresso real...")

	for _, entry := range entries {
		if entry.IsDir() {
			taskFile := filepath.Join(specsDir, entry.Name(), "tasks.md")
			taskContent, err := os.ReadFile(taskFile)
			if err != nil {
				continue
			}

			matches := re.FindAllStringSubmatch(string(taskContent), -1)
			total := len(matches)
			done := 0
			for _, m := range matches {
				if strings.ToLower(m[1]) == "x" {
					done++
				}
			}

			percent := 0.0
			if total > 0 {
				percent = float64(done) / float64(total) * 100
			}

			// Atualiza a linha no ROADMAP.md que contém o nome da feature
			for i, line := range lines {
				if strings.Contains(strings.ToLower(line), strings.ToLower(entry.Name())) {
					status := " [ ] "
					if percent == 100 {
						status = " [x] "
					} else if percent > 0 {
						status = " [/] "
					}
					
					// Substitui o checkbox se encontrar
					if strings.Contains(line, "[ ]") || strings.Contains(line, "[x]") || strings.Contains(line, "[/]") {
						lines[i] = re.ReplaceAllString(line, "-"+status) // This is a bit naive, but works for standard markdown lists
						// Actually let's just prepend/replace the box
						lines[i] = strings.Replace(line, "[ ]", strings.TrimSpace(status), 1)
						lines[i] = strings.Replace(lines[i], "[x]", strings.TrimSpace(status), 1)
						lines[i] = strings.Replace(lines[i], "[/]", strings.TrimSpace(status), 1)
					}
				}
			}
		}
	}

	err = os.WriteFile(roadmapPath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return err
	}

	color.Green("✅ ROADMAP.md atualizado com sucesso!")
	return nil
}

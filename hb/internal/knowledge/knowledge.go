package knowledge

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

// LearningEntry representa uma entrada no LEARNINGS.md
type LearningEntry struct {
	Subject     string
	Description string
	Category    string // Pattern, Bug, Optimization
}

// AddLearning adiciona um novo aprendizado ao arquivo LEARNINGS.md
func AddLearning(root string, entry LearningEntry) error {
	learningsPath := filepath.Join(root, ".specs", "project", "LEARNINGS.md")
	os.MkdirAll(filepath.Dir(learningsPath), 0755)

	// Se o arquivo não existir, cria com header
	if _, err := os.Stat(learningsPath); os.IsNotExist(err) {
		header := "# Project Learnings\n\nThis file records patterns, bugs, and optimizations discovered during development.\n\n"
		os.WriteFile(learningsPath, []byte(header), 0644)
	}

	date := time.Now().Format("2006-01-02")
	content := fmt.Sprintf("## [%s] %s\n- **Category**: %s\n- **Description**: %s\n\n", 
		date, entry.Subject, entry.Category, entry.Description)

	f, err := os.OpenFile(learningsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return err
	}

	color.Green("✅ Aprendizado registrado em LEARNINGS.md")
	return nil
}

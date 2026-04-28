package graph

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateMermaidGraph creates a Mermaid diagram of the project features and their status
func GenerateMermaidGraph(root string) (string, error) {
	featuresDir := filepath.Join(root, ".specs", "features")
	entries, err := os.ReadDir(featuresDir)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.WriteString("graph TD\n")
	builder.WriteString("    Hub[Skills Hub] --> Features\n")
	builder.WriteString("    subgraph Features\n")

	for _, entry := range entries {
		if entry.IsDir() {
			slug := entry.Name()
			status := "InProgress"
			// Check if completed (simplification for YOLO)
			taskFile := filepath.Join(featuresDir, slug, "tasks.md")
			content, err := os.ReadFile(taskFile)
			if err == nil && !strings.Contains(string(content), "- [ ]") {
				status = "Completed"
			}

			color := "orange"
			if status == "Completed" {
				color = "green"
			}
			
			builder.WriteString(fmt.Sprintf("        %s[%s]\n", strings.ReplaceAll(slug, "-", "_"), slug))
			builder.WriteString(fmt.Sprintf("        style %s fill:%s\n", strings.ReplaceAll(slug, "-", "_"), color))
		}
	}

	builder.WriteString("    end\n")
	return builder.String(), nil
}

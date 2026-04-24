package sdd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func StartFeature(root, featureName string) error {
	slug := strings.ToLower(strings.ReplaceAll(featureName, " ", "-"))
	featureDir := filepath.Join(root, ".specs", "features", slug)

	if _, err := os.Stat(featureDir); err == nil {
		return fmt.Errorf("feature '%s' já existe em %s", slug, featureDir)
	}

	if err := os.MkdirAll(featureDir, 0755); err != nil {
		return err
	}

	files := map[string]string{
		"spec.md":     fmt.Sprintf("# Specification: %s\n\n## Goal\n\n## Functional Requirements\n\n## Acceptance Criteria (BDD)\n", featureName),
		"plan.md":     "# Technical Plan\n\n## Architecture\n\n## Mermaid Diagrams\n",
		"contract.md": "# Development Contract\n\n## Delivery Agreement\n\n## Sensors\n",
		"tasks.md":    "# Atomic Tasks\n\n- [ ] [TASK-1] Initialize feature artifacts\n",
	}

	for filename, content := range files {
		path := filepath.Join(featureDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

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
		"spec.md": fmt.Sprintf("# Specification: %s\n\n## Goal\n\n## Acceptance Criteria (BDD)\n- **Given** \n- **When** \n- **Then** \n", featureName),
		"plan.md": fmt.Sprintf("# Technical Plan: %s\n\n## Architecture\n\n## Mermaid Diagrams\n```mermaid\ngraph TD\n    A --> B\n```\n", featureName),
		"tasks.md": fmt.Sprintf("# Tasks: %s\n\n## Phase 1: Planning\n- [ ] T1: Define specifications. `id: TASK-001`\n- [ ] T2: Technical design. `id: TASK-002`\n\n## Phase 2: Implementation\n- [ ] T3: Core logic. `id: TASK-003`\n\n## Phase 3: Verification\n- [ ] T4: Quality audit. `id: TASK-004`\n", featureName),
		"contract.md": fmt.Sprintf("# Validation Contract: %s\n\n## Validation Sensors\n- [ ] Sensor: Test Coverage > 80%%\n- [ ] Sensor: Lint Passing\n- [ ] Sensor: BDD Scenarios Validated\n", featureName),
		"validation-report.md": fmt.Sprintf("# Validation Report: %s\n\n## Results\n- **Score**: 0/100\n- **Status**: Pending\n\n## Evidence\n- [ ] Evidence 1\n", featureName),
		"verification-report.md": fmt.Sprintf("# Quality Verification: %s\n\n## Checklist\n- [ ] Clean Code (SOLID, DRY)\n- [ ] Security Audit\n- [ ] Performance Baseline\n", featureName),
	}

	for filename, content := range files {
		path := filepath.Join(featureDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

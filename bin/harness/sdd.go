package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindWorkspaceRoot searches upwards from the current directory for the .specs directory
func FindWorkspaceRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		specsPath := filepath.Join(dir, ".specs")
		if info, err := os.Stat(specsPath); err == nil && info.IsDir() {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("workspace root (.specs directory) not found")
}

// GetSDDPhase reads .specs/project/STATE.md and returns the current phase
func GetSDDPhase() (string, error) {
	root, err := FindWorkspaceRoot()
	if err != nil {
		return "", err
	}

	statePath := filepath.Join(root, ".specs", "project", "STATE.md")
	file, err := os.Open(statePath)
	if err != nil {
		return "", fmt.Errorf("failed to open STATE.md: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inYaml := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "```yaml" {
			inYaml = true
			continue
		}
		if line == "```" && inYaml {
			inYaml = false
			continue
		}

		if inYaml {
			if strings.HasPrefix(line, "phase:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					phase := strings.Trim(parts[1], " \t\"'")
					return phase, nil
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("phase not found in .specs/project/STATE.md")
}

// VerifySDDGated checks if the current SDD phase allows file mutations
func VerifySDDGated() error {
	phase, err := GetSDDPhase()
	if err != nil {
		// If we can't find or read the state, default to strict: block the operation
		return fmt.Errorf("sdd check failed: %w", err)
	}

	phaseUpper := strings.ToUpper(phase)
	if phaseUpper == "IMPLEMENT" || phaseUpper == "VERIFY" {
		return nil
	}

	return fmt.Errorf("sdd workflow gate blocked: operations that modify the workspace are only allowed during IMPLEMENT or VERIFY phases (current phase: %s)", phase)
}

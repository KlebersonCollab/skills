package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetSDDPhase reads .specs/project/STATE.md and returns the current phase.
func GetSDDPhase(root string) (string, error) {
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
					return strings.Trim(parts[1], " \t\"'"), nil
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("phase not found in .specs/project/STATE.md")
}

// VerifySDDGated checks if the current SDD phase allows file mutations.
// Returns nil if mutations are allowed (IMPLEMENT or VERIFY), error otherwise.
func VerifySDDGated(root string) error {
	phase, err := GetSDDPhase(root)
	if err != nil {
		return fmt.Errorf("sdd check failed: %w", err)
	}

	phaseUpper := strings.ToUpper(phase)
	if phaseUpper == "IMPLEMENT" || phaseUpper == "VERIFY" {
		return nil
	}

	return fmt.Errorf("sdd workflow gate blocked: operations that modify the workspace are only allowed during IMPLEMENT or VERIFY phases (current phase: %s)", phase)
}

package git

import (
	"bytes"
	"os/exec"
	"strings"
)

// StagedDiff returns the unified diff of currently staged changes.
func StagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--unified=3")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return stdout.String(), nil
}

// DiffHunk represents a single hunk in a unified diff.
type DiffHunk struct {
	FilePath string
	OldStart int
	OldLines int
	NewStart int
	NewLines int
	Lines    []string
}

// GetDiffStats returns the total additions and removals in the staged changes.
func GetDiffStats() (int, int, error) {
	diff, err := StagedDiff()
	if err != nil {
		return 0, 0, err
	}

	hunks, err := ParseDiff(diff)
	if err != nil {
		return 0, 0, err
	}

	added := 0
	removed := 0

	for _, hunk := range hunks {
		for _, line := range hunk.Lines {
			if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
				added++
			} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
				removed++
			}
		}
	}

	return added, removed, nil
}

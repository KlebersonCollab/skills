package git

import (
	"os/exec"
	"strconv"
	"strings"
)

// GetDiffStats retorna o número de linhas adicionadas e removidas
func GetDiffStats() (added int, removed int, err error) {
	// git diff --numstat
	cmd := exec.Command("git", "diff", "--numstat")
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			a, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			added += a
			removed += r
		}
	}
	return added, removed, nil
}

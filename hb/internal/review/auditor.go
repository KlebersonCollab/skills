package review

import (
	"fmt"
	"os"
	"strings"

	"github.com/klebersonromero/hb/internal/git"
	"github.com/klebersonromero/hb/internal/sdd"
)

// AuditResult represents the outcome of a semantic review.
type AuditResult struct {
	Pass     bool
	Score    int
	Warnings []string
}

// AuditStagedChanges performs a semantic review of staged changes against SDD context.
func AuditStagedChanges(ctx *sdd.FeatureContext, hunks []git.DiffHunk) *AuditResult {
	result := &AuditResult{
		Pass:  true,
		Score: 100,
	}

	for _, hunk := range hunks {
		// 1. Check for Junk (Debug prints)
		checkJunk(hunk, result)

		// 2. Check Task Alignment (Simple keyword matching for now)
		checkTaskAlignment(ctx, hunk, result)
	}

	if result.Score < 70 {
		result.Pass = false
	}

	return result
}

func checkJunk(hunk git.DiffHunk, result *AuditResult) {
	junkPatterns := []string{
		"fmt.Printf", "fmt.Println", "console.log", "DEBUG:", "FIXME:",
	}

	for _, line := range hunk.Lines {
		if !strings.HasPrefix(line, "+") {
			continue
		}

		for _, pattern := range junkPatterns {
			if strings.Contains(line, pattern) {
				result.Warnings = append(result.Warnings, fmt.Sprintf("Possível código de debug encontrado em %s: %s", hunk.FilePath, strings.TrimSpace(line[1:])))
				result.Score -= 10
			}
		}
	}
}

func checkTaskAlignment(ctx *sdd.FeatureContext, hunk git.DiffHunk, result *AuditResult) {
	taskDesc := strings.ToLower(ctx.ActiveTask.Description)
	fileName := strings.ToLower(filepathBase(hunk.FilePath))

	// 1. Plan Alignment: Check if the file is mentioned in plan.md
	planPath := filepathJoin(ctx.Path, "plan.md")
	planContent, err := osReadFile(planPath)
	if err == nil {
		if !strings.Contains(strings.ToLower(string(planContent)), fileName) {
			result.Warnings = append(result.Warnings, fmt.Sprintf("Arquivo %s não está mapeado no plan.md desta feature.", hunk.FilePath))
			result.Score -= 5
		}
	}

	// 2. Trash Detection: Temporary or log files
	if strings.HasSuffix(fileName, ".log") || strings.HasSuffix(fileName, ".tmp") {
		result.Warnings = append(result.Warnings, fmt.Sprintf("Arquivo temporário ou log detectado no stage: %s", hunk.FilePath))
		result.Score -= 20
	}

	_ = taskDesc
}

func filepathJoin(parts ...string) string {
	return strings.Join(parts, "/")
}

func osReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func filepathBase(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

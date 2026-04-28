package security

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/klebersonromero/hb/internal/git"
)

// SecurityResult holds the findings of a scan
type SecurityResult struct {
	Violations []string
}

// Issue represents a security problem found in a file (compatible with harness)
type Issue struct {
	Type    string
	Message string
	File    string
	Line    int
}

// SecurityRules defines patterns for secrets and PII
var SecurityRules = map[string]*regexp.Regexp{
	"API Key":        regexp.MustCompile(`(?i)(api_key|apikey|secret|token|password)[\s:=]+['"][a-zA-Z0-9\-_]{20,}['"]`),
	"Brazilian CPF":  regexp.MustCompile(`\d{3}\.\d{3}\.\d{3}-\d{2}`),
	"Brazilian RG":   regexp.MustCompile(`\d{2}\.\d{3}\.\d{3}-[\dX]`),
	"Email":          regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`),
}

// ScanHunks checks git hunks for security violations
func ScanHunks(hunks []git.DiffHunk) *SecurityResult {
	result := &SecurityResult{}

	for _, hunk := range hunks {
		for _, line := range hunk.Lines {
			if !strings.HasPrefix(line, "+") {
				continue
			}
			
			cleanLine := strings.TrimPrefix(line, "+")
			for name, re := range SecurityRules {
				if re.MatchString(cleanLine) {
					result.Violations = append(result.Violations, fmt.Sprintf("⚠️ [%s] Detectado em %s: %s", name, hunk.FilePath, strings.TrimSpace(cleanLine)))
				}
			}
		}
	}

	return result
}

// ScanPath scans all files in a directory for security violations
func ScanPath(root string) ([]Issue, error) {
	var issues []Issue
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		// Skip .git and binary files
		if strings.Contains(path, ".git") || info.Size() > 1000000 {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			for name, re := range SecurityRules {
				if re.MatchString(line) {
					issues = append(issues, Issue{
						Type:    name,
						Message: "Detectado padrão sensível",
						File:    path,
						Line:    i + 1,
					})
				}
			}
		}
		return nil
	})
	return issues, err
}

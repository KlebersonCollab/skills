package security

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// Issue representa um problema de segurança detectado
type Issue struct {
	Type     string
	Severity string // "Critical", "Warning", "Info"
	Message  string
	File     string
	Line     int
}

// Patterns comuns de segredos
var secretPatterns = map[string]*regexp.Regexp{
	"Generic API Key": regexp.MustCompile(`(?i)(key|api|token|secret|password|passwd)[-|_]?(key|token|secret|password|passwd)?\s*[:|=]\s*["']?[0-9a-zA-Z]{16,}["']?`),
	"AWS Access Key":  regexp.MustCompile(`(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}`),
	"Private Key":     regexp.MustCompile(`(?s)-----BEGIN (RSA|EC|DSA|GPG|OPENSSH) PRIVATE KEY-----`),
	"GitHub Token":    regexp.MustCompile(`(ghp|gho|ghu|ghs|ghr)_[a-zA-Z0-9]{36}`),
}

// ScanPath executa uma auditoria de segurança em um arquivo ou diretório
func ScanPath(root string) ([]Issue, error) {
	var issues []Issue

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Ignorar diretórios técnicos
		if d.IsDir() {
			if d.Name() == ".git" || d.Name() == "node_modules" || d.Name() == ".venv" {
				return filepath.SkipDir
			}
			return nil
		}

		// Analisar apenas arquivos de texto/código
		if isCodeFile(path) {
			fileIssues, err := scanFile(path)
			if err != nil {
				return err
			}
			issues = append(issues, fileIssues...)
		}

		return nil
	})

	return issues, err
}

func isCodeFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	codeExts := map[string]bool{
		".go":   true, ".py":   true, ".js":   true, ".ts":   true,
		".html": true, ".css":  true, ".yml":  true, ".yaml": true,
		".json": true, ".md":   true, ".sh":   true, ".env":  true,
	}
	return codeExts[ext]
}

func scanFile(path string) ([]Issue, error) {
	var issues []Issue
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		for name, pattern := range secretPatterns {
			if pattern.MatchString(line) {
				issues = append(issues, Issue{
					Type:     "Secret Leak",
					Severity: "Critical",
					Message:  fmt.Sprintf("Detectado possível segredo: %s", name),
					File:     path,
					Line:     lineNum,
				})
				color.Yellow("⚠️  Cuidado: Segredo detectado em %s:%d (%s)", path, lineNum, name)
			}
		}
	}

	return issues, scanner.Err()
}

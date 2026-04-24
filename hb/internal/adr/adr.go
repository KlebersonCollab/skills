package adr

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

const adrTemplate = `# ADR-%04d: %s

*   **Status**: %s
*   **Date**: %s

## Context
Describe the situation and the problem being addressed.

## Decision
Explain the proposed solution and the rationale behind it.

## Consequences
What are the expected outcomes, trade-offs, and impacts?
`

// CreateNew gera um novo arquivo ADR
func CreateNew(root, title string) (string, error) {
	adrDir := filepath.Join(root, ".specs", "architecture")
	os.MkdirAll(adrDir, 0755)

	// Encontrar o próximo ID
	nextID := 1
	files, _ := os.ReadDir(adrDir)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "ADR-") {
			var id int
			fmt.Sscanf(f.Name(), "ADR-%d", &id)
			if id >= nextID {
				nextID = id + 1
			}
		}
	}

	// Sanitizar título para o nome do arquivo
	safeTitle := strings.ToLower(title)
	reg := regexp.MustCompile("[^a-z0-9]+")
	safeTitle = reg.ReplaceAllString(safeTitle, "-")
	safeTitle = strings.Trim(safeTitle, "-")

	fileName := fmt.Sprintf("ADR-%04d-%s.md", nextID, safeTitle)
	filePath := filepath.Join(adrDir, fileName)

	content := fmt.Sprintf(adrTemplate, nextID, title, "Proposed", time.Now().Format("2006-01-02"))
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

// List exibe a lista de ADRs
func List(root string) error {
	adrDir := filepath.Join(root, ".specs", "architecture")
	if _, err := os.Stat(adrDir); os.IsNotExist(err) {
		return fmt.Errorf("diretório de ADRs não existe em %s", adrDir)
	}

	files, err := os.ReadDir(adrDir)
	if err != nil {
		return err
	}

	color.Cyan("%-10s | %-40s | %-10s", "ID", "Title", "Status")
	color.Cyan(strings.Repeat("-", 65))

	var adrFiles []string
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "ADR-") && strings.HasSuffix(f.Name(), ".md") {
			adrFiles = append(adrFiles, f.Name())
		}
	}
	sort.Strings(adrFiles)

	for _, fileName := range adrFiles {
		fullPath := filepath.Join(adrDir, fileName)
		content, _ := os.ReadFile(fullPath)
		
		status := "Unknown"
		title := strings.TrimSuffix(strings.TrimPrefix(fileName, "ADR-XXXX-"), ".md")
		
		// Parse status and title from content
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "# ADR-") {
				parts := strings.SplitN(line, ": ", 2)
				if len(parts) == 2 {
					title = parts[1]
				}
			}
			if strings.HasPrefix(line, "*   **Status**:") {
				status = strings.TrimSpace(strings.TrimPrefix(line, "*   **Status**: "))
			}
		}

		id := fileName[0:8]
		color.White("%-10s | %-40s | %-10s", id, truncate(title, 40), status)
	}

	return nil
}

func truncate(s string, l int) string {
	if len(s) > l {
		return s[:l-3] + "..."
	}
	return s
}

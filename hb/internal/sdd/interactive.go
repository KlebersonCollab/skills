package sdd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// SelectFeatureInteractively lists features and lets the user pick one
func SelectFeatureInteractively(root string) (string, error) {
	specsDir := filepath.Join(root, ".specs", "features")
	entries, err := os.ReadDir(specsDir)
	if err != nil {
		return "", err
	}

	var features []string
	for _, entry := range entries {
		if entry.IsDir() {
			features = append(features, entry.Name())
		}
	}

	if len(features) == 0 {
		return "", fmt.Errorf("nenhuma feature encontrada")
	}

	color.Cyan("\n📁 Selecione uma Feature:")
	for i, f := range features {
		fmt.Printf("[%d] %s\n", i+1, f)
	}

	fmt.Print("\nOpção: ")
	var input string
	fmt.Scanln(&input)
	idx, err := strconv.Atoi(input)
	if err != nil || idx < 1 || idx > len(features) {
		return "", fmt.Errorf("opção inválida")
	}

	return features[idx-1], nil
}

// SelectTaskInteractively lists tasks for a feature and lets the user pick one to complete
func SelectTaskInteractively(root, feature string) (string, error) {
	taskFile := filepath.Join(root, ".specs", "features", feature, "tasks.md")
	content, err := os.ReadFile(taskFile)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	var tasks []string
	var taskIDs []string

	color.Cyan("\n✅ Selecione a Tarefa para concluir (%s):", feature)
	count := 1
	for _, line := range lines {
		if strings.Contains(line, "- [ ]") {
			tasks = append(tasks, line)
			// Extract ID or description
			clean := strings.TrimSpace(strings.Replace(line, "- [ ]", "", 1))
			taskIDs = append(taskIDs, clean)
			fmt.Printf("[%d] %s\n", count, clean)
			count++
		}
	}

	if len(tasks) == 0 {
		return "", fmt.Errorf("nenhuma tarefa pendente encontrada")
	}

	fmt.Print("\nOpção: ")
	var input string
	fmt.Scanln(&input)
	idx, err := strconv.Atoi(input)
	if err != nil || idx < 1 || idx > len(tasks) {
		return "", fmt.Errorf("opção inválida")
	}

	// Returns the original description to match in sdd task
	return taskIDs[idx-1], nil
}

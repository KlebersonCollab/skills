package sdd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Task represents an atomic SDD task.
type Task struct {
	ID          string
	Description string
	Completed   bool
}

// FeatureContext holds the current state of a feature.
type FeatureContext struct {
	Name       string
	Path       string
	ActiveTask *Task
}

// LoadActiveContext scans for a feature directory and finds the active task.
func LoadActiveContext(root string) (*FeatureContext, error) {
	featuresDir := filepath.Join(root, ".specs", "features")
	entries, err := os.ReadDir(featuresDir)
	if err != nil {
		return nil, fmt.Errorf("could not read features dir: %v", err)
	}

	// For simplicity, we find the first directory that has a tasks.md with pending tasks
	// In a real scenario, we might use a .feature-active file
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		tasksPath := filepath.Join(featuresDir, entry.Name(), "tasks.md")
		tasks, err := ParseTasks(tasksPath)
		if err != nil {
			continue
		}

		for _, t := range tasks {
			if !t.Completed {
				return &FeatureContext{
					Name:       entry.Name(),
					Path:       filepath.Join(featuresDir, entry.Name()),
					ActiveTask: &t,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("no active task found in any feature")
}

// ParseTasks reads a tasks.md file and extracts tasks.
func ParseTasks(path string) ([]Task, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- [ ]") || strings.HasPrefix(line, "- [x]") {
			completed := strings.HasPrefix(line, "- [x]")
			desc := strings.TrimSpace(line[5:])
			
			// Extract ID if present (e.g., `id: REV-001`)
			id := ""
			if idx := strings.Index(desc, "`id:"); idx != -1 {
				id = strings.Trim(desc[idx+4:], "` ")
				desc = strings.TrimSpace(desc[:idx])
			}

			tasks = append(tasks, Task{
				ID:          id,
				Description: desc,
				Completed:   completed,
			})
		}
	}

	return tasks, nil
}

package harness

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

)

// Rehydrate Snapshot structure
type Snapshot struct {
	State     string
	Memory    string
	Learnings string
	Feature   string
	Spec      string
	Tasks     string
}

func Rehydrate(root, featureName string) (string, error) {
	snapshot := &Snapshot{}

	// 1. Core Memory (Tríade)
	statePath := filepath.Join(root, ".specs", "project", "STATE.md")
	memoryPath := filepath.Join(root, ".specs", "project", "MEMORY.md")
	learningsPath := filepath.Join(root, ".specs", "project", "LEARNINGS.md")

	snapshot.State = readFileIfExists(statePath)
	snapshot.Memory = readFileIfExists(memoryPath)
	snapshot.Learnings = readFileIfExists(learningsPath)

	// 2. Feature Specific Context
	if featureName != "" {
		featureDir := filepath.Join(root, ".specs", "features", featureName)
		if _, err := os.Stat(featureDir); os.IsNotExist(err) {
			// Try slug
			featureName = strings.ToLower(strings.ReplaceAll(featureName, " ", "-"))
			featureDir = filepath.Join(root, ".specs", "features", featureName)
		}

		snapshot.Feature = featureName
		snapshot.Spec = readFileIfExists(filepath.Join(featureDir, "spec.md"))
		snapshot.Tasks = readFileIfExists(filepath.Join(featureDir, "tasks.md"))
	}

	// 3. Construct Output Block
	var sb strings.Builder
	sb.WriteString("<!-- HARNESS REHYDRATE SNAPSHOT START -->\n")
	sb.WriteString("# 🧠 OPERATIONAL CONTEXT REHYDRATION\n\n")
	
	sb.WriteString("## 🚀 Current State\n")
	sb.WriteString(snapshot.State + "\n\n")

	sb.WriteString("## 📜 Project Memory\n")
	sb.WriteString(snapshot.Memory + "\n\n")

	sb.WriteString("## 💡 Recent Learnings\n")
	sb.WriteString(snapshot.Learnings + "\n\n")

	if snapshot.Feature != "" {
		sb.WriteString(fmt.Sprintf("## 🛠️ Active Feature: %s\n", snapshot.Feature))
		sb.WriteString("### Specifications\n")
		sb.WriteString(snapshot.Spec + "\n\n")
		sb.WriteString("### Task Progress\n")
		sb.WriteString(snapshot.Tasks + "\n\n")
	}

	sb.WriteString("<!-- HARNESS REHYDRATE SNAPSHOT END -->")

	return sb.String(), nil
}

func readFileIfExists(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return "_File not found or empty._"
	}
	return string(content)
}

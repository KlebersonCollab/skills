package knowledge

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// KnowledgeBundle represents a serialized set of learnings.
type KnowledgeBundle struct {
	ID        string    `json:"id"`
	Feature   string    `json:"feature"`
	Timestamp time.Time `json:"timestamp"`
	Learnings string    `json:"learnings"`
	Memory    string    `json:"memory"`
}

// ExportFeature packs learnings from a feature into a bundle.
func ExportFeature(root, featureName string) (*KnowledgeBundle, error) {
	featurePath := filepath.Join(root, ".specs", "features", featureName)
	
	learnings, _ := os.ReadFile(filepath.Join(featurePath, "LEARNINGS.md"))
	memory, _ := os.ReadFile(filepath.Join(featurePath, "MEMORY.md"))

	bundle := &KnowledgeBundle{
		ID:        fmt.Sprintf("KB-%d", time.Now().Unix()),
		Feature:   featureName,
		Timestamp: time.Now(),
		Learnings: string(learnings),
		Memory:    string(memory),
	}

	return bundle, nil
}

// SaveBundle saves the bundle to the global knowledge hub.
func SaveBundle(root string, bundle *KnowledgeBundle) error {
	hubDir := filepath.Join(root, ".specs", "knowledge", "global")
	if err := os.MkdirAll(hubDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s_%s.json", bundle.Feature, bundle.ID)
	return os.WriteFile(filepath.Join(hubDir, filename), data, 0644)
}

// ImportBundle reads a bundle from the global hub.
func ImportBundle(root, bundleName string) (*KnowledgeBundle, error) {
	path := filepath.Join(root, ".specs", "knowledge", "global", bundleName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var bundle KnowledgeBundle
	if err := json.Unmarshal(data, &bundle); err != nil {
		return nil, err
	}

	return &bundle, nil
}

// MergeKnowledge injects bundle learnings into the target feature.
func MergeKnowledge(root, featureName string, bundle *KnowledgeBundle) error {
	featurePath := filepath.Join(root, ".specs", "features", featureName)
	
	// Append Learnings
	learningsPath := filepath.Join(featurePath, "LEARNINGS.md")
	f, err := os.OpenFile(learningsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	header := fmt.Sprintf("\n\n--- Imported from %s (%s) ---\n", bundle.Feature, bundle.ID)
	if _, err := f.WriteString(header + bundle.Learnings); err != nil {
		return err
	}

	return nil
}

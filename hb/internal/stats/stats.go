package stats

// Dummy comment for git stats testing

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/klebersonromero/hb/internal/domain"
)

const StatsFileName = "STATS.tmp"

// LoadStats carrega as estatísticas da sessão atual
func LoadStats(root string) (*domain.SessionStats, error) {
	statsFile := filepath.Join(root, ".specs", "project", StatsFileName)
	data, err := os.ReadFile(statsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Retorna nil se o arquivo não existe (sessão não iniciada)
		}
		return nil, err
	}

	var stats domain.SessionStats
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

// SaveStats persiste as estatísticas no arquivo temporário
func SaveStats(root string, stats *domain.SessionStats) error {
	statsDir := filepath.Join(root, ".specs", "project")
	if _, err := os.Stat(statsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(statsDir, 0755); err != nil {
			return err
		}
	}

	statsFile := filepath.Join(statsDir, StatsFileName)
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(statsFile, data, 0644)
}

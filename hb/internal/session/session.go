package session

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/domain"
	"github.com/klebersonromero/hb/internal/stats"
)

func Start(root string) error {
	sessionFile := filepath.Join(root, ".specs", "project", "SESSION.tmp")
	now := time.Now()
	startTime := now.Format(time.RFC3339)
	
	err := os.WriteFile(sessionFile, []byte(startTime), 0644)
	if err != nil {
		return err
	}

	// Inicializa estatísticas
	initialStats := &domain.SessionStats{
		SessionID: fmt.Sprintf("session-%d", now.Unix()),
		StartTime: now,
		Models:    make(map[string]*domain.ModelUsage),
	}
	stats.SaveStats(root, initialStats)

	color.Cyan("🎬 Sessão iniciada em %s", startTime)
	color.Blue("O estado operacional está sendo monitorado.")
	return nil
}

func Stop(root string) error {
	sessionFile := filepath.Join(root, ".specs", "project", "SESSION.tmp")
	startTimeStr, err := os.ReadFile(sessionFile)
	if err != nil {
		return fmt.Errorf("nenhuma sessão ativa encontrada. Use 'hb session start' primeiro")
	}

	startTime, _ := time.Parse(time.RFC3339, string(startTimeStr))
	duration := time.Since(startTime)

	color.Cyan("🛑 Sessão encerrada.")
	fmt.Printf("Duração: %s\n", duration.Round(time.Second))

	// Exibe sumário de estatísticas se existir
	s, _ := stats.LoadStats(root)
	if s != nil {
		fmt.Printf("Custo Total: $%.4f\n", s.TotalCostUSD)
		if s.Git.LinesAdded > 0 || s.Git.LinesRemoved > 0 {
			fmt.Printf("Alterações: +%d -%d linhas\n", s.Git.LinesAdded, s.Git.LinesRemoved)
		}
	}

	// Remove os arquivos temporários
	os.Remove(sessionFile)
	os.Remove(filepath.Join(root, ".specs", "project", stats.StatsFileName))

	color.Yellow("\nSugestão: Execute 'hb learn' para capturar os aprendizados desta sessão.")
	color.Blue("Não esqueça de rodar 'hb sync' para persistir o estado global.")

	return nil
}

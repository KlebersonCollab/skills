package session

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

func Start(root string) error {
	sessionFile := filepath.Join(root, ".specs", "project", "SESSION.tmp")
	startTime := time.Now().Format(time.RFC3339)
	
	err := os.WriteFile(sessionFile, []byte(startTime), 0644)
	if err != nil {
		return err
	}

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

	// Remove o arquivo de sessão
	os.Remove(sessionFile)

	color.Yellow("\nSugestão: Execute 'hb learn' para capturar os aprendizados desta sessão.")
	color.Blue("Não esqueça de rodar 'hb sync' para persistir o estado global.")

	return nil
}

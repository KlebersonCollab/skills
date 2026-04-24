package mentor

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func AuditFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	
	color.Blue("👨‍🏫 Mentor Audit: %s", path)
	violations := 0

	// Princípio: KISS (Tamanho do arquivo)
	if len(lines) > 500 {
		color.Yellow("⚠️  Princípio KISS: Arquivo muito longo (%d linhas). Considere quebrar em módulos menores.", len(lines))
		violations++
	}

	// Princípio: DRY (Repetições básicas)
	// (Simulação de detecção de duplicidade ou falta de abstração)
	if strings.Count(string(content), "func ") > 20 {
		color.Yellow("⚠️  Princípio SRP: Muitas funções em um único arquivo. Verifique se as responsabilidades estão misturadas.")
		violations++
	}

	if violations == 0 {
		color.Green("✅ Nenhuma violação óbvia de princípios encontrada. Bom trabalho!")
	} else {
		fmt.Printf("\nTotal de avisos: %d\n", violations)
	}

	return nil
}

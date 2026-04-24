package report

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/domain"
)

// PrintSummary exibe uma tabela elegante com os resultados do audit usando tabwriter nativo
func PrintSummary(results []domain.AuditResult) {
	// minwidth, tabwidth, padding, padchar, flags
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	
	// Header
	fmt.Fprintln(w, "SKILL\tVERSION\tCATEGORY\tSTATUS\tISSUES")
	
	passed := 0
	failed := 0

	for _, res := range results {
		status := "PASS"
		if !res.Success() {
			status = "FAIL"
			failed++
		} else {
			passed++
		}

		issuesCount := strconv.Itoa(len(res.Issues))
		
		// Colorir para o terminal (tabwriter preserva as sequências de escape se configurado)
		colorStatus := color.GreenString(status)
		if status == "FAIL" {
			colorStatus = color.RedString(status)
			issuesCount = color.RedString(issuesCount)
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", 
			res.Skill.Name, 
			res.Skill.Version, 
			res.Skill.Category, 
			colorStatus, 
			issuesCount,
		)
	}

	w.Flush()

	// Sumário final
	fmt.Println()
	color.Cyan("--- Final Summary ---")
	fmt.Printf("Total Skills: %d\n", len(results))
	color.Green("Passed:       %d", passed)
	if failed > 0 {
		color.Red("Failed:       %d", failed)
	}
	fmt.Println()
}

// PrintSkillList exibe a tabela de todas as skills encontradas
func PrintSkillList(skills []domain.Skill) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "SKILL\tVERSION\tCATEGORY\tPATH")

	for _, s := range skills {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			s.Name,
			s.Version,
			s.Category,
			s.Path,
		)
	}
	w.Flush()
	fmt.Printf("\nTotal: %d skills.\n", len(skills))
}

// PrintIssues detalha os erros encontrados
func PrintIssues(results []domain.AuditResult) {
	for _, res := range results {
		if !res.Success() {
			color.Red("\n[FAIL] %s:", res.Skill.Name)
			for _, issue := range res.Issues {
				fmt.Printf("  - %s: %s\n", color.YellowString(issue.File), issue.Message)
			}
		}
	}
}

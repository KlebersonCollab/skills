package harness

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

)

type Evaluation struct {
	Design        float64
	Originality   float64
	Craft         float64
	Functionality float64
	TotalScore    float64
}

func (e *Evaluation) Calculate() {
	e.TotalScore = (e.Design * 0.3) + (e.Originality * 0.2) + (e.Craft * 0.3) + (e.Functionality * 0.2)
}

func RunEvaluation(root, feature string, design, orig, craft, funcScore float64) (*Evaluation, error) {
	eval := &Evaluation{
		Design:        design,
		Originality:   orig,
		Craft:         craft,
		Functionality: funcScore,
	}
	eval.Calculate()

	featureDir := filepath.Join(root, ".specs", "features", feature)
	if _, err := os.Stat(featureDir); os.IsNotExist(err) {
		// Tenta buscar por slug
		feature = strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
		featureDir = filepath.Join(root, ".specs", "features", feature)
	}

	reportPath := filepath.Join(featureDir, "validation-report.md")
	
	content := fmt.Sprintf(`# Validation Report: %s

## 📊 GAN-Style Score: %.1f/10.0

| Criterion | Weight | Score | Weighted |
|-----------|--------|-------|----------|
| **Design Quality** | 0.3 | %.1f | %.2f |
| **Originality** | 0.2 | %.1f | %.2f |
| **Craft** | 0.3 | %.1f | %.2f |
| **Functionality** | 0.2 | %.1f | %.2f |
| **TOTAL** | **1.0** | | **%.1f** |

## 📝 Evidence & Feedback
- [Recorded via HB-CLI Harness]

## 🛡️ Status
%s
`, feature, eval.TotalScore, 
	eval.Design, eval.Design*0.3,
	eval.Originality, eval.Originality*0.2,
	eval.Craft, eval.Craft*0.3,
	eval.Functionality, eval.Functionality*0.2,
	eval.TotalScore,
	getStatusMessage(eval.TotalScore))

	err := os.WriteFile(reportPath, []byte(content), 0644)
	if err != nil {
		return nil, err
	}

	return eval, nil
}

func getStatusMessage(score float64) string {
	if score >= 7.0 {
		return "✅ **APPROVED**: Ready for production sync."
	}
	return "❌ **REJECTED**: Score below 7.0. Improvements required."
}

// GetLatestScore procura o score mais recente nos relatórios de validação
func GetLatestScore(root string) (float64, string, error) {
	specsDir := filepath.Join(root, ".specs", "features")
	entries, err := os.ReadDir(specsDir)
	if err != nil {
		return 0, "", nil
	}

	var lowestScore float64 = 10.0
	var foundFeature string
	found := false

	// Regex para encontrar o score no markdown
	re := regexp.MustCompile(`GAN-Style Score: ([\d.]+)/10.0`)

	for _, entry := range entries {
		if entry.IsDir() {
			reportPath := filepath.Join(specsDir, entry.Name(), "validation-report.md")
			content, err := os.ReadFile(reportPath)
			if err != nil {
				continue
			}

			matches := re.FindStringSubmatch(string(content))
			if len(matches) > 1 {
				score, _ := strconv.ParseFloat(matches[1], 64)
				if score < lowestScore {
					lowestScore = score
					foundFeature = entry.Name()
				}
				found = true
			}
		}
	}

	if !found {
		return 10.0, "", nil // Se não houver relatórios, assume que está tudo bem (ou muda para rigoroso se preferir)
	}

	return lowestScore, foundFeature, nil
}

package checker

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/scanner"
)

// CheckVersionParity valida se as versões no README.md batem com os SKILL.md
func CheckVersionParity(root string) error {
	readmePath := filepath.Join(root, "README.md")
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("README.md not found: %w", err)
	}

	// Regex mais flexível: | # | **[Name](path/)** | Desc | `version` |
	re := regexp.MustCompile(`(?m)\|\s*\d+\s*\|\s*\*\*\[.*?\]\((.*?)/?\)\s*\*\*\s*\|.*?\|\s*` + "`" + `(.*?)` + "`" + `\s*\|`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	skills, err := scanner.FindSkills(root)
	if err != nil {
		return err
	}

	skillMap := make(map[string]string)
	for _, s := range skills {
		relPath, _ := filepath.Rel(root, s.Path)
		skillMap[strings.TrimSuffix(relPath, "/")] = s.Version
	}

	errorsFound := false
	for _, match := range matches {
		skillRelPath := strings.Trim(match[1], "./ ")
		skillRelPath = strings.TrimSuffix(skillRelPath, "/")
		readmeVersion := strings.TrimSpace(match[2])

		actualVersion, exists := skillMap[skillRelPath]
		if !exists {
			// Tentar sem o prefixo ./
			actualVersion, exists = skillMap[strings.TrimPrefix(skillRelPath, "./")]
		}

		if exists {
			if actualVersion != readmeVersion {
				color.Red("   ❌ %s: README (%s) != SKILL.md (%s)", skillRelPath, readmeVersion, actualVersion)
				errorsFound = true
			} else {
				color.Cyan("   ✅ %s: %s (Consistent)", skillRelPath, actualVersion)
			}
		}
	}

	if errorsFound {
		return fmt.Errorf("version inconsistency detected between README.md and skills")
	}
	return nil
}

// CheckContextDrift verifica se a memória (STATE/LEARNINGS) está desatualizada em relação ao código
func CheckContextDrift(root string) error {
	specsDir := filepath.Join(root, ".specs", "project")
	memoryFiles := []string{
		filepath.Join(specsDir, "STATE.md"),
		filepath.Join(specsDir, "MEMORY.md"),
		filepath.Join(specsDir, "LEARNINGS.md"),
	}

	var lastMemoryUpdate time.Time
	for _, f := range memoryFiles {
		info, err := os.Stat(f)
		if err == nil {
			if info.ModTime().After(lastMemoryUpdate) {
				lastMemoryUpdate = info.ModTime()
			}
		}
	}

	if lastMemoryUpdate.IsZero() {
		return fmt.Errorf("memory files not found in .specs/project")
	}

	color.Blue("   -> Last Memory Update: %s", lastMemoryUpdate.Format("2006-01-02 15:04:05"))

	var unsynced []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		// Ignorar pastas de infra e specs
		rel, _ := filepath.Rel(root, path)
		if strings.HasPrefix(rel, ".specs") || strings.HasPrefix(rel, "hb") || strings.HasPrefix(rel, "bin") || strings.HasPrefix(rel, "vendor") || strings.Contains(rel, "__pycache__") {
			return nil
		}

		// Monitorar apenas extensões relevantes
		ext := filepath.Ext(path)
		relevant := false
		for _, e := range []string{".go", ".py", ".md", ".js", ".ts", ".dart"} {
			if ext == e {
				relevant = true
				break
			}
		}

		if relevant && info.ModTime().After(lastMemoryUpdate) {
			unsynced = append(unsynced, rel)
		}
		return nil
	})

	if err != nil {
		return err
	}

	if len(unsynced) > 0 {
		color.Yellow("\n   ⚠️  CONTEXT DRIFT DETECTED!")
		color.Yellow("   The following files were modified AFTER the last memory update:")
		for i, f := range unsynced {
			if i >= 5 {
				color.Yellow("   ... and %d more files.", len(unsynced)-5)
				break
			}
			color.Yellow("     - %s", f)
		}
		color.Yellow("\n   👉 Action: Update .specs/project/STATE.md or LEARNINGS.md to sync context.")
	} else {
		color.Green("   ✅ Memory is in sync with codebase changes.")
	}

	return nil
}

// CheckLinks verifica se links locais (arquivos/pastas) existem
func CheckLinks(root string) error {
	color.Blue("\n🔍 Validando links internos do ecossistema...")

	if _, err := os.Stat(root); os.IsNotExist(err) {
		return fmt.Errorf("root directory does not exist: %s", root)
	}

	errorsFound := false
	// Regex simples: [label](target)
	re := regexp.MustCompile(`\[.*?\]\((.*?)\)`)

	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if info == nil || info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		// Ignorar pastas técnicas e pastas de agentes
		rel, _ := filepath.Rel(root, path)
		if strings.HasPrefix(rel, ".specs") || strings.HasPrefix(rel, "hb") || strings.HasPrefix(rel, "bin") || 
		   strings.HasPrefix(rel, "vendor") || strings.Contains(rel, "__pycache__") ||
		   strings.HasPrefix(rel, ".gemini") || strings.HasPrefix(rel, ".claude") || strings.HasPrefix(rel, ".agent") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		matches := re.FindAllStringSubmatch(string(content), -1)
		for _, match := range matches {
			if len(match) < 2 {
				continue
			}
			linkTarget := match[1]

			// Filtrar links externos, âncoras ou templates manualmente
			if strings.HasPrefix(linkTarget, "http") || strings.HasPrefix(linkTarget, "#") || 
			   strings.HasPrefix(linkTarget, "{{") || strings.HasPrefix(linkTarget, "file://") {
				continue
			}

			// Heurística básica para evitar falsos positivos de código:
			// Um link local deve ter uma barra '/' ou uma extensão conhecida
			isPotentialPath := strings.Contains(linkTarget, "/") || 
							   strings.HasSuffix(linkTarget, ".md") || 
							   strings.HasSuffix(linkTarget, ".py") || 
							   strings.HasSuffix(linkTarget, ".go") ||
							   strings.HasSuffix(linkTarget, ".png") ||
							   strings.HasSuffix(linkTarget, ".jpg")

			if !isPotentialPath {
				continue
			}
			
			// Limpar query params ou âncoras no final do link
			linkTarget = strings.Split(linkTarget, "#")[0]
			linkTarget = strings.Split(linkTarget, "?")[0]

			if linkTarget == "" {
				continue
			}

			// Caminho absoluto do alvo
			targetPath := filepath.Join(filepath.Dir(path), linkTarget)
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				color.Red("   ❌ Link quebrado em %s: -> %s", rel, linkTarget)
				errorsFound = true
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	if errorsFound {
		return fmt.Errorf("broken links detected in documentation")
	}

	color.Green("   ✅ Todos os links internos estão funcionais.")
	return nil
}

// CheckSemanticState cruza informações do STATE.md com a existência real de arquivos/diretórios
func CheckSemanticState(root string) error {
	color.Blue("\n🧠 Iniciando Validação Semântica de Estado...")

	statePath := filepath.Join(root, ".specs", "project", "STATE.md")
	content, err := os.ReadFile(statePath)
	if err != nil {
		return fmt.Errorf("STATE.md not found at %s", statePath)
	}

	// Regex para identificar features concluídas: - [x] **Nome da Feature**
	re := regexp.MustCompile(`- \[x\] \*\*([^*]+)\*\*`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	if len(matches) == 0 {
		color.Yellow("   ⚠️  Nenhuma feature concluída encontrada para validação no STATE.md")
		return nil
	}

	errorsFound := false
	for _, match := range matches {
		featureName := strings.TrimSpace(match[1])
		
		// Heurística de validação baseada no nome
		// 1. Se contém "Skill", procurar pasta correspondente
		if strings.Contains(strings.ToLower(featureName), "skill") || 
		   strings.Contains(strings.ToLower(featureName), "expert") ||
		   strings.Contains(strings.ToLower(featureName), "enrichment") {
			
			// Tentar extrair o slug (ex: "Golang Expert" -> "golang-expert")
			slug := strings.ToLower(featureName)
			slug = strings.Split(slug, " enrichment")[0]
			slug = strings.Split(slug, " skill")[0]
			slug = strings.Split(slug, " enrichment")[0]
			slug = strings.ReplaceAll(slug, " ", "-")
			slug = strings.TrimSpace(slug)

			// Verificar se a pasta existe
			path := filepath.Join(root, slug)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				// Busca fuzzy simples nas pastas do root
				found := false
				entries, _ := os.ReadDir(root)
				for _, e := range entries {
					if e.IsDir() && strings.Contains(strings.ToLower(e.Name()), strings.Split(slug, "-")[0]) {
						found = true
						break
					}
				}
				
				if !found {
					color.Red("   ❌ Inconsistência Semântica: STATE.md indica '%s' como concluído, mas diretório '%s' não existe.", featureName, slug)
					errorsFound = true
				}
			} else {
				color.Cyan("   ✅ Validado: %s", featureName)
			}
		} else {
			// Apenas logar que foi detectado mas a heurística não se aplica
			color.Black("   🔍 Ignorado (Heurística n/a): %s", featureName)
		}
	}

	if errorsFound {
		return fmt.Errorf("semantic state mismatch detected")
	}

	color.Green("   ✅ Estado semântico consistente com o File System.")
	return nil
}

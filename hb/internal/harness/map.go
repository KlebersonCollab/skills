package harness

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// ExecuteKnowledgeSync identifica e sugere atualizações para o mapa de conhecimento
func ExecuteKnowledgeSync(root string, apply bool) error {
	color.Blue("🔍 Sincronizando Mapa de Conhecimento...")

	// 1. Obter arquivos modificados
	modified, err := getGitModifiedFiles(root)
	if err != nil {
		return err
	}

	// 2. Identificar Features modificadas
	features := identifyModifiedFeatures(modified)
	if len(features) == 0 {
		color.Yellow("Nenhuma feature nova ou modificada detectada para mapeamento.")
		return nil
	}

	// 3. Analisar relações faltantes
	mermaidPath := filepath.Join(root, "KNOWLEDGE-MAP.mermaid")
	currentMermaid, _ := os.ReadFile(mermaidPath)
	mermaidContent := string(currentMermaid)

	newRelations := []string{}
	for _, f := range features {
		links, _ := discoverFeatureLinks(root, f)
		for _, link := range links {
			rel := fmt.Sprintf("%s --(implements)--> %s", f, link)
			if !strings.Contains(mermaidContent, rel) {
				newRelations = append(newRelations, rel)
			}
		}
	}

	if len(newRelations) == 0 {
		color.Green("✅ O mapa de conhecimento já está atualizado com as mudanças recentes.")
		return nil
	}

	color.Cyan("\nNovas relações sugeridas:")
	for _, r := range newRelations {
		fmt.Printf(" - %s\n", r)
	}

	if apply {
		return applyNewRelations(mermaidPath, newRelations)
	}

	color.Yellow("\nUse --apply para atualizar o KNOWLEDGE-MAP.mermaid automaticamente.")
	return nil
}

func getGitModifiedFiles(root string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "HEAD")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		// Se falhar (talvez sem commits), tenta diff com staged
		cmd = exec.Command("git", "diff", "--cached", "--name-only")
		cmd.Dir = root
		out, _ = cmd.Output()
	}
	
	lines := strings.Split(string(out), "\n")
	
	// Também pega arquivos não trackeados
	cmd = exec.Command("git", "ls-files", "--others", "--exclude-standard")
	cmd.Dir = root
	out2, _ := cmd.Output()
	lines = append(lines, strings.Split(string(out2), "\n")...)

	var files []string
	for _, l := range lines {
		if l != "" {
			files = append(files, l)
		}
	}
	return files, nil
}

func identifyModifiedFeatures(files []string) []string {
	featureMap := make(map[string]bool)
	for _, f := range files {
		if strings.HasPrefix(f, ".specs/features/") {
			parts := strings.Split(f, "/")
			if len(parts) > 2 {
				featureMap[parts[2]] = true
			}
		}
	}
	
	var list []string
	for f := range featureMap {
		list = append(list, f)
	}
	return list
}

func discoverFeatureLinks(root, feature string) ([]string, error) {
	featureDir := filepath.Join(root, ".specs", "features", feature)
	files := []string{"spec.md", "plan.md"}
	
	links := make(map[string]bool)
	
	// Lista de todas as skills para cruzamento
	// No nosso caso, as skills estão na raiz ou em .agent, .claude, etc.
	// Mas o padrão é o nome da skill ser o diretório.
	
	for _, filename := range files {
		content, err := os.ReadFile(filepath.Join(featureDir, filename))
		if err != nil {
			continue
		}
		
		// Procura menções a skills conhecidas (ex: harness-expert)
		// Simplificação: qualquer palavra com hífen que pareça uma skill
		words := strings.Fields(string(content))
		for _, w := range words {
			w = strings.Trim(w, " `(),.[]\"'")
			if strings.Contains(w, "-") && len(w) > 5 && !strings.HasPrefix(w, "-") {
				// Ignora extensões de arquivos comuns
				if !strings.HasSuffix(w, ".md") && !strings.HasSuffix(w, ".go") {
					links[w] = true
				}
			}
		}
	}

	var list []string
	for l := range links {
		// Normaliza para o formato do mermaid (snake_case IDs)
		list = append(list, strings.ReplaceAll(l, "-", "_"))
	}
	return list, nil
}

func applyNewRelations(path string, relations []string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, r := range relations {
		_, err := f.WriteString(fmt.Sprintf("    %s\n", r))
		if err != nil {
			return err
		}
	}

	color.Green("✅ Mapa de conhecimento atualizado!")
	return nil
}

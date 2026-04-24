package syncer

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/fatih/color"
)

const (
	MarkerStart = "<!-- GLOBAL_MANDATES_START -->"
	MarkerEnd   = "<!-- GLOBAL_MANDATES_END -->"
)

var agentFolders = []string{".gemini", ".claude", ".agent", ".agents"}

// SyncMandates injeta o conteúdo de GLOBAL_MANDATES.md nos arquivos dos agentes
func SyncMandates(root string) error {
	mandatesPath := filepath.Join(root, ".specs", "codebase", "GLOBAL_MANDATES.md")

	// Se não existir, tenta criar pasta
	os.MkdirAll(filepath.Dir(mandatesPath), 0755)

	mandates, err := os.ReadFile(mandatesPath)
	if err != nil {
		return fmt.Errorf("failed to read GLOBAL_MANDATES.md: %w", err)
	}

	newBlock := fmt.Sprintf("%s\n\n%s\n\n%s", MarkerStart, string(mandates), MarkerEnd)
	re := regexp.MustCompile(fmt.Sprintf("(?s)%s.*?%s", regexp.QuoteMeta(MarkerStart), regexp.QuoteMeta(MarkerEnd)))

	for _, folder := range agentFolders {
		agentDir := filepath.Join(root, folder)
		// Garantir que a pasta do agente existe
		os.MkdirAll(agentDir, 0755)

		fileName := strings.ToUpper(strings.TrimPrefix(folder, ".")) + ".md"
		if folder == ".agents" {
			fileName = "AGENTS.md"
		}
		agentFile := filepath.Join(agentDir, fileName)

		content, err := os.ReadFile(agentFile)
		var newContent []byte
		if err == nil && re.Match(content) {
			newContent = re.ReplaceAll(content, []byte(newBlock))
		} else {
			header := fmt.Sprintf("# %s Project Instructions\n\n", strings.ToUpper(strings.TrimPrefix(folder, ".")))
			newContent = []byte(header + newBlock + "\n\n--- \n*Mandate updated via HB-CLI.*")
		}

		err = os.WriteFile(agentFile, newContent, 0644)
		if err != nil {
			color.Yellow("   ! Warning: Failed to write to %s: %v", agentFile, err)
		} else {
			color.Cyan("   + Mandates synced: %s", agentFile)
		}
	}
	return nil
}

// SyncSkills sincroniza as pastas de skills para as pastas de agentes em paralelo
func SyncSkills(root string) error {
	skillPaths := make(map[string]string) // Name -> Absolute Path

	// 1. Busca no Root
	entries, _ := os.ReadDir(root)
	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			p := filepath.Join(root, e.Name())
			if _, err := os.Stat(filepath.Join(p, "SKILL.md")); err == nil {
				skillPaths[e.Name()] = p
			}
		}
	}

	// 2. Busca em .agents/skills (caso o usuário tenha instalado via CLI)
	agentsSkillsPath := filepath.Join(root, ".agents", "skills")
	if entries, err := os.ReadDir(agentsSkillsPath); err == nil {
		for _, e := range entries {
			if e.IsDir() {
				p := filepath.Join(agentsSkillsPath, e.Name())
				if _, err := os.Stat(filepath.Join(p, "SKILL.md")); err == nil {
					skillPaths[e.Name()] = p
				}
			}
		}
	}

	var wg sync.WaitGroup
	for _, folder := range agentFolders {
		agentSkillsDir := filepath.Join(root, folder, "skills")
		os.MkdirAll(agentSkillsDir, 0755)

		for name, path := range skillPaths {
			wg.Add(1)
			go func(sName, sPath, target string) {
				defer wg.Done()
				dst := filepath.Join(target, sName)
				err := CopyDir(sPath, dst)
				if err != nil {
					color.Red("   ! Error syncing %s to %s: %v", sName, target, err)
				}
			}(name, path, agentSkillsDir)
		}
	}
	wg.Wait()
	color.Green("   + All skills synced in parallel.")
	return nil
}

// CopyDir realiza a cópia recursiva de um diretório
func CopyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if entry.Name() == "__pycache__" || entry.Name() == ".pytest_cache" {
				continue
			}
			err = CopyDir(sourcePath, destPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourcePath, destPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync()
}

// DownloadMandates baixa os mandatos globais do repositório remoto
func DownloadMandates(root string) error {
	repoURL := "https://raw.githubusercontent.com/KlebersonCollab/skills/main/.specs/codebase/GLOBAL_MANDATES.md"
	color.Cyan("   -> Baixando mandatos globais de %s...", repoURL)

	mandatesPath := filepath.Join(root, ".specs", "codebase", "GLOBAL_MANDATES.md")
	os.MkdirAll(filepath.Dir(mandatesPath), 0755)

	cmd := exec.Command("curl", "-s", "-o", mandatesPath, repoURL)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to download mandates: %v\nOutput: %s", err, string(out))
	}

	color.Green("   ✅ Mandatos globais baixados com sucesso!")
	return nil
}

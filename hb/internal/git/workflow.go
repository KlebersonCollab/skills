package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Start(feature string) error {
	slug := strings.ToLower(strings.ReplaceAll(feature, " ", "-"))
	branch := fmt.Sprintf("feat/%s", slug)
	
	cmd := exec.Command("git", "checkout", "-b", branch)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao criar branch: %s", string(out))
	}

	color.Green("🌿 Branch '%s' criada e ativa.", branch)
	return nil
}

func Commit(mType, scope, message string) error {
	commitMsg := fmt.Sprintf("%s(%s): %s", mType, scope, message)
	
	// Primeiro faz o add .
	exec.Command("git", "add", ".").Run()

	cmd := exec.Command("git", "commit", "-m", commitMsg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao fazer commit: %s", string(out))
	}

	color.Green("✅ Commit realizado: %s", commitMsg)
	return nil
}

func Sync() error {
	color.Blue("🔄 Sincronizando com rebase...")
	cmd := exec.Command("git", "pull", "--rebase", "origin", "main")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro no sync/rebase: %s", string(out))
	}

	color.Green("✨ Repositório sincronizado com a main.")
	return nil
}

package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/domain"
	"github.com/klebersonromero/hb/internal/scanner"
	"github.com/klebersonromero/hb/internal/syncer"
	"gopkg.in/yaml.v3"
)

type Options struct {
	Name   string
	Target string
	Force  bool
	Remote bool
	Root   string
}

// Install orquestra a instalação local ou remota e lida com dependências
func Install(opts Options) error {
	// Mapa para evitar loops infinitos de dependências
	processed := make(map[string]bool)
	return installRecursive(opts, processed)
}

func installRecursive(opts Options, processed map[string]bool) error {
	if processed[opts.Name] {
		return nil
	}
	processed[opts.Name] = true

	var err error
	if opts.Remote {
		err = installRemote(opts)
	} else {
		err = installLocal(opts)
	}

	if err != nil {
		return err
	}

	// 1. Verificar dependências (uses)
	skillPath := filepath.Join(opts.Target, opts.Name)
	skillFile := filepath.Join(skillPath, "SKILL.md")

	// Criar uma skill temporária para o parser
	tempSkill := domain.Skill{Name: opts.Name}
	parseSkillMetadata(skillFile, &tempSkill)

	if len(tempSkill.Uses) > 0 {
		color.Blue("   🔗 Resolvendo dependências para %s: %v", opts.Name, tempSkill.Uses)
		for _, dep := range tempSkill.Uses {
			depOpts := opts
			depOpts.Name = dep
			err = installRecursive(depOpts, processed)
			if err != nil {
				color.Yellow("   ⚠️  Falha ao instalar dependência %s: %v", dep, err)
			}
		}
	}

	return nil
}

// Update força a reinstalação de uma ou todas as skills
func Update(opts Options, all bool) error {
	opts.Force = true // Forçar sobrescrita

	if all {
		color.Blue("🔄 Atualizando todas as skills instaladas em %s...", opts.Target)
		entries, err := os.ReadDir(opts.Target)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				skillName := entry.Name()
				color.Cyan("   -> Atualizando %s...", skillName)
				updateOpts := opts
				updateOpts.Name = skillName
				Install(updateOpts)
			}
		}
		return nil
	}

	color.Blue("🔄 Atualizando skill: %s", opts.Name)
	return Install(opts)
}

func installLocal(opts Options) error {
	skills, err := scanner.FindSkills(opts.Root)
	if err != nil {
		return err
	}

	var srcPath string
	for _, s := range skills {
		if s.Name == opts.Name || filepath.Base(s.Path) == opts.Name {
			srcPath = s.Path
			break
		}
	}

	if srcPath == "" {
		return fmt.Errorf("skill '%s' not found in local Hub. Try using --remote", opts.Name)
	}

	destPath := filepath.Join(opts.Target, filepath.Base(srcPath))

	if _, err := os.Stat(destPath); err == nil && !opts.Force {
		return fmt.Errorf("skill already exists at %s. Use --force to overwrite", destPath)
	}

	color.Cyan("   -> Copying from %s to %s", srcPath, destPath)
	return syncer.CopyDir(srcPath, destPath)
}

func installRemote(opts Options) error {
	repoURL := "https://github.com/KlebersonCollab/skills.git"
	tmpDir, err := os.MkdirTemp("", "hb-install-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	color.Cyan("   -> Cloning remote repository (optimized)...")
	cloneCmd := exec.Command("git", "clone", "--filter=blob:none", "--sparse", repoURL, tmpDir)
	if out, err := cloneCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git clone failed: %v\nOutput: %s", err, string(out))
	}

	color.Cyan("   -> Fetching skill via Sparse-Checkout...")
	sparseCmd := exec.Command("git", "-C", tmpDir, "sparse-checkout", "add", opts.Name, ".agents/skills/"+opts.Name, ".agent/skills/"+opts.Name)
	if out, err := sparseCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git sparse-checkout failed: %v\nOutput: %s", err, string(out))
	}

	// Localizar a skill no repo clonado
	possiblePaths := []string{
		filepath.Join(tmpDir, opts.Name),
		filepath.Join(tmpDir, ".agents", "skills", opts.Name),
		filepath.Join(tmpDir, ".agent", "skills", opts.Name),
	}

	var srcSkill string
	for _, p := range possiblePaths {
		if _, err := os.Stat(filepath.Join(p, "SKILL.md")); err == nil {
			srcSkill = p
			break
		}
	}

	if srcSkill == "" {
		return fmt.Errorf("skill '%s' not found in remote repository", opts.Name)
	}

	destPath := filepath.Join(opts.Target, opts.Name)
	if _, err := os.Stat(destPath); err == nil && !opts.Force {
		return fmt.Errorf("skill already exists at %s. Use --force to overwrite", destPath)
	}

	color.Cyan("   -> Installing %s in %s...", opts.Name, destPath)
	return syncer.CopyDir(srcSkill, destPath)
}

// Duplicação da lógica de parsing para evitar circular import
func parseSkillMetadata(path string, skill *domain.Skill) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return
	}

	var meta struct {
		Uses []string `yaml:"uses"`
	}
	_ = yaml.Unmarshal([]byte(parts[1]), &meta)
	skill.Uses = meta.Uses
}

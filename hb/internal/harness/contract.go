package harness

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// ContractIssue representa um desvio do contrato
type ContractIssue struct {
	Type    string // "Missing File", "Missing Symbol"
	Name    string
	Target  string
	Status  string // "✅ Found", "❌ Missing"
}

// ExecuteContractCheck verifica se a implementação condiz com o contract.md
func ExecuteContractCheck(root, feature string) (bool, error) {
	if feature == "" {
		return false, fmt.Errorf("nome da feature é obrigatório")
	}

	contractPath := filepath.Join(root, ".specs", "features", feature, "contract.md")
	if _, err := os.Stat(contractPath); os.IsNotExist(err) {
		return false, fmt.Errorf("contract.md não encontrado para a feature '%s'", feature)
	}

	color.Blue("🔍 Validando Conformidade Arquitetural (Contract Guardian) para: %s", feature)

	// 1. Extrair contratos do Markdown
	files, symbols, err := parseContractMD(contractPath)
	if err != nil {
		return false, err
	}

	issues := 0
	
	// 2. Validar Arquivos
	fmt.Println("\n📂 Verificando Artefatos:")
	for _, f := range files {
		found := false
		// Busca o arquivo no projeto
		_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() { return err }
			if d.Name() == f {
				found = true
				return filepath.SkipDir
			}
			return nil
		})
		
		if found {
			color.Green("   ✅ %s", f)
		} else {
			color.Red("   ❌ %s (Não encontrado)", f)
			issues++
		}
	}

	// 3. Validar Símbolos/Interfaces
	fmt.Println("\n🔗 Verificando Interfaces/Comandos:")
	for _, s := range symbols {
		// Busca o símbolo via grep
		found := searchSymbol(root, s)
		if found {
			color.Green("   ✅ %s", s)
		} else {
			color.Red("   ❌ %s (Implementação não detectada)", s)
			issues++
		}
	}

	if issues == 0 {
		color.Green("\n✨ Conformidade absoluta detectada! O código respeita o contrato.")
		return true, nil
	}

	color.Yellow("\n⚠️  Detectado drift de contrato (%d itens divergentes).", issues)
	return false, nil
}

func parseContractMD(path string) ([]string, []string, error) {
	var files []string
	var symbols []string

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	reSymbol := regexp.MustCompile("`([^`]+)`")
	reFile := regexp.MustCompile(`\*\*File\*\*:\s*` + "`([^`]+)`")

	for scanner.Scan() {
		line := scanner.Text()
		
		// 1. Tentar encontrar definições de arquivos **File**: `name`
		if matches := reFile.FindStringSubmatch(line); len(matches) > 1 {
			files = append(files, matches[1])
			continue
		}

		// 2. Se for uma linha de tabela, extrair símbolos em backticks
		if strings.Contains(line, "|") {
			matches := reSymbol.FindAllStringSubmatch(line, -1)
			for _, m := range matches {
				if len(m) > 1 {
					s := m[1]
					// Ignorar coisas que parecem flags ou paths genéricos
					if !strings.HasPrefix(s, "-") && !strings.Contains(s, "[") {
						symbols = append(symbols, s)
					}
				}
			}
		}
	}

	return files, symbols, scanner.Err()
}

func searchSymbol(root, symbol string) bool {
	// Se for um comando CLI, tenta buscar o nome dele ou partes chaves
	cleanSymbol := symbol
	if strings.HasPrefix(symbol, "hb ") {
		parts := strings.Split(symbol, " ")
		cleanSymbol = parts[len(parts)-1] // Pega o último termo (o comando)
	}

	// Executa um grep simples ignorando binários e pastas técnicas
	// Buscamos pelo símbolo em aspas ou após nomes de comandos
	cmd := exec.Command("grep", "-rE", `"`+cleanSymbol+`"|Use:\s*"`+cleanSymbol+`(\s|")|func\s+`+cleanSymbol, root, 
		"--exclude-dir=.git", "--exclude-dir=node_modules", "--exclude-dir=bin", "--exclude-dir=.specs")
	
	err := cmd.Run()
	return err == nil // Se exit code for 0, encontrou
}

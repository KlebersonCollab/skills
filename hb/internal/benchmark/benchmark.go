package benchmark

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Result struct {
	Target string
	NsPerOp float64
	Date   string
}

// Run executa os benchmarks para o target especificado
func Run(root, target string) (*Result, error) {
	color.Blue("🚀 Iniciando benchmarks para: %s", target)

	// Detectar linguagem
	isGo := false
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
		isGo = true
	} else if _, err := os.Stat(filepath.Join(root, "hb", "go.mod")); err == nil {
		isGo = true
	}

	if !isGo {
		return nil, fmt.Errorf("linguagem não suportada ou go.mod não encontrado")
	}

	cmd := exec.Command("go", "test", "-bench=.", "-benchmem", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("erro ao executar benchmark: %s", string(out))
	}

	return parseGoBenchmark(target, string(out))
}

func parseGoBenchmark(target, output string) (*Result, error) {
	// Exemplo: BenchmarkFib10-8  3000000  424 ns/op
	re := regexp.MustCompile(`Benchmark\S+\s+\d+\s+([\d\.]+)\s+ns/op`)
	matches := re.FindStringSubmatch(output)

	if len(matches) < 2 {
		return nil, fmt.Errorf("nenhum benchmark encontrado na saída")
	}

	val, _ := strconv.ParseFloat(matches[1], 64)
	return &Result{
		Target:  target,
		NsPerOp: val,
		Date:    time.Now().Format("2006-01-02"),
	}, nil
}

// Compare compara o resultado atual com o baseline e retorna a variação (%)
func Compare(current, baseline *Result) float64 {
	if baseline == nil || baseline.NsPerOp == 0 {
		return 0
	}
	return ((current.NsPerOp - baseline.NsPerOp) / baseline.NsPerOp) * 100
}

// LoadBaseline busca o último benchmark no LEARNINGS.md
func LoadBaseline(root, target string) (*Result, error) {
	path := filepath.Join(root, ".specs", "project", "LEARNINGS.md")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lastRes *Result
	scanner := bufio.NewScanner(file)
	// Procurar por padrão: [BENCHMARK] <target>: <val> ns/op (<date>)
	pattern := fmt.Sprintf(`\[BENCHMARK\] %s: ([\d\.]+) ns/op \(([\d-]+)\)`, regexp.QuoteMeta(target))
	re := regexp.MustCompile(pattern)

	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) >= 3 {
			val, _ := strconv.ParseFloat(matches[1], 64)
			lastRes = &Result{
				Target:  target,
				NsPerOp: val,
				Date:    matches[2],
			}
		}
	}

	return lastRes, nil
}

// SaveBaseline persiste o resultado no LEARNINGS.md
func SaveBaseline(root, target string, res *Result) error {
	path := filepath.Join(root, ".specs", "project", "LEARNINGS.md")
	
	entry := fmt.Sprintf("\n### [BENCHMARK] %s: %.2f ns/op (%s)\n", target, res.NsPerOp, res.Date)
	
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(entry)
	return err
}

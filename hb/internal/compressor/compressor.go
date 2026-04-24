package compressor

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var preservePatterns = []string{
	"```[\\s\\S]*?```",  // Code blocks
	"`[^`]*?`",          // Inline code
	"\\[.*?\\]\\(.*?\\)", // Links
	"^#+ .*$",           // Headings
	"\\|.*?\\|",          // Tables
	"https?://\\S+",     // Raw URLs
}

var fillerWords = []string{
	"just", "really", "basically", "actually", "simply", "very", "of course",
	"happy to", "sure", "certainly", "please", "kindly",
	"realmente", "simplesmente", "basicamente", "apenas", "só", "com certeza",
}

var articles = []string{
	"the", "a", "an", "o", "as", "os", "um", "uma", "uns", "umas", "de", "do", "da", "dos", "das",
}

// CompressFile comprime um arquivo Markdown e salva um backup .original.md se necessário
func CompressFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	originalSize := len(content)
	compressed := CompressText(string(content))
	compressedSize := len(compressed)

	// Backup se for a primeira vez
	backupPath := path + ".original.md"
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		os.WriteFile(backupPath, content, 0644)
		color.Cyan("   + Backup criado: %s", filepathBase(backupPath))
	}

	err = os.WriteFile(path, []byte(compressed), 0644)
	if err != nil {
		return err
	}

	reduction := 100 - (float64(compressedSize) / float64(originalSize) * 100)
	color.Green("   ✅ %s comprimido: %d -> %d bytes (%.1f%% salvo)", filepathBase(path), originalSize, compressedSize, reduction)
	return nil
}

// CompressText aplica a minificação Caveman no texto
func CompressText(text string) string {
	placeholderMap := make(map[string]string)
	counter := 0

	// 1. Proteger elementos técnicos
	combinedPreserve := strings.Join(preservePatterns, "|")
	rePreserve := regexp.MustCompile("(?m)" + combinedPreserve)

	text = rePreserve.ReplaceAllStringFunc(text, func(match string) string {
		placeholder := fmt.Sprintf("__PRESERVE_%d__", counter)
		placeholderMap[placeholder] = match
		counter++
		return placeholder
	})

	// 2. Compressão Linguística
	// Remover fillers e artigos (case insensitive)
	allRemovables := append(fillerWords, articles...)
	for _, word := range allRemovables {
		re := regexp.MustCompile("(?i)\\b" + regexp.QuoteMeta(word) + "\\b")
		text = re.ReplaceAllString(text, "")
	}

	// 3. Limpeza de espaços
	reSpace := regexp.MustCompile(` +`)
	text = reSpace.ReplaceAllString(text, " ")
	
	reNewlines := regexp.MustCompile(`\n{3,}`)
	text = reNewlines.ReplaceAllString(text, "\n\n")
	
	text = strings.TrimSpace(text)

	// 4. Restaurar elementos técnicos
	for placeholder, original := range placeholderMap {
		text = strings.ReplaceAll(text, placeholder, original)
	}

	return text
}

func filepathBase(path string) string {
	parts := strings.Split(path, string(os.PathSeparator))
	return parts[len(parts)-1]
}

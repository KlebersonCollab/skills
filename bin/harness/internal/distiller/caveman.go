// Package distiller provides token compression and text distillation for agent history.
package distiller

import (
	"fmt"
	"regexp"
	"strings"
)

// Level represents the intensity of Caveman compression.
type Level string

const (
	Lite Level = "lite"
	Full Level = "full"
	Ultra Level = "ultra"
)

var (
	fillerRE       = regexp.MustCompile(`(?i)\b(really|basically|just|actually|simply|very|of course|absolutely|sure|certainly|obviously|clearly)\b\s*`)
	articleRE      = regexp.MustCompile(`(?i)\b(the|a|an)\b\s*`)
	conjunctionRE  = regexp.MustCompile(`(?i)\b(and|or|but|because|since|although)\b\s*`)
	causalityRE    = regexp.MustCompile(`(?i)\s+\b(leads to|causes|results in)\b\s+`)
	stepPattern    = regexp.MustCompile(`(?im)^\s*(\d+[\.:]\s+|-\s*(step|phase)\s+\d+|(step|phase)\s+\d+[\.:]?\s+)`)
	securityKW     = []string{"security warning", "vulnerability", "cve-", "exploit", "secret", "private key", "auth bypass", "sql injection", "unsafe"}
	irreversibleKW = []string{"permanent", "cannot be undone", "drop table", "delete all", "purge database", "destructive", "rm -rf", "irreversible"}
)

// ToCaveman compresses text according to the specified level.
// Returns original text with bypass marker if Auto-Clarity triggers.
func ToCaveman(text string, level Level) string {
	if CheckAutoClarity(text) {
		return "[AUTO-CLARITY TRIGGERED: Safety Valve Active - Bypassing Compression]\n" + text
	}

	text = fillerRE.ReplaceAllString(text, "")

	if level != Lite {
		text = articleRE.ReplaceAllString(text, "")
		text = strings.NewReplacer(
			"implement a solution for", "fix",
			"implement", "do",
			"extensive", "big",
		).Replace(text)
	}

	if level == Ultra {
		text = abbreviateText(text)
	}

	return strings.TrimSpace(text)
}

// CheckAutoClarity implements the Safety Valve.
// Returns true if text MUST NOT be compressed (security, irreversible, multi-step).
func CheckAutoClarity(text string) bool {
	lower := strings.ToLower(text)
	for _, kw := range securityKW {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	for _, kw := range irreversibleKW {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return stepPattern.MatchString(text)
}

// IsDestructiveCommand checks if command matches known destructive patterns.
func IsDestructiveCommand(cmd string) bool {
	lower := strings.ToLower(cmd)
	patterns := []string{
		"rm -rf", "rm -r /", "rm -rf /", "rm -fr",
		"dd if=", "mkfs.", "mkfs ",
		"drop table", "drop database",
		"truncate table",
		"purge", "shutdown",
		":(){ :|:& };:",
		"> /dev/", "> /dev/sda", "> /dev/nvme",
		"chmod 777 /", "chown -r",
		"mount --bind",
		"wget ", "curl ",
	}
	for _, p := range patterns {
		if strings.Contains(lower, p) {
			return true
		}
	}
	return false
}

// CompactHistory performs micro-compaction on old tool outputs in a node slice.
// Nodes older than maxAgeTurns and larger than maxCharLen are replaced with stubs.
func CompactHistory[T any](nodes []T, maxAgeTurns int, maxCharLen int, getContent func(T) (string, func(T, string) T)) []T {
	if maxAgeTurns <= 0 {
		maxAgeTurns = 10
	}
	if maxCharLen <= 0 {
		maxCharLen = 5000
	}

	total := len(nodes)
	compacted := make([]T, 0, total)

	for i, n := range nodes {
		content, setter := getContent(n)
		if len(content) > maxCharLen && (total-i) > maxAgeTurns {
			originalLen := len(content)
			n = setter(n, fmt.Sprintf("[Old tool result cleared. Tool: system. Summary: Long output (%d bytes).]", originalLen))
		}
		compacted = append(compacted, n)
	}
	return compacted
}

// abbreviateText applies ultra-level abbreviations, protecting backticked code blocks.
func abbreviateText(text string) string {
	parts := strings.Split(text, "`")
	for i := 0; i < len(parts); i += 2 {
		parts[i] = causalityRE.ReplaceAllString(parts[i], " → ")
		replacer := strings.NewReplacer(
			"database", "DB",
			"databases", "DBs",
			"authentication", "auth",
			"configuration", "config",
			"request", "req",
			"requests", "reqs",
			"response", "res",
			"responses", "res",
			"function", "fn",
			"functions", "fns",
			"implementation", "impl",
			"connection", "conn",
			"connections", "conns",
		)
		parts[i] = replacer.Replace(parts[i])
		parts[i] = conjunctionRE.ReplaceAllString(parts[i], "")
	}
	return strings.Join(parts, "`")
}

// CompactHistorySimple is a simplified version of CompactHistory for session.Node.
type NodeCompactor struct {
	MaxAgeTurns int
	MaxCharLen  int
}

// Compact compacts a slice of nodes (interface{} with Role and Content fields).
type HasContent interface {
	GetContent() string
	SetContent(string)
	GetRole() string
}

// Compact compacts nodes in place, returning a new slice.
func Compact(nodes []HasContent, maxAgeTurns int, maxCharLen int) []HasContent {
	if maxAgeTurns <= 0 {
		maxAgeTurns = 10
	}
	if maxCharLen <= 0 {
		maxCharLen = 5000
	}

	total := len(nodes)
	compacted := make([]HasContent, 0, total)

	for i, n := range nodes {
		if n.GetRole() == "system" && len(n.GetContent()) > maxCharLen && (total-i) > maxAgeTurns {
			originalLen := len(n.GetContent())
			n.SetContent(fmt.Sprintf("[Old tool result cleared. Summary: Long output (%d bytes).]", originalLen))
		}
		compacted = append(compacted, n)
	}
	return compacted
}

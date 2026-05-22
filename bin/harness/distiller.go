package main

import (
	"fmt"
	"regexp"
	"strings"
)

// DistillLevel represents the intensity of Caveman compression.
type DistillLevel string

const (
	Lite  DistillLevel = "lite"
	Full  DistillLevel = "full"
	Ultra DistillLevel = "ultra"
)

var (
	// Filler words removed in all levels
	fillerRE = regexp.MustCompile(`(?i)\b(really|basically|just|actually|simply|very|of course|absolutely|sure|certainly|obviously|clearly)\b\s*`)

	// Articles removed in full and ultra
	articleRE = regexp.MustCompile(`(?i)\b(the|a|an)\b\s*`)

	// Conjunctions stripped in ultra
	conjunctionRE = regexp.MustCompile(`(?i)\b(and|or|but|because|since|although)\b\s*`)

	// Causality arrows for ultra
	causalityRE = regexp.MustCompile(`(?i)\s+\b(leads to|causes|results in)\b\s+`)

	// Multi-step patterns for Auto-Clarity
	stepPattern = regexp.MustCompile(`(?im)^\s*(\d+[\.:]\s+|-\s*(step|phase)\s+\d+|(step|phase)\s+\d+[\.:]?\s+)`)
)

// Security and irreversible keywords for Auto-Clarity
var securityKeywords = []string{
	"security warning", "vulnerability", "cve-", "exploit",
	"secret", "private key", "auth bypass", "sql injection", "unsafe",
}
var irreversibleKeywords = []string{
	"permanent", "cannot be undone", "drop table", "delete all",
	"purge database", "destructive", "rm -rf", "irreversible",
}

// abbreviateText applies ultra-level abbreviations, protecting backticked code blocks.
func abbreviateText(text string) string {
	parts := strings.Split(text, "`")
	for i := 0; i < len(parts); i += 2 {
		// Causality arrows
		parts[i] = causalityRE.ReplaceAllString(parts[i], " → ")

		// Prose abbreviations
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

		// Strip conjunctions
		parts[i] = conjunctionRE.ReplaceAllString(parts[i], "")
	}
	return strings.Join(parts, "`")
}

// CheckAutoClarity implements the Safety Valve.
// Returns true if the text MUST NOT be compressed (security, irreversible, multi-step).
func CheckAutoClarity(text string) bool {
	lower := strings.ToLower(text)

	for _, kw := range securityKeywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	for _, kw := range irreversibleKeywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	if stepPattern.MatchString(text) {
		return true
	}
	return false
}

// ToCaveman compresses the given text according to the specified level.
// Returns the original text unchanged if Auto-Clarity triggers.
func ToCaveman(text string, level DistillLevel) string {
	if CheckAutoClarity(text) {
		return "[AUTO-CLARITY TRIGGERED: Safety Valve Active - Bypassing Compression]\n" + text
	}

	// Remove filler words (all levels)
	text = fillerRE.ReplaceAllString(text, "")

	if level != Lite {
		// Drop articles
		text = articleRE.ReplaceAllString(text, "")

		// Short direct synonyms
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

// isDestructiveCommand checks if command matches known destructive patterns.
// Used by Safety Approval Gate before execution.
func isDestructiveCommand(cmd string) bool {
	lower := strings.ToLower(cmd)
	patterns := []string{
		"rm -rf", "rm -r /", "rm -rf /", "rm -fr",
		"dd if=", "mkfs.", "mkfs ",
		"drop table", "drop database",
		"truncate table",
		"purge", "shutdown",
		":(){ :|:& };:", // fork bomb
		"> /dev/", "> /dev/sda", "> /dev/nvme",
		"chmod 777 /", "chown -r",
		"mount --bind",
		"wget ", "curl ", // fetching external scripts is dangerous
	}
	for _, p := range patterns {
		if strings.Contains(lower, p) {
			return true
		}
	}
	return false
}

// CompactHistory performs micro-compaction on old tool outputs in the node list.
// Nodes older than maxAgeTurns and larger than maxCharLen are replaced with stubs.
func CompactHistory(nodes []Node, maxAgeTurns int, maxCharLen int) []Node {
	if maxAgeTurns <= 0 {
		maxAgeTurns = 10
	}
	if maxCharLen <= 0 {
		maxCharLen = 5000
	}

	total := len(nodes)
	compacted := make([]Node, 0, total)

	for i, n := range nodes {
		// Only compact system (tool) nodes
		if n.Role == "system" && len(n.Content) > maxCharLen && (total-i) > maxAgeTurns {
			originalLen := len(n.Content)
			summary := originalLen
			if summary > 200 {
				summary = 200
			}
			n.Content = fmt.Sprintf("[Old tool result cleared. Tool: system. Summary: Long tool output (%d bytes).]", originalLen)
			_ = summary // preserve original first bytes if needed
		}
		compacted = append(compacted, n)
	}
	return compacted
}

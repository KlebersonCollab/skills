package main

import (
	"fmt"
	"strings"
)

const (
	CReset = "\033[0m"
	CBold  = "\033[1m"
	CDim   = "\033[2m"
	CAccent   = "\033[38;5;99m"
	CYellow   = "\033[38;5;208m"
	CCyan     = "\033[38;5;86m"
)

func formatInline(res string) string {
	// Double asterisks: **bold** -> CYellow + Bold
	for {
		start := strings.Index(res, "**")
		if start < 0 {
			break
		}
		end := strings.Index(res[start+2:], "**")
		if end < 0 {
			break
		}
		endIdx := start + 2 + end
		boldText := res[start+2 : endIdx]
		res = res[:start] + CBold + CYellow + boldText + CReset + res[endIdx+2:]
	}

	// Single asterisks: *italic* -> CDim
	for {
		start := strings.Index(res, "*")
		if start < 0 {
			break
		}
		end := strings.Index(res[start+1:], "*")
		if end < 0 {
			break
		}
		endIdx := start + 1 + end
		italicText := res[start+1 : endIdx]
		res = res[:start] + CDim + italicText + CReset + res[endIdx+1:]
	}

	// Inline code: `code` -> CCyan
	for {
		start := strings.Index(res, "`")
		if start < 0 {
			break
		}
		end := strings.Index(res[start+1:], "`")
		if end < 0 {
			break
		}
		endIdx := start + 1 + end
		codeText := res[start+1 : endIdx]
		res = res[:start] + CCyan + codeText + CReset + res[endIdx+1:]
	}

	return res
}

func main() {
	testCases := []string{
		"**bold**",
		"*italic*",
		"`code`",
		"**bold** and *italic* and `code`",
		"**unclosed",
		"*unclosed",
		"`unclosed",
		"**bold *italic* bold**",
		"this * that * other",
		"nested **bold **within** bold**",
		"**a*b**",
		"*a**b*",
		"* O git log -p --all",
		"\"+     return Speak(spoken, vs.apiKey)\"",
		"| Arquivo | Item Órfão |",
		"---",
		"# Header",
	}

	fmt.Println("Running test cases...")
	for i, tc := range testCases {
		fmt.Printf("TC %d: %q\n", i, tc)
		formatted := formatInline(tc)
		fmt.Printf("   Result: %q\n", formatted)
	}
	fmt.Println("All test cases completed successfully (no infinite loops)!")
}

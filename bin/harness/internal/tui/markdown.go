package tui

import (
	"strings"
)

// RenderMarkdown parses markdown text and returns a beautiful terminal-formatted string
// with colored headers, elegant boxed tables, cards for code blocks, and bold/italic elements.
func RenderMarkdown(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inCodeBlock := false
	var codeLines []string

	// Table parsing state
	inTable := false
	var tableRows [][]string

	// Flush table helper
	flushTable := func() {
		if len(tableRows) == 0 {
			inTable = false
			return
		}
		// Calculate column widths
		numCols := 0
		for _, row := range tableRows {
			if len(row) > numCols {
				numCols = len(row)
			}
		}
		colWidths := make([]int, numCols)
		for _, row := range tableRows {
			for i, cell := range row {
				cleanCell := stripANSIAndMD(cell)
				if len(cleanCell) > colWidths[i] {
					colWidths[i] = len(cleanCell)
				}
			}
		}

		// Draw top border
		var top strings.Builder
		top.WriteString("  ┌")
		for i, w := range colWidths {
			top.WriteString(strings.Repeat("─", w+2))
			if i < len(colWidths)-1 {
				top.WriteString("┬")
			}
		}
		top.WriteString("┐")
		result = append(result, top.String())

		// Draw rows
		for rowIdx, row := range tableRows {
			// Check if separator line
			isSeparator := true
			for _, cell := range row {
				clean := strings.TrimSpace(cell)
				if clean != "" && !isCharsOnly(clean, '-') && !isCharsOnly(clean, ':') {
					isSeparator = false
					break
				}
			}
			if isSeparator {
				// Draw middle separator
				var sep strings.Builder
				sep.WriteString("  ├")
				for i, w := range colWidths {
					sep.WriteString(strings.Repeat("─", w+2))
					if i < len(colWidths)-1 {
						sep.WriteString("┼")
					}
				}
				sep.WriteString("┤")
				result = append(result, sep.String())
				continue
			}

			var rLine strings.Builder
			rLine.WriteString("  │")
			for i := 0; i < numCols; i++ {
				cellVal := ""
				if i < len(row) {
					cellVal = strings.TrimSpace(row[i])
				}
				cleanCell := stripANSIAndMD(cellVal)
				padding := colWidths[i] - len(cleanCell)

				// Highlight header row
				formattedCell := formatInline(cellVal)
				if rowIdx == 0 {
					formattedCell = CBold + CAccent + formattedCell + CReset
				} else {
					formattedCell = CYellow + formattedCell + CReset
				}

				rLine.WriteString(" ")
				rLine.WriteString(formattedCell)
				rLine.WriteString(strings.Repeat(" ", padding))
				rLine.WriteString(" │")
			}
			result = append(result, rLine.String())
		}

		// Draw bottom border
		var bottom strings.Builder
		bottom.WriteString("  └")
		for i, w := range colWidths {
			bottom.WriteString(strings.Repeat("─", w+2))
			if i < len(colWidths)-1 {
				bottom.WriteString("┴")
			}
		}
		bottom.WriteString("┘")
		result = append(result, bottom.String())

		tableRows = nil
		inTable = false
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// 1. Handle Code Blocks
		if strings.HasPrefix(trimmed, "```") {
			if inCodeBlock {
				// End of code block: draw elegant card
				result = append(result, "  ┌"+strings.Repeat("─", 60))
				for _, cl := range codeLines {
					result = append(result, "  │ "+CCyan+cl+CReset)
				}
				result = append(result, "  └"+strings.Repeat("─", 60))
				codeLines = nil
				inCodeBlock = false
			} else {
				// Start of code block
				inCodeBlock = true
			}
			continue
		}

		if inCodeBlock {
			codeLines = append(codeLines, line)
			continue
		}

		// 2. Handle Tables
		if strings.HasPrefix(trimmed, "|") && strings.HasSuffix(trimmed, "|") {
			inTable = true
			parts := strings.Split(trimmed, "|")
			// Remove first and last empty elements
			if len(parts) > 2 {
				parts = parts[1 : len(parts)-1]
			}
			tableRows = append(tableRows, parts)
			continue
		} else if inTable {
			flushTable()
		}

		// 3. Horizontal Rules
		if trimmed == "---" {
			result = append(result, "  "+CDarkGray+strings.Repeat("─", 64)+CReset)
			continue
		}

		// 4. Headers
		if strings.HasPrefix(trimmed, "# ") {
			headerText := strings.TrimPrefix(trimmed, "# ")
			result = append(result, "\n  "+CBold+CAccent+strings.ToUpper(headerText)+CReset)
			continue
		}
		if strings.HasPrefix(trimmed, "## ") {
			headerText := strings.TrimPrefix(trimmed, "## ")
			result = append(result, "\n  "+CBold+CCyan+headerText+CReset)
			continue
		}
		if strings.HasPrefix(trimmed, "### ") {
			headerText := strings.TrimPrefix(trimmed, "### ")
			result = append(result, "\n  "+CBold+CGreen+headerText+CReset)
			continue
		}

		// 5. Unordered lists
		if strings.HasPrefix(trimmed, "- ") {
			itemText := strings.TrimPrefix(trimmed, "- ")
			formatted := formatInline(itemText)
			result = append(result, "    "+CCyan+"•"+CReset+" "+formatted)
			continue
		}
		if strings.HasPrefix(trimmed, "* ") {
			itemText := strings.TrimPrefix(trimmed, "* ")
			formatted := formatInline(itemText)
			result = append(result, "    "+CCyan+"•"+CReset+" "+formatted)
			continue
		}

		// 6. Normal line
		formatted := formatInline(line)
		result = append(result, "  "+formatted)
	}

	if inTable {
		flushTable()
	}

	return strings.Join(result, "\n")
}

func isCharsOnly(s string, char byte) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != char && s[i] != ' ' && s[i] != ':' {
			return false
		}
	}
	return true
}

func stripANSIAndMD(s string) string {
	var clean strings.Builder
	inESC := false
	for i := 0; i < len(s); i++ {
		if s[i] == '\033' {
			inESC = true
			continue
		}
		if inESC {
			if s[i] == 'm' {
				inESC = false
			}
			continue
		}
		clean.WriteByte(s[i])
	}

	res := clean.String()
	res = strings.ReplaceAll(res, "**", "")
	res = strings.ReplaceAll(res, "*", "")
	res = strings.ReplaceAll(res, "`", "")
	return res
}

func formatInline(s string) string {
	res := s

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

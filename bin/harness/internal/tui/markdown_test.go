package tui

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	input := `# Header 1
Some normal text with **bold** and *italic* and ` + "`code`" + `.

- List item 1
- List item 2

| Header A | Header B |
|---|---|
| Value A | Value B |
`
	rendered := RenderMarkdown(input)

	// Check headers
	if !strings.Contains(rendered, "HEADER 1") {
		t.Error("expected Header 1 to be capitalized and present")
	}

	// Check bold
	if !strings.Contains(rendered, CBold) {
		t.Error("expected bold ANSI tags to be present")
	}

	// Check lists
	if !strings.Contains(rendered, "•") || !strings.Contains(rendered, "List item 1") {
		t.Error("expected bullet points to be formatted")
	}

	// Check tables
	if !strings.Contains(rendered, "┌") || !strings.Contains(rendered, "├") || !strings.Contains(rendered, "└") {
		t.Error("expected table borders to be rendered")
	}
	if !strings.Contains(rendered, "Header A") || !strings.Contains(rendered, "Value A") {
		t.Error("expected table values to be present")
	}
}

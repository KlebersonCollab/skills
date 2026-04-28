package git

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var hunkHeader = regexp.MustCompile(`^@@ -(\d+),?(\d*) \+(\d+),?(\d*) @@`)

// ParseDiff parses a unified diff string into individual file hunks.
func ParseDiff(diffStr string) ([]DiffHunk, error) {
	var hunks []DiffHunk
	var currentFile string
	
	scanner := bufio.NewScanner(strings.NewReader(diffStr))
	var currentHunk *DiffHunk

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "diff --git") {
			// New file detected
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				currentFile = strings.TrimPrefix(parts[3], "b/")
			}
			continue
		}

		if hunkHeader.MatchString(line) {
			if currentHunk != nil {
				hunks = append(hunks, *currentHunk)
			}
			
			matches := hunkHeader.FindStringSubmatch(line)
			oldStart, _ := strconv.Atoi(matches[1])
			oldLines, _ := strconv.Atoi(matches[2])
			newStart, _ := strconv.Atoi(matches[3])
			newLines, _ := strconv.Atoi(matches[4])

			currentHunk = &DiffHunk{
				FilePath: currentFile,
				OldStart: oldStart,
				OldLines: oldLines,
				NewStart: newStart,
				NewLines: newLines,
				Lines:    []string{line},
			}
			continue
		}

		if currentHunk != nil {
			currentHunk.Lines = append(currentHunk.Lines, line)
		}
	}

	if currentHunk != nil {
		hunks = append(hunks, *currentHunk)
	}

	return hunks, nil
}

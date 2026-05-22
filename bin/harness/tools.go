package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// resolvePath returns the absolute path within the workspace root
func resolvePath(p string) (string, error) {
	root, err := FindWorkspaceRoot()
	if err != nil {
		return "", err
	}
	if filepath.IsAbs(p) {
		// Clean the path and verify it is under root
		absClean := filepath.Clean(p)
		rootClean := filepath.Clean(root)
		if !strings.HasPrefix(absClean, rootClean) {
			return "", fmt.Errorf("security boundary violation: path '%s' is outside workspace root '%s'", p, root)
		}
		return absClean, nil
	}
	return filepath.Clean(filepath.Join(root, p)), nil
}

// ReadFile performs safe, bounded file reading (up to 5MB)
func ReadFile(path string) (string, error) {
	resolved, err := resolvePath(path)
	if err != nil {
		return "", err
	}

	file, err := os.Open(resolved)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 5MB limit
	const maxLimit = 5 * 1024 * 1024
	limitedReader := io.LimitReader(file, maxLimit)

	content, err := io.ReadAll(limitedReader)
	if err != nil {
		return "", err
	}

	// Check if file was truncated
	fi, err := file.Stat()
	if err == nil && fi.Size() > maxLimit {
		return string(content), fmt.Errorf("warning: file exceeded 5MB size limit and was truncated")
	}

	return string(content), nil
}

// WriteFile writes or completely overwrites a file (enforcing SDD gate)
func WriteFile(path string, content string) error {
	if err := VerifySDDGated(); err != nil {
		return err
	}

	resolved, err := resolvePath(path)
	if err != nil {
		return err
	}

	dir := filepath.Dir(resolved)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory tree: %w", err)
	}

	err = os.WriteFile(resolved, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// PatchFile surgically replaces a contiguous block of text (enforcing SDD gate)
func PatchFile(path string, target string, replacement string) error {
	if err := VerifySDDGated(); err != nil {
		return err
	}

	resolved, err := resolvePath(path)
	if err != nil {
		return err
	}

	bytes, err := os.ReadFile(resolved)
	if err != nil {
		return fmt.Errorf("failed to read file for patching: %w", err)
	}

	content := string(bytes)
	count := strings.Count(content, target)
	if count == 0 {
		return fmt.Errorf("patch failed: target text not found in '%s'", path)
	}
	if count > 1 {
		return fmt.Errorf("patch failed: target text matches %d times in '%s' (must be unique)", count, path)
	}

	newContent := strings.Replace(content, target, replacement, 1)
	err = os.WriteFile(resolved, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write patched file: %w", err)
	}

	return nil
}

// ExecuteCommand executes a subprocess command via system shell (enforcing SDD gate)
func ExecuteCommand(command string) (string, int, error) {
	if err := VerifySDDGated(); err != nil {
		return "", -1, err
	}

	root, err := FindWorkspaceRoot()
	if err != nil {
		return "", -1, err
	}

	// OS-agnostic command routing
	var shell, shellFlag string
	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		shellFlag = "/c"
	} else {
		shell = "bash"
		shellFlag = "-c"
	}

	cmd := exec.Command(shell, shellFlag, command)
	cmd.Dir = root

	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)

	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			return output, -1, err
		}
	}

	return output, exitCode, nil
}

// isBinary checks if a file is binary by reading its first 512 bytes.
// Returns true if a null byte is found (typical binary indicator).
func isBinary(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	buf := make([]byte, 512)
	n, err := io.ReadFull(f, buf)
	if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
		return false
	}
	buf = buf[:n]
	for _, b := range buf {
		if b == 0 {
			return true
		}
	}
	return false
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 64*1024) // 64KB buffer
		return &b
	},
}

// SearchFiles performs highly parallel, pure Go OS-agnostic search on files by name patterns and/or contents query.
// Uses a worker pool pattern: dispatcher walks the tree, workers read and match files concurrently.
func SearchFiles(pattern string, query string) (string, error) {
	root, err := FindWorkspaceRoot()
	if err != nil {
		return "", err
	}

	patternLower := strings.ToLower(pattern)
	queryLower := strings.ToLower(query)

	const maxMatches = 100
	results := make([]string, 0, maxMatches)
	resultsMu := sync.Mutex{}

	// Channel for file paths (buffered)
	jobs := make(chan string, 1024)

	// Worker pool
	numWorkers := runtime.NumCPU()
	if numWorkers < 2 {
		numWorkers = 2 // at least 2 workers for parallelism
	}
	var wg sync.WaitGroup

	type matchResult struct {
		lines []string
	}
	resultsChan := make(chan matchResult, numWorkers*2)

	// Worker function
	worker := func() {
		defer wg.Done()
		for path := range jobs {
			name := filepath.Base(path)
			nameLower := strings.ToLower(name)

			// Match filename pattern if provided
			if pattern != "" {
				match, err := filepath.Match(patternLower, nameLower)
				if err != nil || !match {
					continue
				}
			}

			// If only matching by name (no query), send result directly
			if query == "" {
				relPath, err := filepath.Rel(root, path)
				if err != nil {
					relPath = path
				}
				resultsChan <- matchResult{lines: []string{relPath}}
				continue
			}

			// Binary detection (only if we need to read content)
			if isBinary(path) {
				continue // skip binary files
			}

			// Open file
			f, err := os.Open(path)
			if err != nil {
				continue
			}

			fi, err := f.Stat()
			if err != nil {
				f.Close()
				continue
			}
			if fi.Size() > 2*1024*1024 {
				f.Close()
				continue // skip files > 2MB
			}

			// Use buffer pool for small files
			var contentBytes []byte
			if fi.Size() <= 64*1024 {
				bufPtr := bufferPool.Get().(*[]byte)
				buf := *bufPtr
				n, readErr := io.ReadFull(f, buf)
				f.Close()
				if readErr != nil && readErr != io.ErrUnexpectedEOF && readErr != io.EOF {
					bufferPool.Put(bufPtr)
					continue
				}
				contentBytes = make([]byte, n)
				copy(contentBytes, buf[:n])
				bufferPool.Put(bufPtr)
			} else {
				contentBytes, err = io.ReadAll(f)
				f.Close()
				if err != nil {
					continue
				}
			}

			contentStr := string(contentBytes)
			if !strings.Contains(strings.ToLower(contentStr), queryLower) {
				continue
			}

			// Extract matched lines
			lines := strings.Split(contentStr, "\n")
			var matchedLines []string
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				relPath = path
			}

			for idx, line := range lines {
				if strings.Contains(strings.ToLower(line), queryLower) {
					preview := strings.TrimSpace(line)
					if len(preview) > 100 {
						preview = preview[:97] + "..."
					}
					matchedLines = append(matchedLines, fmt.Sprintf("%s:%d: %s", relPath, idx+1, preview))
				}
			}

			if len(matchedLines) == 0 {
				matchedLines = append(matchedLines, fmt.Sprintf("%s: (matched in binary/unstructured format)", relPath))
			}

			resultsChan <- matchResult{lines: matchedLines}
		}
	}

	// Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// Dispatcher: walk directory tree
	dispatcherDone := make(chan error, 1)
	go func() {
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return nil // skip problematic paths
			}

			name := d.Name()
			// Skip hidden directories, build artifacts
			if d.IsDir() {
				if strings.HasPrefix(name, ".") && name != "." && name != ".." {
					return filepath.SkipDir
				}
				if name == "node_modules" || name == "dist" || name == "build" || name == "vendor" {
					return filepath.SkipDir
				}
				return nil
			}

			// Optional: pre-filter by name pattern to reduce jobs sent (only if both pattern and query)
			if pattern != "" && query == "" {
				// If only pattern, we already match in worker, but we can send all matching names to reduce jobs
				nameLower := strings.ToLower(name)
				match, _ := filepath.Match(patternLower, nameLower)
				if !match {
					return nil
				}
			}

			jobs <- path
			return nil
		})
		close(jobs)
		dispatcherDone <- err
	}()

	// Wait for all workers to finish, then close results channel
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collector: read from resultsChan and aggregate
	collectorWg := sync.WaitGroup{}
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		for res := range resultsChan {
			resultsMu.Lock()
			if len(results) >= maxMatches {
				resultsMu.Unlock()
				continue // skip extra results after cap
			}
			remaining := maxMatches - len(results)
			for _, line := range res.lines {
				if remaining <= 0 {
					break
				}
				results = append(results, line)
				remaining--
			}
			resultsMu.Unlock()
		}
	}()

	// Wait for dispatcher
	dispatchErr := <-dispatcherDone

	// Wait for collector to finish
	collectorWg.Wait()

	if dispatchErr != nil {
		return "", dispatchErr
	}

	if len(results) == 0 {
		return "Nenhum arquivo correspondente foi encontrado.", nil
	}

	// Log if truncated
	if cap(results) > maxMatches {
		truncatedCount := cap(results) - maxMatches
		results = append(results, fmt.Sprintf("... (e mais %d correspondências foram truncadas)", truncatedCount))
	}

	return strings.Join(results, "\n"), nil
}


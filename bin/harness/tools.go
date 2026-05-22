package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

// SearchFiles performs pure Go OS-agnostic search on files by name patterns and/or contents query
func SearchFiles(pattern string, query string) (string, error) {
	root, err := FindWorkspaceRoot()
	if err != nil {
		return "", err
	}

	var results []string
	patternLower := strings.ToLower(pattern)
	queryLower := strings.ToLower(query)

	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // ignore individual path errors to continue walk
		}

		name := d.Name()
		nameLower := strings.ToLower(name)

		// Performance: Skip system, harness or build dirs
		if d.IsDir() {
			if strings.HasPrefix(name, ".") && name != "." && name != ".." {
				return filepath.SkipDir
			}
			if name == "node_modules" || name == "dist" || name == "build" || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		// Match filename pattern if provided
		if pattern != "" {
			match, err := filepath.Match(patternLower, nameLower)
			if err != nil || !match {
				return nil
			}
		}

		// Match content query if provided
		if query != "" {
			file, err := os.Open(path)
			if err != nil {
				return nil
			}
			defer file.Close()

			fi, err := file.Stat()
			if err != nil || fi.Size() > 2*1024*1024 {
				return nil // Skip large files (2MB limit) to prevent token bloat/OOM
			}

			contentBytes, err := io.ReadAll(file)
			if err != nil {
				return nil
			}

			contentStr := string(contentBytes)
			if !strings.Contains(strings.ToLower(contentStr), queryLower) {
				return nil
			}

			// Extract matched lines with line numbers
			lines := strings.Split(contentStr, "\n")
			matchedInFile := false
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				relPath = path
			}

			for idx, line := range lines {
				if strings.Contains(strings.ToLower(line), queryLower) {
					matchedInFile = true
					preview := strings.TrimSpace(line)
					if len(preview) > 100 {
						preview = preview[:97] + "..."
					}
					results = append(results, fmt.Sprintf("%s:%d: %s", relPath, idx+1, preview))
				}
			}

			if !matchedInFile {
				results = append(results, fmt.Sprintf("%s: (matched in binary/unstructured format)", relPath))
			}
		} else {
			// Just matching the file name pattern
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				relPath = path
			}
			results = append(results, relPath)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "Nenhum arquivo correspondente foi encontrado.", nil
	}

	// Cap results to prevent token blowout
	const maxMatches = 100
	if len(results) > maxMatches {
		truncatedCount := len(results) - maxMatches
		results = append(results[:maxMatches], fmt.Sprintf("... (e mais %d correspondências foram truncadas)", truncatedCount))
	}

	return strings.Join(results, "\n"), nil
}


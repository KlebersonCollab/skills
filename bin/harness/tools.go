package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
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

// ExecuteCommand executes a subprocess command via bash (enforcing SDD gate)
func ExecuteCommand(command string) (string, int, error) {
	if err := VerifySDDGated(); err != nil {
		return "", -1, err
	}

	root, err := FindWorkspaceRoot()
	if err != nil {
		return "", -1, err
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = root

	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)

	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			return output, -1, err
		}
	}

	return output, exitCode, nil
}

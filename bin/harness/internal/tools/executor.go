// Package tools provides file system, command execution, web search, code search, and content fetch tools.
package tools

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ToolResult holds the outcome of a tool execution.
type ToolResult struct {
	Output   string
	ExitCode int
	Error    error
}

// ToolExecutor defines the interface for executing agent tools (DIP).
type ToolExecutor interface {
	// Execute runs a tool by name with the given arguments.
	Execute(name string, args map[string]any) ToolResult
}

// Registry manages available tool executors (OCP: register, not modify).
type Registry struct {
	tools map[string]ToolExecutor
}

// NewRegistry creates an empty tool registry.
func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]ToolExecutor),
	}
}

// Register adds a tool executor. Panics if name already exists.
func (r *Registry) Register(name string, executor ToolExecutor) {
	if _, exists := r.tools[name]; exists {
		panic(fmt.Sprintf("tool '%s' already registered", name))
	}
	r.tools[name] = executor
}

// Execute runs a registered tool. Returns error if tool not found.
func (r *Registry) Execute(name string, args map[string]any) ToolResult {
	executor, exists := r.tools[name]
	if !exists {
		return ToolResult{Error: fmt.Errorf("tool '%s' not found", name)}
	}
	return executor.Execute(name, args)
}

// Names returns the list of registered tool names.
func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.tools))
	for n := range r.tools {
		names = append(names, n)
	}
	return names
}

// --- Built-in tool executors ---

// ReadFileExecutor implements the read_file tool.
type ReadFileExecutor struct {
	Root string // Workspace root for security boundary
}

func (e *ReadFileExecutor) Execute(name string, args map[string]any) ToolResult {
	path, _ := args["path"].(string)
	if path == "" {
		return ToolResult{Error: fmt.Errorf("missing required argument: path")}
	}

	resolved, err := resolvePath(path, e.Root)
	if err != nil {
		return ToolResult{Error: err}
	}

	file, err := os.Open(resolved)
	if err != nil {
		return ToolResult{Error: err}
	}
	defer file.Close()

	const maxLimit = 5 * 1024 * 1024
	limitedReader := io.LimitReader(file, maxLimit)

	content, err := io.ReadAll(limitedReader)
	if err != nil {
		return ToolResult{Error: err}
	}

	return ToolResult{Output: string(content)}
}

// WriteFileExecutor implements the write_file tool.
type WriteFileExecutor struct {
	Root string
}

func (e *WriteFileExecutor) Execute(name string, args map[string]any) ToolResult {
	path, _ := args["path"].(string)
	content, _ := args["content"].(string)
	if path == "" {
		return ToolResult{Error: fmt.Errorf("missing required argument: path")}
	}

	resolved, err := resolvePath(path, e.Root)
	if err != nil {
		return ToolResult{Error: err}
	}

	dir := filepath.Dir(resolved)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return ToolResult{Error: fmt.Errorf("failed to create directory tree: %w", err)}
	}

	if err := os.WriteFile(resolved, []byte(content), 0644); err != nil {
		return ToolResult{Error: fmt.Errorf("failed to write file: %w", err)}
	}

	return ToolResult{Output: "File written successfully."}
}

// PatchFileExecutor implements the patch_file tool.
type PatchFileExecutor struct {
	Root string
}

func (e *PatchFileExecutor) Execute(name string, args map[string]any) ToolResult {
	path, _ := args["path"].(string)
	target, _ := args["target"].(string)
	replacement, _ := args["replacement"].(string)
	if path == "" || target == "" {
		return ToolResult{Error: fmt.Errorf("missing required arguments: path, target")}
	}

	resolved, err := resolvePath(path, e.Root)
	if err != nil {
		return ToolResult{Error: err}
	}

	bytes, err := os.ReadFile(resolved)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("failed to read file for patching: %w", err)}
	}

	content := string(bytes)
	count := strings.Count(content, target)
	if count == 0 {
		return ToolResult{Error: fmt.Errorf("patch failed: target text not found in '%s'", path)}
	}
	if count > 1 {
		return ToolResult{Error: fmt.Errorf("patch failed: target text matches %d times in '%s' (must be unique)", count, path)}
	}

	newContent := strings.Replace(content, target, replacement, 1)
	if err := os.WriteFile(resolved, []byte(newContent), 0644); err != nil {
		return ToolResult{Error: fmt.Errorf("failed to write patched file: %w", err)}
	}

	return ToolResult{Output: "Patch applied successfully."}
}

// ExecCommandExecutor implements the execute_command tool.
type ExecCommandExecutor struct {
	Root string
}

func (e *ExecCommandExecutor) Execute(name string, args map[string]any) ToolResult {
	command, _ := args["command"].(string)
	if command == "" {
		return ToolResult{Error: fmt.Errorf("missing required argument: command")}
	}

	var shell, shellFlag string
	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		shellFlag = "/c"
	} else {
		shell = "bash"
		shellFlag = "-c"
	}

	cmd := exec.Command(shell, shellFlag, command)
	cmd.Dir = e.Root

	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)

	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			return ToolResult{Output: output, ExitCode: -1, Error: err}
		}
	}

	return ToolResult{Output: output, ExitCode: exitCode}
}

// SearchExecutor implements the search_files tool.
type SearchExecutor struct {
	Root string
}

func (e *SearchExecutor) Execute(name string, args map[string]any) ToolResult {
	pattern, _ := args["pattern"].(string)
	query, _ := args["query"].(string)

	result, err := e.search(pattern, query)
	if err != nil {
		return ToolResult{Error: err}
	}
	return ToolResult{Output: result}
}

func (e *SearchExecutor) search(pattern string, query string) (string, error) {
	root := e.Root
	patternLower := strings.ToLower(pattern)
	queryLower := strings.ToLower(query)

	const maxMatches = 100
	results := make([]string, 0, maxMatches)
	resultsMu := sync.Mutex{}
	jobs := make(chan string, 1024)

	numWorkers := runtime.NumCPU()
	if numWorkers < 2 {
		numWorkers = 2
	}
	var wg sync.WaitGroup

	type matchResult struct {
		lines []string
	}
	resultsChan := make(chan matchResult, numWorkers*2)

	worker := func() {
		defer wg.Done()
		for path := range jobs {
			name := filepath.Base(path)
			nameLower := strings.ToLower(name)

			if pattern != "" {
				match, err := filepath.Match(patternLower, nameLower)
				if err != nil || !match {
					continue
				}
			}

			if query == "" {
				relPath, err := filepath.Rel(root, path)
				if err != nil {
					relPath = path
				}
				resultsChan <- matchResult{lines: []string{relPath}}
				continue
			}

			if isBinary(path) {
				continue
			}

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
				continue
			}

			contentBytes, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				continue
			}

			contentStr := string(contentBytes)
			if !strings.Contains(strings.ToLower(contentStr), queryLower) {
				continue
			}

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

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	dispatcherDone := make(chan error, 1)
	go func() {
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return nil
			}
			name := d.Name()
			if d.IsDir() {
				if strings.HasPrefix(name, ".") && name != "." && name != ".." {
					return filepath.SkipDir
				}
				if name == "node_modules" || name == "dist" || name == "build" || name == "vendor" {
					return filepath.SkipDir
				}
				return nil
			}
			if pattern != "" && query == "" {
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

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	collectorWg := sync.WaitGroup{}
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		for res := range resultsChan {
			resultsMu.Lock()
			if len(results) >= maxMatches {
				resultsMu.Unlock()
				continue
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

	dispatchErr := <-dispatcherDone
	collectorWg.Wait()

	if dispatchErr != nil {
		return "", dispatchErr
	}

	if len(results) == 0 {
		return "Nenhum arquivo correspondente foi encontrado.", nil
	}

	return strings.Join(results, "\n"), nil
}

// --- Helpers ---

func resolvePath(p string, root string) (string, error) {
	if root == "" {
		var err error
		root, err = findWorkspaceRoot()
		if err != nil {
			return "", err
		}
	}
	if filepath.IsAbs(p) {
		absClean := filepath.Clean(p)
		rootClean := filepath.Clean(root)
		if !strings.HasPrefix(absClean, rootClean) {
			return "", fmt.Errorf("security boundary violation: path '%s' is outside workspace root '%s'", p, root)
		}
		return absClean, nil
	}
	return filepath.Clean(filepath.Join(root, p)), nil
}

func findWorkspaceRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		specsPath := filepath.Join(dir, ".specs")
		if info, err := os.Stat(specsPath); err == nil && info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("workspace root (.specs directory) not found")
}

// --- MCP Tool Executor (no import cycle: uses Callback) ---

// MCPToolExecutor wraps an MCP client tool using a callback to avoid import cycles.
type MCPToolExecutor struct {
	// CallFn is called with the tool name and args, returns output text or error.
	CallFn      func(ctx context.Context, toolName string, args map[string]interface{}) (string, error)
	ToolName    string
	Description string
}

func (e *MCPToolExecutor) Execute(name string, args map[string]any) ToolResult {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mcpArgs := make(map[string]interface{})
	for k, v := range args {
		mcpArgs[k] = v
	}

	output, err := e.CallFn(ctx, e.ToolName, mcpArgs)
	if err != nil {
		return ToolResult{Error: fmt.Errorf("mcp:%s: %w", e.ToolName, err)}
	}

	return ToolResult{Output: output}
}

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


// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Create temp workspace with 2000 small Go files + 2000 non-Go files
	tmpDir, err := os.MkdirTemp("", "harness-smoke-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create .harness/config.json to make FindWorkspaceRoot work
	harnessDir := filepath.Join(tmpDir, ".harness")
	os.MkdirAll(harnessDir, 0755)
	os.WriteFile(filepath.Join(harnessDir, "config.json"), []byte(`{"active_provider":"mock","providers":{"mock":{}}}`), 0644)

	// Create 2000 Go files and 2000 text files
	fmt.Println("Creating 4000 test files...")
	for i := 0; i < 2000; i++ {
		content := fmt.Sprintf("package main\n\nfunc file_%d() {\n\tprintln(\"hello\")\n}\n", i)
		os.WriteFile(filepath.Join(tmpDir, fmt.Sprintf("file_%04d.go", i)), []byte(content), 0644)
	}
	for i := 0; i < 2000; i++ {
		content := fmt.Sprintf("This is text file number %d with some random content.\n", i)
		os.WriteFile(filepath.Join(tmpDir, fmt.Sprintf("readme_%04d.txt", i)), []byte(content), 0644)
	}

	// Temporarily change working directory so FindWorkspaceRoot uses tmpDir
	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	// We need to import the search function, but it's in package main of another module.
	// Instead, let's just time the `rg` command vs fallback? But the harness uses pure Go.
	// Better: we can run the test binary from this directory.
	fmt.Println("Smoke test: workspace with 4000 files ready.")
	fmt.Println("Run: cd", tmpDir, "&& go test -bench=BenchmarkLargeWorkspace -benchtime=1x -count=1 ./...")

	// Actually let's just do a manual time measurement using the built binary
	// We'll use the existing benchmark test via command line.
}

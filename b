ok`); err == nil { t.Error("expected error for unclosed tag") }
}

func TestIsDestructiveCommand(t *testing.T) {
	if !isDestructiveCommand("rm -rf /") { t.Error("'rm -rf /' should be destructive") }
	if !isDestructiveCommand("DROP TABLE users") { t.Error("'DROP TABLE' should be destructive") }
	if isDestructiveCommand("ls -la") { t.Error("'ls -la' should NOT be destructive") }
	if isDestructiveCommand("echo hello") { t.Error("'echo hello' should NOT be destructive") }
}

func TestCallLLMRetryOnTimeout(t *testing.T) {
	var callCount int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&callCount, 1)
		time.Sleep(50 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"text":"test"}`))
	}))
	defer server.Close()
	originalTimeout := llmRequestTimeout
	llmRequestTimeout = 15 * time.Millisecond
	defer func() { llmRequestTimeout = originalTimeout }()
	cfg := &AppConfig{
		ActiveProvider: "mock-llm",
		Providers: map[string]ProviderConfig{
			"mock-llm": {URL: server.URL, Headers: map[string]string{"Content-Type": "application/json"},
				BodyTemplate: `{"prompt": "{{prompt}}"}`,
				ResponsePath: "text",
			},
		},
	}
	_, err := CallLLM(cfg, "test", "")
	if err == nil { t.Errorf("expected timeout error, got success") }
	if c := atomic.LoadInt64(&callCount); c != 3 { t.Errorf("expected 3 attempts, got %d", c) }
}

func TestCallLLMStream_Basic(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		chunks := []string{
			`data: {"choices":[{"delta":{"content":"Hello "}}]}` + "\n\n",
			`data: {"choices":[{"delta":{"content":"world"}}]}` + "\n\n",
			"data: [DONE]\n\n",
		}
		for _, chunk := range chunks {
			fmt.Fprint(w, chunk)
			w.(http.Flusher).Flush()
		}
	}))
	defer server.Close()
	cfg := &AppConfig{
		ActiveProvider: "mock-stream",
		Providers: map[string]ProviderConfig{
			"mock-stream": {
				URL:                server.URL,
				Headers:            map[string]string{"Content-Type": "application/json"},
				BodyTemplate:       `{"prompt": "{{prompt}}"}`,
				Streaming:          true,
				StreamResponsePath: "choices.0.delta.content",
			},
		},
	}
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	result, err := CallLLMStream(cfg, "test", "")
	if err != nil { t.Fatalf("CallLLMStream failed: %v", err) }
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	partialOutput := buf.String()
	if result != "Hello world" { t.Errorf("expected 'Hello world', got '%s'", result) }
	if !strings.Contains(partialOutput, "Hello ") || !strings.Contains(partialOutput, "world") { t.Errorf("expected partial tokens in stdout, got '%s'", partialOutput) }
}

// Ripgrep integration tests
func TestIsRipgrepAvailable(t *testing.T) {
	available := isRipgrepAvailable()
	_ = available // just ensure function runs without crash
}

func TestSearchWithRipgrep_Fallback(t *testing.T) {
	res, err := SearchFiles("tools.go", "")
	if err != nil { t.Fatalf("SearchFiles failed: %v", err) }
	if !strings.Contains(res, "bin/harness/tools.go") {
		t.Errorf("expected to find bin/harness/tools.go, got:\n%s", res)
	}
}

func TestSearchWithRipgrep_Content(t *testing.T) {
	res, err := SearchFiles("main.go", "runInteractiveLoop")
	if err != nil { t.Fatalf("SearchFiles failed: %v", err) }
	if !strings.Contains(res, "bin/harness/main.go:") {
		t.Errorf("expected to find main.go, got:\n%s", res)
	}
}

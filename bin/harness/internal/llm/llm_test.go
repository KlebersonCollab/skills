package llm

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// --- Mocks ---

type mockProvider struct {
	name      string
	failCount int64 // fail for this many calls, then succeed
	callCount int64
}

func (m *mockProvider) Name() string { return m.name }

func (m *mockProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	atomic.AddInt64(&m.callCount, 1)
	fail := atomic.LoadInt64(&m.failCount)
	if fail > 0 {
		atomic.AddInt64(&m.failCount, -1)
		return "", fmt.Errorf("mock %s: simulated failure", m.name)
	}
	return fmt.Sprintf("[%s response to: %s]", m.name, prompt[:min(len(prompt), 30)]), nil
}

func (m *mockProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	atomic.AddInt64(&m.callCount, 1)
	ch := make(chan string, 2)
	ch <- fmt.Sprintf("[%s stream]", m.name)
	ch <- "\n"
	close(ch)
	return ch, nil
}

// --- Tests ---

func TestProviderRegistry_RegisterAndGet(t *testing.T) {
	r := NewProviderRegistry()
	mock := &mockProvider{name: "mock-1"}

	if err := r.Register(mock); err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	got, err := r.Get("mock-1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got.Name() != "mock-1" {
		t.Errorf("expected name 'mock-1', got %q", got.Name())
	}
}

func TestProviderRegistry_RegisterDuplicate(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "dup"})
	err := r.Register(&mockProvider{name: "dup"})
	if err == nil {
		t.Error("expected error for duplicate registration")
	}
}

func TestProviderRegistry_GetNotFound(t *testing.T) {
	r := NewProviderRegistry()
	_, err := r.Get("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent provider")
	}
}

func TestProviderRegistry_Names(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "a"})
	r.Register(&mockProvider{name: "b"})
	r.Register(&mockProvider{name: "c"})

	names := r.Names()
	if len(names) != 3 {
		t.Errorf("expected 3 names, got %d", len(names))
	}
}

func TestProviderRegistry_FallbackOrder(t *testing.T) {
	r := NewProviderRegistry()
	a := &mockProvider{name: "a"}
	b := &mockProvider{name: "b"}
	c := &mockProvider{name: "c"}

	r.Register(a)
	r.Register(b)
	r.Register(c)

	r.SetFallbackOrder([]string{"a", "c"})
	order := r.FallbackOrder()

	if len(order) != 2 || order[0] != "a" || order[1] != "c" {
		t.Errorf("expected fallback order [a c], got %v", order)
	}
}

func TestProviderRegistry_FallbackOrder_IgnoredUnknown(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "real"})

	r.SetFallbackOrder([]string{"real", "fake"})
	order := r.FallbackOrder()

	if len(order) != 1 || order[0] != "real" {
		t.Errorf("expected only 'real', got %v", order)
	}
}

func TestCompleteWithFallback_FirstSucceeds(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "primary"})
	r.Register(&mockProvider{name: "secondary"})

	r.SetFallbackOrder([]string{"primary", "secondary"})

	result, err := r.CompleteWithFallback(context.Background(), "hello", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(result, "primary") {
		t.Errorf("expected primary response, got %q", result)
	}
}

func TestCompleteWithFallback_FirstFailsSecondSucceeds(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "primary", failCount: 1})
	r.Register(&mockProvider{name: "secondary"})

	r.SetFallbackOrder([]string{"primary", "secondary"})

	result, err := r.CompleteWithFallback(context.Background(), "hello", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(result, "secondary") {
		t.Errorf("expected secondary response after fallback, got %q", result)
	}
}

func TestCompleteWithFallback_AllFail(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "p1", failCount: 5})
	r.Register(&mockProvider{name: "p2", failCount: 5})

	r.SetFallbackOrder([]string{"p1", "p2"})

	_, err := r.CompleteWithFallback(context.Background(), "hello", nil, nil)
	if err == nil {
		t.Error("expected error when all providers fail")
	}
}

func TestCompleteWithFallback_NoProviders(t *testing.T) {
	r := NewProviderRegistry()
	_, err := r.CompleteWithFallback(context.Background(), "hello", nil, nil)
	if err == nil {
		t.Error("expected error with no providers")
	}
}

func TestStreamWithFallback(t *testing.T) {
	r := NewProviderRegistry()
	r.Register(&mockProvider{name: "streamer"})
	r.SetFallbackOrder([]string{"streamer"})

	ch, err := r.StreamWithFallback(context.Background(), "hello", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result string
	for token := range ch {
		result += token
	}
	if !strings.Contains(result, "streamer") {
		t.Errorf("expected streamer response, got %q", result)
	}
}

func TestExtractToken(t *testing.T) {
	token := ExtractToken(`{"text": "hello"}`, "text")
	if token != "hello" {
		t.Errorf("expected 'hello', got %q", token)
	}
}

func TestExtractToken_EmptyPath(t *testing.T) {
	token := ExtractToken(`{"a": 1}`, "nonexistent")
	if token != "" {
		t.Errorf("expected empty for nonexistent path, got %q", token)
	}
}

func TestExtractToken_InvalidJSON(t *testing.T) {
	token := ExtractToken(`not json`, "path")
	if token != "" {
		t.Errorf("expected empty for invalid JSON, got %q", token)
	}
}

func TestRetryCall_Success(t *testing.T) {
	callCount := 0
	result, err := RetryCall(context.Background(), RetryConfig{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond},
		func(ctx context.Context) (string, error) {
			callCount++
			return "success", nil
		})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "success" {
		t.Errorf("expected 'success', got %q", result)
	}
	if callCount != 1 {
		t.Errorf("expected 1 call, got %d", callCount)
	}
}

func TestRetryCall_AllFail(t *testing.T) {
	callCount := 0
	_, err := RetryCall(context.Background(), RetryConfig{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond},
		func(ctx context.Context) (string, error) {
			callCount++
			return "", fmt.Errorf("fail")
		})
	if err == nil {
		t.Error("expected error")
	}
	if callCount != 3 {
		t.Errorf("expected 3 calls, got %d", callCount)
	}
}

func TestRetryCall_ContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := RetryCall(ctx, RetryConfig{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond},
		func(ctx context.Context) (string, error) {
			return "", fmt.Errorf("should not be called")
		})
	if err == nil {
		t.Error("expected error from canceled context")
	}
}

func TestResolveTemplate(t *testing.T) {
	tmpl := `{"prompt": "{{prompt}}", "history": "{{history}}"}`
	result := ResolveTemplate(tmpl, "hi", "previous")
	if !strings.Contains(result, "hi") || !strings.Contains(result, "previous") {
		t.Errorf("expected template to contain prompt and history, got %q", result)
	}
}

func TestNewGeminiProvider(t *testing.T) {
	p := NewGeminiProvider("test-key", "")
	if p.Name() != "gemini" {
		t.Errorf("expected name 'gemini', got %q", p.Name())
	}
}

func TestNewDeepSeekProvider(t *testing.T) {
	p := NewDeepSeekProvider("test-key", "", nil)
	if p.Name() != "deepseek" {
		t.Errorf("expected name 'deepseek', got %q", p.Name())
	}
}

func TestNewOllamaProvider(t *testing.T) {
	p := NewOllamaProvider("", "")
	if p.Name() != "ollama" {
		t.Errorf("expected name 'ollama', got %q", p.Name())
	}
}

func TestWebSearchRegistry(t *testing.T) {
	r := NewWebSearchRegistry("perp-key", "", "gemini-key")
	if r == nil {
		t.Fatal("expected non-nil registry")
	}
}

func TestNewWebSearchRegistry_Empty(t *testing.T) {
	r := NewWebSearchRegistry("", "", "")
	result, err := r.Search(context.Background(), "test", WebSearchOpts{})
	if err != nil {
		t.Fatalf("fallback search should not error: %v", err)
	}
	if result.Provider != "none" {
		t.Errorf("expected 'none' provider for empty keys, got %q", result.Provider)
	}
}

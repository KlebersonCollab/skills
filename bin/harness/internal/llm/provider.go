// Package llm provides abstractions for LLM providers and the registry.
package llm

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

// LLMProvider defines the interface that all LLM providers must implement.
// This is the core abstraction for Dependency Inversion (DIP).
type LLMProvider interface {
	// Name returns the provider identifier (e.g., "gemini", "deepseek").
	Name() string

	// Complete sends a prompt and returns the full response text.
	Complete(ctx context.Context, prompt string, opts map[string]any) (string, error)

	// Stream sends a prompt and returns a channel of text chunks.
	// The channel is closed when streaming completes or an error occurs.
	Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error)
}

// ProviderRegistry manages registered LLM providers with fallback support.
// OCP: new providers are added via Register, not by modifying switch statements.
// Fallback: providers are tried in order if the primary fails.
type ProviderRegistry struct {
	mu            sync.RWMutex
	providers     map[string]LLMProvider
	fallbackOrder []string  // Ordem de tentativa para fallback
}

// NewProviderRegistry creates an empty provider registry.
func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{
		providers: make(map[string]LLMProvider),
	}
}

// Register adds a provider to the registry. Returns error if name already exists.
func (r *ProviderRegistry) Register(p LLMProvider) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := p.Name()
	if _, exists := r.providers[name]; exists {
		return fmt.Errorf("provider '%s' already registered", name)
	}
	r.providers[name] = p
	// Auto-adiciona à fallback order se ainda nao estiver
	found := false
	for _, n := range r.fallbackOrder {
		if n == name {
			found = true
			break
		}
	}
	if !found {
		r.fallbackOrder = append(r.fallbackOrder, name)
	}
	return nil
}

// SetFallbackOrder defines the order in which providers are tried for fallback.
// Only registered providers are kept; unknown names are ignored.
func (r *ProviderRegistry) SetFallbackOrder(names []string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order := make([]string, 0, len(names))
	for _, name := range names {
		if _, exists := r.providers[name]; exists {
			order = append(order, name)
		}
	}
	if len(order) > 0 {
		r.fallbackOrder = order
	}
}

// Get retrieves a provider by name. Returns error if not found.
func (r *ProviderRegistry) Get(name string) (LLMProvider, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.providers[name]
	if !exists {
		return nil, fmt.Errorf("provider '%s' not found in registry", name)
	}
	return p, nil
}

// Names returns the list of registered provider names.
func (r *ProviderRegistry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.providers))
	for n := range r.providers {
		names = append(names, n)
	}
	return names
}

// FallbackOrder returns the current fallback order.
func (r *ProviderRegistry) FallbackOrder() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order := make([]string, len(r.fallbackOrder))
	copy(order, r.fallbackOrder)
	return order
}

// CompleteWithFallback tries providers in fallback order until one succeeds.
// Returns the first successful response. If all fail, returns the last error.
func (r *ProviderRegistry) CompleteWithFallback(ctx context.Context, prompt string, opts map[string]any, logFn func(format string, args ...any)) (string, error) {
	r.mu.RLock()
	order := make([]string, len(r.fallbackOrder))
	copy(order, r.fallbackOrder)
	r.mu.RUnlock()

	if len(order) == 0 {
		return "", fmt.Errorf("no providers registered")
	}

	var lastErr error
	for i, name := range order {
		p, err := r.Get(name)
		if err != nil {
			lastErr = err
			continue
		}

		if i > 0 && logFn != nil {
			logFn("\033[38;5;208m⚠️  Fallback para %s (tentativa %d/%d)...\033[0m\n", strings.ToUpper(name), i+1, len(order))
		}

		result, err := p.Complete(ctx, prompt, opts)
		if err == nil {
			if i > 0 && logFn != nil {
				logFn("\033[38;5;32m✅ %s respondeu com sucesso\033[0m\n", strings.ToUpper(name))
			}
			return result, nil
		}
		lastErr = err

		if logFn != nil {
			logFn("\033[38;5;196m❌ %s falhou: %v\033[0m\n", strings.ToUpper(name), err)
		}
	}

	return "", fmt.Errorf("all %d providers failed: %w", len(order), lastErr)
}

// StreamWithFallback tries providers in fallback order for streaming.
// Returns the first successful stream channel.
func (r *ProviderRegistry) StreamWithFallback(ctx context.Context, prompt string, opts map[string]any, logFn func(format string, args ...any)) (<-chan string, error) {
	r.mu.RLock()
	order := make([]string, len(r.fallbackOrder))
	copy(order, r.fallbackOrder)
	r.mu.RUnlock()

	if len(order) == 0 {
		return nil, fmt.Errorf("no providers registered")
	}

	var lastErr error
	for i, name := range order {
		p, err := r.Get(name)
		if err != nil {
			lastErr = err
			continue
		}

		if i > 0 && logFn != nil {
			logFn("\033[38;5;208m⚠️  Fallback streaming para %s...\033[0m\n", strings.ToUpper(name))
		}

		ch, err := p.Stream(ctx, prompt, opts)
		if err == nil {
			return ch, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("all %d providers failed for streaming: %w", len(order), lastErr)
}

// ActiveProvider holds the currently active provider name and configuration.
type ActiveProvider struct {
	Name     string
	Provider LLMProvider
}

// CallOpts defines options for an LLM call.
type CallOpts struct {
	Streaming     bool
	History       string
	MaxTokens     int
	ThinkingLevel string // "off", "minimal", "low", "medium", "high", "xhigh"
}

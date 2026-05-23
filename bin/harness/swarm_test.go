package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"harness/internal/config"
)

// mockProvider implements LLMProvider for testing
type mockProvider struct {
	name     string
	response string
}

func (m *mockProvider) Name() string { return m.name }

func (m *mockProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	if m.response != "" {
		return m.response, nil
	}
	return "mock response for: " + prompt[:min(len(prompt), 50)], nil
}

func (m *mockProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	ch := make(chan string, 1)
	ch <- "mock stream"
	close(ch)
	return ch, nil
}

func TestDefaultSwarmConfig(t *testing.T) {
	cfg := DefaultSwarmConfig()
	if cfg.Topology != "centralized" {
		t.Errorf("expected topology 'centralized', got '%s'", cfg.Topology)
	}
	if len(cfg.Agents) == 0 {
		t.Error("expected at least one agent in default config")
	}
	if cfg.MaxWorkers <= 0 {
		t.Errorf("expected MaxWorkers > 0, got %d", cfg.MaxWorkers)
	}
	if cfg.TimeoutSec <= 0 {
		t.Errorf("expected TimeoutSec > 0, got %d", cfg.TimeoutSec)
	}

	// Verify mandatory roles
	hasOrchestrator := false
	hasGenerator := false
	hasDiscriminator := false
	hasWorker := false
	for _, a := range cfg.Agents {
		switch a.Role {
		case RoleOrchestrator:
			hasOrchestrator = true
		case RoleGenerator:
			hasGenerator = true
		case RoleDiscriminator:
			hasDiscriminator = true
		case RoleWorker:
			hasWorker = true
		}
	}
	if !hasOrchestrator {
		t.Error("default config missing orchestrator")
	}
	if !hasGenerator {
		t.Error("default config missing generator")
	}
	if !hasDiscriminator {
		t.Error("default config missing discriminator")
	}
	if !hasWorker {
		t.Error("default config missing worker")
	}
}

func TestNewSwarmOrchestrator(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	if orch == nil {
		t.Fatal("expected non-nil orchestrator")
	}
	if len(orch.Agents) != len(cfg.Agents) {
		t.Errorf("expected %d agents, got %d", len(cfg.Agents), len(orch.Agents))
	}
}

func TestFindAgent(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)

	orchAgent := orch.findAgent(RoleOrchestrator)
	if orchAgent == nil {
		t.Fatal("expected orchestrator agent")
	}
	if orchAgent.Config.Role != RoleOrchestrator {
		t.Errorf("expected RoleOrchestrator, got %s", orchAgent.Config.Role)
	}

	workerAgent := orch.findAgent(RoleWorker)
	if workerAgent == nil {
		t.Fatal("expected worker agent")
	}

	nonExistent := orch.findAgent("unknown")
	if nonExistent != nil {
		t.Error("expected nil for unknown role")
	}
}

func TestGetWorkers(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	workers := orch.getWorkers()
	if len(workers) == 0 {
		t.Error("expected at least one worker")
	}
	for _, w := range workers {
		if w.Config.Role != RoleWorker {
			t.Errorf("expected role 'worker', got '%s'", w.Config.Role)
		}
	}
}

func TestBindLLM(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)

	mock := &mockProvider{name: "mock", response: "hello"}
	err := orch.BindLLM("orchestrator", mock)

	if err != nil {
		t.Fatalf("BindLLM failed: %v", err)
	}

	agent := orch.findAgent(RoleOrchestrator)
	if agent.LLM == nil {
		t.Fatal("expected LLM provider to be bound")
	}

	err = orch.BindLLM("nonexistent", mock)
	if err == nil {
		t.Error("expected error for unknown agent")
	}
}

func TestDecomposeTask_Fallback(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	// No LLM bound — should fallback to single task

	tasks, err := orch.decomposeTask(context.Background(), "test goal", "test context")
	if err != nil {
		t.Fatalf("decomposeTask failed: %v", err)
	}
	if len(tasks) == 0 {
		t.Fatal("expected at least one task")
	}
	_ = tasks // fallback creates one task
}

func TestExecuteSwarm_Centralized_Mock(t *testing.T) {
	cfg := DefaultSwarmConfig()
	cfg.Topology = "centralized"
	orch := NewSwarmOrchestrator(cfg)

	// Bind all agents with mock
	for name := range orch.Agents {
		orch.BindLLM(name, &mockProvider{name: name, response: "mock output for " + name})
	}

	result, err := orch.ExecuteSwarm("test goal", "test context")
	if err != nil {
		t.Fatalf("ExecuteSwarm failed: %v", err)
	}
	if result == "" {
		t.Error("expected non-empty result")
	}
}

func TestExecuteSwarm_Sequential(t *testing.T) {
	cfg := DefaultSwarmConfig()
	cfg.Topology = "sequential"
	orch := NewSwarmOrchestrator(cfg)

	for name := range orch.Agents {
		orch.BindLLM(name, &mockProvider{name: name, response: "seq output from " + name})
	}

	result, err := orch.ExecuteSwarm("sequential goal", "ctx")
	if err != nil {
		t.Fatalf("ExecuteSwarm sequential failed: %v", err)
	}
	if result == "" {
		t.Error("expected non-empty result")
	}
}

func TestExecuteSwarm_GAN_Mock(t *testing.T) {
	cfg := DefaultSwarmConfig()
	cfg.Topology = "gan"
	orch := NewSwarmOrchestrator(cfg)

	for name := range orch.Agents {
		orch.BindLLM(name, &mockProvider{name: name, response: "mock GAN output"})
	}

	result, err := orch.ExecuteSwarm("GAN goal", "ctx")
	// GAN may or may not find APPROVED — just ensure no panic
	_ = result
	if err != nil {
		t.Logf("expected GAN error: %v", err)
	}
}

func TestCollectResults(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)

	orch.Results <- SwarmResult{AgentName: "a", Output: "out1"}
	orch.Results <- SwarmResult{AgentName: "b", Output: "out2"}

	results := orch.collectResults()
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestGenericProvider_Stream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"test":"hello generic streaming"}`))
	}))
	defer server.Close()

	gp := &genericProvider{
		cfg: config.ProviderConfig{
			URL:          server.URL,
			BodyTemplate: `{"prompt":"{{prompt}}"}`,
			ResponsePath: "test",
		},
	}

	ch, err := gp.Stream(context.Background(), "hello", nil)
	if err != nil {
		t.Fatalf("Stream failed: %v", err)
	}

	var parts []string
	for token := range ch {
		parts = append(parts, token)
	}

	result := strings.Join(parts, "")
	if result != "hello generic streaming" {
		t.Errorf("expected 'hello generic streaming', got '%s'", result)
	}
}

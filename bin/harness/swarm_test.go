
package main

import (
	"strings"
	"testing"
)

func TestDefaultSwarmConfig(t *testing.T) {
	cfg := DefaultSwarmConfig()
	if cfg.MaxWorkers != 3 {
		t.Errorf("expected MaxWorkers=3, got %d", cfg.MaxWorkers)
	}
	if len(cfg.Agents) != 4 {
		t.Errorf("expected 4 agents, got %d", len(cfg.Agents))
	}
	hasOrch := false
	hasGen := false
	hasDisc := false
	for _, a := range cfg.Agents {
		switch a.Role {
		case RoleOrchestrator:
			hasOrch = true
		case RoleGenerator:
			hasGen = true
		case RoleDiscriminator:
			hasDisc = true
		}
	}
	if !hasOrch {
		t.Error("missing orchestrator")
	}
	if !hasGen {
		t.Error("missing generator")
	}
	if !hasDisc {
		t.Error("missing discriminator")
	}
}

func TestNewSwarmOrchestrator(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	if orch == nil {
		t.Fatal("expected non-nil orchestrator")
	}
	if len(orch.Agents) != 4 {
		t.Errorf("expected 4 agents, got %d", len(orch.Agents))
	}
}

func TestFindAgent(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	
	gen := orch.findAgent(RoleGenerator)
	if gen == nil {
		t.Fatal("expected to find generator")
	}
	if gen.Config.Role != RoleGenerator {
		t.Errorf("expected generator role, got %s", gen.Config.Role)
	}
	
	none := orch.findAgent("nobody")
	if none != nil {
		t.Error("expected nil for nonexistent agent")
	}
}

func TestGetWorkers(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	
	workers := orch.getWorkers()
	if len(workers) != 1 {
		t.Errorf("expected 1 worker, got %d", len(workers))
	}
}

func TestBindLLM(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	
	err := orch.BindLLM("generator", func(prompt string) (string, error) {
		return "mock output", nil
	})
	if err != nil {
		t.Fatalf("BindLLM failed: %v", err)
	}
	
	gen := orch.findAgent(RoleGenerator)
	out, err := gen.LLM("test")
	if err != nil {
		t.Fatalf("LLM call failed: %v", err)
	}
	if out != "mock output" {
		t.Errorf("expected 'mock output', got '%s'", out)
	}
}

func TestDecomposeTask_Fallback(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	orch.BindLLM("orchestrator", func(prompt string) (string, error) {
		return "invalid json", nil
	})
	
	tasks, err := orch.DecomposeTask("test goal", "test context")
	if err != nil {
		t.Fatalf("DecomposeTask failed: %v", err)
	}
	if len(tasks) != 1 {
		t.Errorf("expected 1 fallback task, got %d", len(tasks))
	}
	if tasks[0].Goal != "test goal" {
		t.Errorf("expected goal 'test goal', got '%s'", tasks[0].Goal)
	}
}

func TestExecuteSwarm_Centralized_Mock(t *testing.T) {
	cfg := DefaultSwarmConfig()
	cfg.Topology = "centralized"
	orch := NewSwarmOrchestrator(cfg)
	
	// Bind all agents with mock LLMs
	for name := range orch.Agents {
		orch.BindLLM(name, func(prompt string) (string, error) {
			if strings.Contains(prompt, "Decompose") {
				return `[{"id":"t1","goal":"subtask 1","context":"ctx"}]`, nil
			}
			return "mock result from " + name, nil
		})
	}
	
	result, err := orch.ExecuteSwarm("main goal", "main context")
	if err != nil {
		t.Fatalf("ExecuteSwarm failed: %v", err)
	}
	if !strings.Contains(result, "mock result") {
		t.Errorf("expected result to contain 'mock result', got: %s", result)
	}
}

func TestExecuteSwarm_Sequential(t *testing.T) {
	cfg := SwarmConfig{
		Topology:   "sequential",
		MaxWorkers: 2,
		Agents: []SwarmAgentConfig{
			{Name: "orch", Role: RoleOrchestrator, SystemPrompt: "orch"},
			{Name: "a1", Role: RoleWorker, SystemPrompt: "worker 1"},
			{Name: "a2", Role: RoleWorker, SystemPrompt: "worker 2"},
		},
	}
	orch := NewSwarmOrchestrator(cfg)
	for name := range orch.Agents {
		orch.BindLLM(name, func(prompt string) (string, error) {
			return "output from " + name, nil
		})
	}
	
	result, err := orch.ExecuteSwarm("test", "ctx")
	if err != nil {
		t.Fatalf("sequential swarm failed: %v", err)
	}
	if !strings.Contains(result, "output from") {
		t.Errorf("expected agent output in result, got: %s", result)
	}
}

func TestExecuteSwarm_GAN_Mock(t *testing.T) {
	cfg := SwarmConfig{
		Topology: "gan",
		Agents: []SwarmAgentConfig{
			{Name: "generator", Role: RoleGenerator, SystemPrompt: "gen"},
			{Name: "discriminator", Role: RoleDiscriminator, SystemPrompt: "disc"},
		},
	}
	orch := NewSwarmOrchestrator(cfg)
	
	genCalls := 0
	orch.BindLLM("generator", func(prompt string) (string, error) {
		genCalls++
		return "proposal v" + string(rune('0'+genCalls)), nil
	})
	orch.BindLLM("discriminator", func(prompt string) (string, error) {
		if genCalls >= 2 {
			return "APPROVED", nil
		}
		return "needs improvement", nil
	})
	
	result, err := orch.ExecuteSwarm("test goal", "test context")
	if err != nil {
		t.Fatalf("GAN swarm failed: %v", err)
	}
	if !strings.Contains(result, "proposal") {
		t.Errorf("expected proposal in result, got: %s", result)
	}
}

func TestCollectResults(t *testing.T) {
	cfg := DefaultSwarmConfig()
	orch := NewSwarmOrchestrator(cfg)
	
	orch.Results <- SwarmResult{AgentName: "a1", Output: "out1"}
	orch.Results <- SwarmResult{AgentName: "a2", Output: "out2"}
	
	results := orch.collectResults()
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

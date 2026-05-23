package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"harness/internal/llm"
)

type SwarmRole string

const (
	RoleOrchestrator  SwarmRole = "orchestrator"
	RoleWorker        SwarmRole = "worker"
	RoleGenerator     SwarmRole = "generator"
	RoleDiscriminator SwarmRole = "discriminator"
)

type SwarmAgentConfig struct {
	Name         string    `json:"name"`
	Role         SwarmRole `json:"role"`
	Provider     string    `json:"provider"`
	SystemPrompt string    `json:"system_prompt"`
	Tools        []string  `json:"tools"`
}

type SwarmConfig struct {
	Agents     []SwarmAgentConfig `json:"agents"`
	Topology   string             `json:"topology"`
	MaxWorkers int                `json:"max_workers"`
	TimeoutSec int                `json:"timeout_sec"`
}

type SwarmResult struct {
	AgentName string     `json:"agent_name"`
	Role      SwarmRole  `json:"role"`
	Output    string     `json:"output"`
	Tokens    int        `json:"tokens"`
	Error     string     `json:"error,omitempty"`
	Duration  time.Duration `json:"duration"`
}

type SubTask struct {
	ID       string `json:"id"`
	Goal     string `json:"goal"`
	Context  string `json:"context"`
	Expected string `json:"expected,omitempty"`
}

type SwarmOrchestrator struct {
	Config  SwarmConfig
	Agents  map[string]*SwarmAgent
	Results chan SwarmResult
	mu      sync.Mutex
}

type SwarmAgent struct {
	Config SwarmAgentConfig
	LLM    llm.LLMProvider
}
func NewSwarmOrchestrator(cfg SwarmConfig) *SwarmOrchestrator {
	agents := make(map[string]*SwarmAgent)
	for _, ac := range cfg.Agents {
		agents[ac.Name] = &SwarmAgent{Config: ac}
	}
	return &SwarmOrchestrator{
		Config:  cfg,
		Agents:  agents,
		Results: make(chan SwarmResult, 100),
	}
}

func (o *SwarmOrchestrator) BindLLM(name string, provider llm.LLMProvider) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	a, ok := o.Agents[name]
	if !ok {
		return fmt.Errorf("agent %s not found", name)
	}
	a.LLM = provider
	return nil
}

func (o *SwarmOrchestrator) ExecuteSwarm(goal, ctxStr string) (string, error) {
	ctx := context.Background()
	if o.Config.TimeoutSec > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(o.Config.TimeoutSec)*time.Second)
		defer cancel()
	}

	switch o.Config.Topology {
	case "gan":
		return o.executeGAN(ctx, goal, ctxStr)
	case "sequential":
		return o.executeSequential(ctx, goal, ctxStr)
	default:
		return o.executeCentralized(ctx, goal, ctxStr)
	}
}

func (o *SwarmOrchestrator) executeCentralized(ctx context.Context, goal, ctxStr string) (string, error) {
	tasks, err := o.decomposeTask(ctx, goal, ctxStr)
	if err != nil {
		return "", err
	}
	var wg sync.WaitGroup
	// Semaphore to limit parallel workers based on MaxWorkers config
	sem := make(chan struct{}, 100)
	if o.Config.MaxWorkers > 0 {
		sem = make(chan struct{}, o.Config.MaxWorkers)
	}
	for _, t := range tasks {
		wg.Add(1)
		go func(t SubTask) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				o.Results <- SwarmResult{AgentName: "orchestrator", Role: RoleWorker, Error: ctx.Err().Error()}
				return
			case sem <- struct{}{}:
				defer func() { <-sem }()
			}
			o.executeWorker(ctx, t)
		}(t)
	}
	wg.Wait()
	return o.mergeResults(ctx, goal, ctxStr, o.collectResults())
}

func (o *SwarmOrchestrator) executeSequential(ctx context.Context, goal, ctxStr string) (string, error) {
	for _, a := range o.Agents {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}
		if a.Config.Role == RoleOrchestrator || a.LLM == nil {
			continue
		}
		p := fmt.Sprintf("Goal: %s\nContext: %s\nAgent %s: %s", goal, ctxStr, a.Config.Name, a.Config.SystemPrompt)
		out, err := a.LLM.Complete(ctx, p, nil)
		if err != nil {
			return "", err
		}
		o.Results <- SwarmResult{AgentName: a.Config.Name, Role: a.Config.Role, Output: out}
		ctxStr = out
	}
	return o.mergeResults(ctx, goal, ctxStr, o.collectResults())
}

func (o *SwarmOrchestrator) executeGAN(ctx context.Context, goal, ctxStr string) (string, error) {
	gen := o.findAgent(RoleGenerator)
	disc := o.findAgent(RoleDiscriminator)
	if gen == nil || gen.LLM == nil || disc == nil || disc.LLM == nil {
		return "", fmt.Errorf("GAN requires generator and discriminator")
	}
	proposal := ""
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}
		p := fmt.Sprintf("Generate for: %s\nContext: %s\nIteration %d/3", goal, ctxStr, i+1)
		out, err := gen.LLM.Complete(ctx, p, nil)
		if err != nil {
			return "", err
		}
		proposal = out
		o.Results <- SwarmResult{AgentName: gen.Config.Name, Role: RoleGenerator, Output: out}
		dp := fmt.Sprintf("Critique:\nGoal: %s\nSolution:\n%s", goal, proposal)
		dout, err := disc.LLM.Complete(ctx, dp, nil)
		if err != nil {
			return "", err
		}
		o.Results <- SwarmResult{AgentName: disc.Config.Name, Role: RoleDiscriminator, Output: dout}
		if strings.Contains(strings.ToUpper(dout), "APPROVED") {
			return proposal, nil
		}
		proposal = fmt.Sprintf("Proposal:\n%s\nCritique:\n%s", proposal, dout)
	}
	return proposal, fmt.Errorf("max GAN iterations")
}

func (o *SwarmOrchestrator) decomposeTask(ctx context.Context, goal, ctxStr string) ([]SubTask, error) {
	orch := o.findAgent(RoleOrchestrator)
	if orch == nil || orch.LLM == nil {
		return []SubTask{{ID: "task-1", Goal: goal, Context: ctxStr}}, nil
	}
	p := fmt.Sprintf("Decompose into JSON subtasks: Goal: %s, Context: %s", goal, ctxStr)
	resp, err := orch.LLM.Complete(ctx, p, nil)
	if err != nil {
		return []SubTask{{ID: "task-1", Goal: goal, Context: ctxStr}}, nil
	}
	var tasks []SubTask
	clean := strings.TrimSpace(resp)
	if idx := strings.Index(clean, "["); idx >= 0 {
		if end := strings.LastIndex(clean, "]"); end > idx {
			clean = clean[idx : end+1]
		}
	}
	if err := json.Unmarshal([]byte(clean), &tasks); err != nil {
		return []SubTask{{ID: "task-1", Goal: goal, Context: ctxStr}}, nil
	}
	return tasks, nil
}

func (o *SwarmOrchestrator) executeWorker(ctx context.Context, task SubTask) {
	workers := o.getWorkers()
	if len(workers) == 0 {
		o.Results <- SwarmResult{AgentName: "none", Role: RoleWorker, Error: "no workers"}
		return
	}
	w := workers[0]
	if w.LLM == nil {
		o.Results <- SwarmResult{AgentName: w.Config.Name, Role: RoleWorker, Error: "no LLM"}
		return
	}

	systemPrompt := w.Config.SystemPrompt
	if len(w.Config.Tools) > 0 {
		systemPrompt += "\n\nAvailable tools: " + strings.Join(w.Config.Tools, ", ")
	}

	p := fmt.Sprintf("System Prompt: %s\n\nTask %s: %s\nContext: %s", systemPrompt, task.ID, task.Goal, task.Context)
	if task.Expected != "" {
		p += fmt.Sprintf("\nExpected Outcome/Format: %s", task.Expected)
	}

	out, err := w.LLM.Complete(ctx, p, nil)
	r := SwarmResult{AgentName: w.Config.Name, Role: RoleWorker, Output: out, Tokens: len(out) / 4}
	if err != nil {
		r.Error = err.Error()
	}
	o.Results <- r
}

func (o *SwarmOrchestrator) mergeResults(ctx context.Context, goal, ctxStr string, results []SwarmResult) (string, error) {
	var parts []string
	for _, r := range results {
		label := fmt.Sprintf("[%s - %s]", r.AgentName, r.Role)
		if r.Error != "" {
			parts = append(parts, label+" ERROR: "+r.Error)
		} else {
			parts = append(parts, label+"\n"+r.Output)
		}
	}
	merged := strings.Join(parts, "\n\n---\n\n")
	orch := o.findAgent(RoleOrchestrator)
	if orch != nil && orch.LLM != nil {
		p := fmt.Sprintf("Merge:\nGoal: %s\nContext: %s\nOutputs:\n%s", goal, ctxStr, merged)
		if final, err := orch.LLM.Complete(ctx, p, nil); err == nil {
			return final, nil
		}
	}
	return merged, nil
}

func (o *SwarmOrchestrator) collectResults() []SwarmResult {
	close(o.Results)
	var res []SwarmResult
	for r := range o.Results {
		res = append(res, r)
	}
	o.Results = make(chan SwarmResult, 100)
	return res
}

func (o *SwarmOrchestrator) findAgent(role SwarmRole) *SwarmAgent {
	for _, a := range o.Agents {
		if a.Config.Role == role {
			return a
		}
	}
	return nil
}

func (o *SwarmOrchestrator) getWorkers() []*SwarmAgent {
	var workers []*SwarmAgent
	for _, a := range o.Agents {
		if a.Config.Role == RoleWorker {
			workers = append(workers, a)
		}
	}
	return workers
}

func DefaultSwarmConfig() SwarmConfig {
	return SwarmConfig{
		Agents: []SwarmAgentConfig{
			{Name: "orchestrator", Role: RoleOrchestrator, Provider: "gemini", SystemPrompt: "You are a swarm orchestrator."},
			{Name: "generator", Role: RoleGenerator, Provider: "gemini", SystemPrompt: "Generate solutions."},
			{Name: "discriminator", Role: RoleDiscriminator, Provider: "deepseek", SystemPrompt: "Critique solutions."},
			{Name: "worker-1", Role: RoleWorker, Provider: "gemini", SystemPrompt: "Execute tasks."},
		},
		Topology: "centralized", MaxWorkers: 3, TimeoutSec: 120,
	}
}

func runSwarm() {
	root := getRoot()
	cfg := loadAppCfg(root)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🐝 Swarm Mode")
	fmt.Println("1) Centralized  2) Sequential  3) GAN")
	fmt.Print("Choice [1]: ")
	c, _ := reader.ReadString('\n')
	topo := "centralized"
	switch strings.TrimSpace(c) {
	case "2":
		topo = "sequential"
	case "3":
		topo = "gan"
	}
	sc := DefaultSwarmConfig()
	sc.Topology = topo
	orch := NewSwarmOrchestrator(sc)
	fmt.Printf("Topology: %s | Agents: %d\n", topo, len(sc.Agents))
	fmt.Print("Goal: ")
	goal, _ := reader.ReadString('\n')
	goal = strings.TrimSpace(goal)
	if goal == "" {
		goal = "General"
	}
	fmt.Print("Context (optional): ")
	ctxStr, _ := reader.ReadString('\n')
	ctxStr = strings.TrimSpace(ctxStr)
	for _, ac := range sc.Agents {
		if cfg != nil {
			p, err := buildSingleProvider(ac.Provider, cfg)
			if err == nil {
				orch.BindLLM(ac.Name, p)
			}
		}
	}
	fmt.Println("\nRunning...")
	// Use ui.StartSpinner from internal
	result, swarmErr := orch.ExecuteSwarm(goal, ctxStr)
	if swarmErr != nil {
		fmt.Printf("\nWarning: %v\n", swarmErr)
	}
	fmt.Printf("\nResult:\n%s\n", result)
}

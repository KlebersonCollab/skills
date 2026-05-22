
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------
// Types
// ---------------------------------------------------------------------------

type SwarmRole string

const (
	RoleOrchestrator SwarmRole = "orchestrator"
	RoleWorker       SwarmRole = "worker"
	RoleGenerator    SwarmRole = "generator"
	RoleDiscriminator SwarmRole = "discriminator"
)

type SwarmAgentConfig struct {
	Name        string     `json:"name"`
	Role        SwarmRole  `json:"role"`
	Provider    string     `json:"provider"`
	SystemPrompt string    `json:"system_prompt"`
	Tools       []string   `json:"tools"` // "read_file", "write_file", etc.
}

type SwarmConfig struct {
	Agents      []SwarmAgentConfig `json:"agents"`
	Topology    string             `json:"topology"`     // "centralized", "gan", "sequential"
	MaxWorkers  int                `json:"max_workers"`
	TimeoutSec  int                `json:"timeout_sec"`
}

// SwarmOrchestrator manages the swarm lifecycle
type SwarmOrchestrator struct {
	Config     SwarmConfig
	Agents     map[string]*SwarmAgent
	Results    chan SwarmResult
	Done       chan struct{}
	mu         sync.Mutex
}

// SwarmAgent represents a single agent instance
type SwarmAgent struct {
	Config  SwarmAgentConfig
	LLM     func(prompt string) (string, error) // simplified LLM call
}

// SwarmResult holds the output of a single agent
type SwarmResult struct {
	AgentName string `json:"agent_name"`
	Role      SwarmRole `json:"role"`
	Output    string `json:"output"`
	Tokens    int    `json:"tokens"`
	Error     string `json:"error,omitempty"`
	Duration  time.Duration `json:"duration"`
}

// SubTask is a unit of work assigned to a worker
type SubTask struct {
	ID       string `json:"id"`
	Goal     string `json:"goal"`
	Context  string `json:"context"`
	Expected string `json:"expected,omitempty"`
}

// ---------------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------------

func NewSwarmOrchestrator(cfg SwarmConfig) *SwarmOrchestrator {
	agents := make(map[string]*SwarmAgent)
	for _, ac := range cfg.Agents {
		agents[ac.Name] = &SwarmAgent{
			Config: ac,
			LLM:    func(prompt string) (string, error) { return "", fmt.Errorf("LLM not bound") },
		}
	}
	return &SwarmOrchestrator{
		Config:  cfg,
		Agents:  agents,
		Results: make(chan SwarmResult, 100),
		Done:    make(chan struct{}),
	}
}

// BindLLM attaches a real LLM caller to an agent
func (o *SwarmOrchestrator) BindLLM(name string, fn func(prompt string) (string, error)) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	agent, ok := o.Agents[name]
	if !ok {
		return fmt.Errorf("agent '%s' not found", name)
	}
	agent.LLM = fn
	return nil
}

// ---------------------------------------------------------------------------
// Orchestrator: Decompose → Delegate → Collect → Merge
// ---------------------------------------------------------------------------

// DecomposeTask splits a user goal into subtasks
func (o *SwarmOrchestrator) DecomposeTask(goal string, context string) ([]SubTask, error) {
	orchestrator, ok := o.Agents["orchestrator"]
	if !ok {
		// Auto-create orchestrator if not explicitly defined
		orchestrator = &SwarmAgent{
			Config: SwarmAgentConfig{
				Name: "orchestrator",
				Role: RoleOrchestrator,
			},
			LLM: func(prompt string) (string, error) { return "", fmt.Errorf("no LLM") },
		}
	}

	prompt := fmt.Sprintf(`You are a swarm orchestrator. Decompose the following goal into subtasks.
Goal: %s
Context: %s

Output a JSON array of subtasks, each with: id, goal, context, expected.
Maximum %d subtasks.
`, goal, context, o.Config.MaxWorkers)

	resp, err := orchestrator.LLM(prompt)
	if err != nil {
		return nil, err
	}

	// Try to parse JSON response
	var tasks []SubTask
	respClean := strings.TrimSpace(resp)
	// Extract JSON array if wrapped in markdown
	if idx := strings.Index(respClean, "["); idx >= 0 {
		if endIdx := strings.LastIndex(respClean, "]"); endIdx > idx {
			respClean = respClean[idx : endIdx+1]
		}
	}
	if err := json.Unmarshal([]byte(respClean), &tasks); err != nil {
		// Fallback: create a single generic subtask
		tasks = []SubTask{{
			ID:      "task-1",
			Goal:    goal,
			Context: context,
		}}
	}
	return tasks, nil
}

// ExecuteSwarm runs the full swarm lifecycle based on topology
func (o *SwarmOrchestrator) ExecuteSwarm(goal string, context string) (string, error) {
	switch o.Config.Topology {
	case "gan":
		return o.executeGAN(goal, context)
	case "sequential":
		return o.executeSequential(goal, context)
	default: // "centralized"
		return o.executeCentralized(goal, context)
	}
}

// executeCentralized: Orchestrator decomposes → Workers execute → Merge
func (o *SwarmOrchestrator) executeCentralized(goal string, context string) (string, error) {
	tasks, err := o.DecomposeTask(goal, context)
	if err != nil {
		return "", fmt.Errorf("task decomposition failed: %w", err)
	}

	var wg sync.WaitGroup
	timeout := time.Duration(o.Config.TimeoutSec) * time.Second

	for _, task := range tasks {
		wg.Add(1)
		go func(t SubTask) {
			defer wg.Done()
			o.executeWorker(t)
		}(task)
	}

	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// all workers finished
	case <-time.After(timeout):
		return "", fmt.Errorf("swarm execution timed out after %d seconds", o.Config.TimeoutSec)
	}

	// Collect results
	results := o.collectResults()
	return o.mergeResults(goal, context, results)
}

// executeSequential: Agents execute one after another, passing context
func (o *SwarmOrchestrator) executeSequential(goal string, context string) (string, error) {
	currentContext := context
	for _, agent := range o.Agents {
		if agent.Config.Role == RoleOrchestrator {
			continue // skip orchestrator in sequential
		}
		prompt := fmt.Sprintf("Goal: %s\nContext:\n%s\n\nAgent '%s' (%s): %s\nProvide your output.", goal, currentContext, agent.Config.Name, agent.Config.Role, agent.Config.SystemPrompt)
		output, err := agent.LLM(prompt)
		if err != nil {
			return "", fmt.Errorf("agent '%s' failed: %w", agent.Config.Name, err)
		}
		o.Results <- SwarmResult{
			AgentName: agent.Config.Name,
			Role:      agent.Config.Role,
			Output:    output,
			Tokens:    len(output) / 4,
		}
		currentContext = output
	}
	results := o.collectResults()
	return o.mergeResults(goal, context, results)
}

// executeGAN: Generator → Discriminator loop with refinement
func (o *SwarmOrchestrator) executeGAN(goal string, context string) (string, error) {
	gen := o.findAgent(RoleGenerator)
	disc := o.findAgent(RoleDiscriminator)

	if gen == nil || disc == nil {
		return "", fmt.Errorf("GAN mode requires both a 'generator' and a 'discriminator' agent")
	}

	maxIterations := 3
	currentProposal := ""

	for i := 0; i < maxIterations; i++ {
		// Generator
		prompt := fmt.Sprintf("Generate a solution for:\nGoal: %s\nContext:\n%s\nPrevious attempt: %s\nIteration %d/%d", goal, context, currentProposal, i+1, maxIterations)
		genOutput, err := gen.LLM(prompt)
		if err != nil {
			return "", fmt.Errorf("generator failed: %w", err)
		}
		currentProposal = genOutput

		o.Results <- SwarmResult{
			AgentName: gen.Config.Name,
			Role:      RoleGenerator,
			Output:    genOutput,
			Duration:  0,
		}

		// Discriminator
		discPrompt := fmt.Sprintf("Critique the following solution for correctness, security, and completeness. Respond with 'APPROVED' if acceptable, or list specific issues:\nGoal: %s\nSolution:\n%s", goal, currentProposal)
		discOutput, err := disc.LLM(discPrompt)
		if err != nil {
			return "", fmt.Errorf("discriminator failed: %w", err)
		}

		o.Results <- SwarmResult{
			AgentName: disc.Config.Name,
			Role:      RoleDiscriminator,
			Output:    discOutput,
			Duration:  0,
		}

		// Check approval
		if strings.Contains(strings.ToUpper(discOutput), "APPROVED") {
			return currentProposal, nil
		}

		// Feed critique back for next iteration
		currentProposal = fmt.Sprintf("Proposal:\n%s\n\nCritique:\n%s", currentProposal, discOutput)
	}

	// Return last proposal even if not approved
	return currentProposal, fmt.Errorf("GAN loop reached max iterations without approval")
}

// ---------------------------------------------------------------------------
// Workers
// ---------------------------------------------------------------------------

func (o *SwarmOrchestrator) executeWorker(task SubTask) {
	agents := o.getWorkers()
	if len(agents) == 0 {
		o.Results <- SwarmResult{
			AgentName: "unknown",
			Role:      RoleWorker,
			Output:    "No workers configured",
			Error:     "no workers available",
		}
		return
	}

	// Pick a random worker
	agent := agents[rand.Intn(len(agents))]
	start := time.Now()

	prompt := fmt.Sprintf(`You are a worker agent '%s'.
Task ID: %s
Goal: %s
Context: %s
Expected outcome: %s

Execute this task to the best of your ability using any available tools.`, agent.Config.Name, task.ID, task.Goal, task.Context, task.Expected)

	output, err := agent.LLM(prompt)
	duration := time.Since(start)

	result := SwarmResult{
		AgentName: agent.Config.Name,
		Role:      RoleWorker,
		Output:    output,
		Tokens:    len(output) / 4,
		Duration:  duration,
	}
	if err != nil {
		result.Error = err.Error()
	}

	o.Results <- result
}

// ---------------------------------------------------------------------------
// Merge
// ---------------------------------------------------------------------------

func (o *SwarmOrchestrator) mergeResults(goal string, context string, results []SwarmResult) (string, error) {
	var parts []string
	for _, r := range results {
		label := fmt.Sprintf("[%s - %s]", r.AgentName, r.Role)
		if r.Error != "" {
			parts = append(parts, fmt.Sprintf("%s ERROR: %s", label, r.Error))
		} else {
			parts = append(parts, fmt.Sprintf("%s\n%s", label, r.Output))
		}
	}

	merged := strings.Join(parts, "\n\n---\n\n")

	// If orchestrator exists, let it do a final merge
	orchestrator := o.findAgent(RoleOrchestrator)
	if orchestrator != nil {
		finalPrompt := fmt.Sprintf("Merge the following worker outputs into a cohesive final response for the user.\nOriginal goal: %s\nContext: %s\n\nWorker outputs:\n%s", goal, context, merged)
		final, err := orchestrator.LLM(finalPrompt)
		if err == nil {
			return final, nil
		}
	}

	return merged, nil
}

func (o *SwarmOrchestrator) collectResults() []SwarmResult {
	close(o.Results)
	var results []SwarmResult
	for r := range o.Results {
		results = append(results, r)
	}
	o.Results = make(chan SwarmResult, 100) // reset channel
	return results
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func (o *SwarmOrchestrator) findAgent(role SwarmRole) *SwarmAgent {
	for _, agent := range o.Agents {
		if agent.Config.Role == role {
			return agent
		}
	}
	return nil
}

func (o *SwarmOrchestrator) getWorkers() []*SwarmAgent {
	var workers []*SwarmAgent
	for _, agent := range o.Agents {
		if agent.Config.Role == RoleWorker {
			workers = append(workers, agent)
		}
	}
	return workers
}

// DefaultSwarmConfig returns a sensible default swarm configuration
func DefaultSwarmConfig() SwarmConfig {
	return SwarmConfig{
		Agents: []SwarmAgentConfig{
			{
				Name:         "orchestrator",
				Role:         RoleOrchestrator,
				Provider:     "gemini",
				SystemPrompt: "You are a swarm orchestrator. Decompose tasks and merge results.",
				Tools:        []string{},
			},
			{
				Name:         "generator",
				Role:         RoleGenerator,
				Provider:     "gemini",
				SystemPrompt: "You are a solution generator. Propose creative and thorough solutions.",
				Tools:        []string{"read_file", "search_files"},
			},
			{
				Name:         "discriminator",
				Role:         RoleDiscriminator,
				Provider:     "deepseek",
				SystemPrompt: "You are a solution critic. Analyze proposals for correctness, security, and completeness.",
				Tools:        []string{"read_file"},
			},
			{
				Name:         "worker-1",
				Role:         RoleWorker,
				Provider:     "gemini",
				SystemPrompt: "You are a worker agent. Execute assigned tasks with precision.",
				Tools:        []string{"read_file", "write_file", "search_files", "execute_command"},
			},
		},
		Topology:    "centralized",
		MaxWorkers:  3,
		TimeoutSec:  120,
	}
}
	config, _ := LoadConfig(configPath)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🐝 Harness Multi-Agent Swarm")
	fmt.Println("1) Centralized  2) Sequential  3) GAN")
	fmt.Print("Escolha (1-3) [1]: ")
	choice, _ := reader.ReadString('\n')
	topology := "centralized"
	switch strings.TrimSpace(choice) {
	case "2":
		topology = "sequential"
	case "3":
		topology = "gan"
	}
	cfg := DefaultSwarmConfig()
	cfg.Topology = topology
	orch := NewSwarmOrchestrator(cfg)
	fmt.Printf("Topology: %s | Agents: %d\n", topology, len(cfg.Agents))
	fmt.Print("\nObjetivo: ")
	goal, _ := reader.ReadString('\n')
	goal = strings.TrimSpace(goal)
	if goal == "" {
		goal = "Sessão exploratória geral"
	}
	fmt.Print("Contexto (opcional): ")
	context, _ := reader.ReadString('\n')
	context = strings.TrimSpace(context)
	for name := range orch.Agents {
		agentName := name
		if config != nil {
			orch.BindLLM(agentName, func(prompt string) (string, error) {
				agent := orch.Agents[agentName]
				augmentedPrompt := agent.Config.SystemPrompt
				if augmentedPrompt != "" {
					augmentedPrompt += "\n\n"
				}
				augmentedPrompt += prompt
				if config.Providers[config.ActiveProvider].Streaming {
					fmt.Printf(" [%s streaming...] ", agentName)
					return CallLLMStream(config, augmentedPrompt, "")
				}
				return CallLLM(config, augmentedPrompt, "")
			})
		} else {
			orch.BindLLM(agentName, func(prompt string) (string, error) {
				return fmt.Sprintf("[%s - Demo]\nRecebido: %.200s...\n(Configure \"harness init\" para LLM real)", agentName, prompt), nil
			})
		}
	}
	fmt.Println("\n⚙️  Swarm em execução...")
	stopSpinner := startSpinner(fmt.Sprintf("Swarm %s com %d agentes", topology, len(cfg.Agents)))
	result, swarmErr := orch.ExecuteSwarm(goal, context)
	close(stopSpinner)
	if swarmErr != nil {
		fmt.Printf("\n⚠️  Concluído com ressalvas: %v\n", swarmErr)
	}
	fmt.Printf("\n📋 Resultado:\n%s\n", result)
	if config != nil {
		sessionHash := md5.Sum([]byte(time.Now().Format(time.RFC3339) + goal))
		sessionID := "swarm-" + hex.EncodeToString(sessionHash[:4])
		sessionFile := filepath.Join(root, ".harness", "sessions", sessionID+".json")
		tree := NewSessionTree(sessionID, "Swarm: "+goal)
		tree.AddNode("system", fmt.Sprintf("topology=%s, agents=%d", topology, len(cfg.Agents)), 0)
		tree.AddNode("user", goal, len(goal)/4)
		tree.AddNode("assistant", result, len(result)/4)
		tree.Save(sessionFile)
		fmt.Printf("💾 Sessão: %s\n", sessionFile)
	}
}

func filepathWalkDir(root string, fn func(path string, d fs.DirEntry, err error) error) error {
	return filepath.WalkDir(root, fn)
}


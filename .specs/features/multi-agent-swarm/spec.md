
# Specification — Multi-Agent Swarm

## ID: `multi-agent-swarm`

## Objective
Extend the Harness Go Agent to support multi-agent swarm orchestration, enabling multiple AI agents to collaborate on complex tasks with distinct roles, parallel execution, and coordinated handoffs.

## Functional Requirements

### FR-1: Swarm Configuration
- Swarm topology defined in `.harness/config.json`
- Each agent has: name, role (orchestrator/worker/generator/discriminator), provider, tools, system prompt
- Support for 3+ agents in a single swarm

### FR-2: Orchestrator Role
- Decomposes user input into subtasks
- Assigns subtasks to worker agents
- Merges worker results into final response
- Handles error recovery and re-delegation

### FR-3: Worker Role
- Receives a well-defined subtask with context
- Executes tools and LLM calls independently
- Reports structured results back to orchestrator
- Runs in isolated context (separate session branch)

### FR-4: GAN Pattern Support
- **Generator**: Creates solution proposals
- **Discriminator**: Audits proposals for quality, security, and correctness
- Iterative refinement loop between G and D

### FR-5: Parallel Execution
- Workers can execute concurrently via goroutines
- Orchestrator collects results via channels
- Configurable timeout per worker

### FR-6: Session DAG Integration
- Each agent's activity logged in the session tree as sub-branches
- Full traceability of which agent did what

## Acceptance Criteria

- [ ] **AC-1**: `harness swarm` command starts an interactive swarm session
- [ ] **AC-2**: Orchestrator can delegate to 2+ workers and merge results
- [ ] **AC-3**: GAN loop (Generator → Discriminator → Refine) works with 2 agents
- [ ] **AC-4**: Session tree shows sub-branches per agent
- [ ] **AC-5**: `go build` succeeds
- [ ] **AC-6**: `go test` passes
- [ ] **AC-7**: `make audit` 100% compliant

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "multi-agent-swarm"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-23T17:55:00Z"
evidence_checksum: "spec-draft-v1"
```

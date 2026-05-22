
# Tasks — Multi-Agent Swarm

## Phase 1: Core Types
- [ ] **T1.1**: Define SwarmConfig, SwarmAgent, SwarmRole types in `swarm.go`
- [ ] **T1.2**: Implement swarm config loading from AppConfig
- [ ] **T1.3**: Add SwarmConfig to AppConfig in `client.go`

## Phase 2: Orchestrator
- [ ] **T2.1**: Implement SwarmOrchestrator struct
- [ ] **T2.2**: Task decomposition (user input → subtasks)
- [ ] **T2.3**: Spawn workers via goroutines with isolated contexts
- [ ] **T2.4**: Result collection via channels

## Phase 3: GAN Pattern
- [ ] **T3.1**: Implement Generator agent role
- [ ] **T3.2**: Implement Discriminator agent role
- [ ] **T3.3**: Implement iterative refinement loop

## Phase 4: CLI & Integration
- [ ] **T4.1**: Add `harness swarm` CLI command
- [ ] **T4.2**: Integrate swarm session with DAG tree
- [ ] **T4.3**: `go test` + `make audit`

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "multi-agent-swarm"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-23T18:05:00Z"
evidence_checksum: "tasks-draft-v1"
```

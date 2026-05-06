# Vanderbilt Prompt Patterns

Advanced prompt engineering patterns to maximize agentic intelligence during brainstorming.

## 1. Flipped Interaction Pattern
Instead of you giving a solution, the agent asks you questions to build the context.
- **Prompt**: "From now on, I want you to ask me questions until you have enough information to [goal]. Ask me one question at a time."

## 2. Cognitive Verifier Pattern
The agent breaks a complex problem into smaller questions to verify its own understanding before providing a final answer.
- **Prompt**: "When I ask you a question, generate 3 additional questions that would help you give a better answer. Answer those questions first, then give the final response."

## 3. Persona Pattern
Assign a specific role to the agent to change its perspective.
- **Prompt**: "Act as a [Senior Architect/UX Specialist/Product Manager] with a focus on [attribute]. Critique the following proposal."

## 4. Alternative Terminology Pattern
Ask for different ways to describe a problem to find new solutions.
- **Prompt**: "Describe [challenge] using 5 different metaphors or industry terminologies. How does each description change the potential solution?"

## 5. Template Pattern
Force the agent to respond in a specific format to ensure completeness.
- **Prompt**: "For every idea you generate, follow this template: Idea | Rationale | Potential Risk | Mitigation."


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.375414Z"
evidence_checksum: "NONE"
```

# Divergence Techniques for Software Design

Effective brainstorming requires a phase of divergence where multiple ideas are generated without immediate judgment. Use these techniques to explore the solution space.

## 1. SCAMPER (Substitute, Combine, Adapt, Modify, Put to another use, Eliminate, Reverse)
A checklist to trigger new ideas by changing existing elements of a system.
- **Substitute**: Replace a library with a more modern one.
- **Combine**: Merge two features into a single unified interface.
- **Adapt**: Reuse a pattern from a different domain (e.g., using a game engine pattern in an enterprise UI).
- **Modify**: Change the scale of a feature (e.g., from synchronous to asynchronous processing).
- **Put to another use**: Use a monitoring tool for business analytics.
- **Eliminate**: Remove redundant steps in a user workflow (KISS).
- **Reverse**: Change the order of operations or the perspective of the user.

## 2. 5 Whys (Root Cause Analysis)
A simple but powerful technique to dig deeper into a problem.
- **Goal**: Ask "Why?" five times until the root cause of a requirement or bug is found.
- **Example**: 
  1. Why is the application slow? (Response: Because the database is busy).
  2. Why is the database busy? (Response: Because of many concurrent read operations).
  3. Why are there so many read operations? (Response: Because we are not caching the results).
  4. Why are we not caching the results? (Response: Because the cache invalidation logic was too complex).
  5. Why was it too complex? (Response: Because the data structure is not optimized for caching).

## 3. First Principles (Deconstruction)
Break down complex problems into their most basic, fundamental truths.
- **Step 1**: Identify current assumptions.
- **Step 2**: Breakdown the problem into its most basic parts.
- **Step 3**: Create a new solution from scratch using only the fundamental truths.
- **Example**: Instead of asking "How can we improve the existing build script?", ask "What is the minimum required to transform source code into a deployable artifact?".

## 4. Mind Mapping
Visually map out ideas and their relationships to see the big picture and find hidden connections.

## 5. Brainwriting
A collaborative technique where team members write down their ideas independently before sharing them to avoid groupthink.

## 6. Random Word Association
Pick a random word and force a connection to the problem. This can spark unexpected and creative solutions.

## Best Practices for AI Brainstorming
1. **Defer Judgment**: Generate many ideas first, then evaluate them in a separate turn.
2. **Seek Divergent Perspectives**: Ask the agent to play different roles (e.g., "Think like a Security Expert", "Think like a UI Designer").
3. **Encourage Wild Ideas**: Don't limit yourself to what is currently implemented; explore what is technically possible.

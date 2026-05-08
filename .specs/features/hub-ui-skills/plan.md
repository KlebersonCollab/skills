# Engineering Plan: Agent Skills Hub

## 🏗 Architecture
The application will be a single-page application (SPA) built with Vite and React.

### Components:
- `SkillRegistry`: A utility script to parse `SKILL.md` files.
- `Dashboard`: Main layout with Sidebar and Main Content.
- `SkillCard`: Card component for the list.
- `EcosystemGraph`: Interactive canvas using `React Flow`.
- `DetailPanel`: Slide-over component for markdown rendering.

### Tech Stack:
- **Frontend**: Vite + React + Vanilla CSS (Design Tokens).
- **Graphing**: React Flow.
- **Markdown**: `react-markdown`.
- **Parsing**: `gray-matter` (for YAML frontmatter).

## 📊 Data Schema
```json
{
  "skills": [
    {
      "id": "sdd",
      "name": "System Design Document",
      "version": "2.3.0",
      "category": "core",
      "description": "...",
      "status": "COMPLETED",
      "path": ".agents/skills/sdd/SKILL.md",
      "conversesWith": ["brainstorming", "architecture"]
    }
  ]
}
```

## 🚀 Implementation Phases

### Phase 1: Data Scaffolding
1. Create `scripts/update-registry.js` to scan `.agents/skills/`.
2. Generate `public/registry.json`.

### Phase 2: UI Foundation
1. Initialize Vite project.
2. Setup CSS variables based on `DESIGN.md`.
3. Create Layout components.

### Phase 3: Graph Integration
1. Implement `EcosystemGraph` with `React Flow`.
2. Define node and edge types.

### Phase 4: Polish & Delivery
1. Add animations.
2. Final audit against `spec.md`.

<!-- @sdd-state -->
```yaml
version: "1.0.0"
feature_id: "hub-ui-skills"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-08T23:03:00Z"
evidence_checksum: "NONE"
```

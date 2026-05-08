# Feature Specification: Agent Skills Hub

## Overview
A centralized dashboard and visualization tool for the Agent Hub's skills. It provides a visual representation of available skills, their metadata, and their logical interconnections (how they "talk" to each other).

## 🎯 Goals
- Provide a high-level overview of all skills in `.agents/skills/`.
- Visualize logical relationships and dependencies between skills using an interactive graph.
- Enable quick access to skill metadata (version, category, status).
- Follow the design aesthetics defined in `DESIGN.md`.

## 🚫 Non-Goals
- Real-time agent execution or prompt testing (this is a governance/visualization tool).
- Editing skill files directly from the UI (v1).

## 👥 User Personas
- **Developer**: Needs to know what tools are available and how to compose them.
- **Project Lead**: Needs to audit the status and versioning of the skills ecosystem.

## 🛠 Functional Requirements (FR)
- **FR1: Skill Discovery**: Search and filter skills by name, category, or status.
- **FR2: Metadata Display**: View `name`, `version`, `description`, and `@sdd-state` for each skill.
- **FR3: Relationship Graph**: An interactive graph showing how skills connect (e.g., `sdd` at the center, connecting to others).
- **FR4: Side Detail Panel**: Detailed view of a selected skill's `SKILL.md` content.

## 📐 Non-Functional Requirements (NFR)
- **Performance**: Initial load under 2 seconds.
- **Aesthetics**: Premium Dark Mode, Glassmorphism, Modern Typography.
- **Responsiveness**: Optimized for Desktop and Tablet.

## ✅ Acceptance Criteria (AC)
- [ ] UI lists all 14+ skills from the `.agents/skills/` directory.
- [ ] Each skill card shows correct version and status badges.
- [ ] The Graph View correctly links skills based on SDD phases.
- [ ] Searching for a skill highlights it in the list and graph.
- [ ] Side panel displays the full content of the selected skill's `SKILL.md`.

<!-- @sdd-state -->
```yaml
version: "1.0.0"
feature_id: "hub-ui-skills"
phase: "SPECIFICATION"
status: "IN_PROGRESS"
last_update: "2026-05-08T19:25:00Z"
evidence_checksum: "NONE"
```

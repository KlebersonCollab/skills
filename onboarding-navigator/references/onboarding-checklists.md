# Onboarding Checklists: Action Plans

Standardized checklists to ensure that no important steps are forgotten during onboarding or the start of new projects.

---

## 1. New Team Member (Skynet Developer)

### A. Access and Environment
- [ ] Clone the central skills repository.
- [ ] Install Gemini CLI.
- [ ] Configure environment variables (e.g., `AZURE_DEVOPS_EXT_PAT`).
- [ ] Setup Python environment via **UV** (`python-uv`).
- [ ] Setup Flutter environment via **FVM** (`flutter-fvm`).

### B. Mandatory Readings (Source of Truth)
- [ ] `README.md` (Overview).
- [ ] `sdd/SKILL.md` (Working Methodology).
- [ ] `.specs/codebase/STACK.md` (What we use).
- [ ] `.specs/project/PROJECT.md` (Where we are going).

### C. First Commit
- [ ] Choose a "Small" task or "Bug fix" from the Roadmap.
- [ ] Execute the full SDD workflow (Spec -> Plan -> Tasks -> Impl -> Review).

---

## 2. Starting a New Feature (Module or Skill)

### A. Planning (Discovery & Specify)
- [ ] Create folder in `.specs/features/`.
- [ ] Audit PRD or initial Requirements.
- [ ] Write `spec.md` with BDD.
- [ ] Create `plan.md` with Mermaid Diagram.
- [ ] Generate atomic task list in `tasks.md`.

### B. Architecture
- [ ] Evaluate if the change requires an ADR (`architecture`).
- [ ] Review SOLID and DRY principles (`clean-code-mentor`).

### C. Execution and Closure
- [ ] Implement with tests (TDD).
- [ ] Audit via `skill-factory-validator` (if it is a skill).
- [ ] Proactively update the ROADMAP and memory logs.

---

## 3. Navigator Tip
"A checklist is not a burden; it is your safety net. Follow the steps and ensure your delivery is unquestionable."

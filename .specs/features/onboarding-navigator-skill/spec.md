# Spec: Onboarding Navigator Skill

## 1. Goal
Create a specialized skill that acts as an interactive guide and cultural mentor for the project ecosystem. This skill will leverage existing documentation (like `GUIA_CENTRAL.md`, `MAPA_NAVEGACAO.md`, and `Technical Decision Making`) to orient new team members and guide the creation of new modules, ensuring adherence to project standards and values.

## 2. Target Audience
New developers joining the team and existing developers starting new modules or features who need to ensure compliance with the global project guidelines.

## 3. Core Requirements (Functional)
- **Knowledge Base Navigation**: Ability to locate and explain project standards based on specific files (`GUIA_CENTRAL.md`, etc.).
- **Decision Making Support**: Guide developers through the "Technical Decision Making" process, helping to structure ADRs or RFCs.
- **Onboarding Checklists**: Generate personalized checklists for new members or new feature setups.
- **Standards Enforcement**: Actively suggest the correct templates (PRD, Spec, Plan) based on the context.
- **Cultural Mentorship**: Explain the "Why" behind project conventions (e.g., "Why simplicity first?").

## 4. Acceptance Criteria (BDD Format)

### Scenario 1: Orientation for New Feature
**Given** a developer wants to start a "New Feature"
**When** the `onboarding-navigator` is asked for guidance
**Then** it must provide a step-by-step guide pointing to the correct spec directory and the required templates (PRD/RFC) according to the complexity.

### Scenario 2: Technical Decision Support
**Given** a developer is unsure about an architectural choice
**When** the `onboarding-navigator` is invoked
**Then** it must ask the diagnostic questions defined in the "Technical Decision Making" guide and help draft an ADR.

### Scenario 3: Onboarding a New Member
**Given** a new member joins the "Skynet" project
**When** the `onboarding-navigator` generates an onboarding plan
**Then** it must include a list of mandatory reads (GUIA_CENTRAL, Culture) and initial environment setup steps.

## 5. Technical Requirements (Non-Functional)
- **Context-Aware**: Must prioritize content from the `.specs/` and local documentation folders.
- **Interactive Tone**: Professional, welcoming, and pedagogical.
- **Visual Mapping**: Use Mermaid diagrams to show the project structure and decision flows.

## 6. Architecture & Design
- `onboarding-navigator/SKILL.md`: Main entry point with the navigation workflow.
- `onboarding-navigator/references/`:
    - `decision-making-framework.md`: Guidance on how to choose and document technologies.
    - `project-structure-guide.md`: Explanation of the "Source of Truth" and folder hierarchy.
    - `onboarding-checklists.md`: Templates for different roles and project stages.
- `onboarding-navigator/examples/`:
    - `welcome-guide.md`: A sample interactive welcome message.
    - `decision-flow.mermaid`: Visual representation of the decision process.

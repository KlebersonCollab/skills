# Specification: Upgrade YouTube Transcript Skill to SDD v2.2.0

## 🎯 Objective
Retrofit the `youtube-transcript` skill to comply with the latest SDD v2.2.0 governance standards, ensuring structural purity, centralized memory, and alignment with the `skill-factory` architect mandates.

## 📋 Requirements
- **FR-1: Structural Purity**: Remove any local `STATE.md`, `MEMORY.md`, or `LEARNINGS.md` (if any existed, though none were found in `list_dir`).
- **FR-2: Centralized Context**: Update the "Context Check" in `SKILL.md` to point to `.specs/project/` artifacts.
- **FR-3: Mandatory Prerequisites**: Align the "Prerequisites" section with the `agent.md` mandates (Mental Checklist, SDD Verification).
- **FR-4: Audit Compliance**: Generate a standard `audit-report.md` using the `skill-factory-validator` logic.
- **FR-5: Registry Update**: Ensure the skill is correctly registered in the project's global catalog (if not already).

## ✅ Acceptance Criteria (BDD)
- **AC-1: Prerequisites Alignment**
  - **Given** I am reading `youtube-transcript/SKILL.md`
  - **When** I check the "Prerequisites" section
  - **Then** it must include the "Session Bootstrap" mental checklist and point to centralized memory.

- **AC-2: Memory Centralization**
  - **Given** the skill directory `youtube-transcript`
  - **When** I search for `STATE.md`, `MEMORY.md`, or `LEARNINGS.md`
  - **Then** these files must not exist locally.

- **AC-3: Validation Score**
  - **Given** the `audit-report.md`
  - **When** the audit is performed
  - **Then** the score must be >= 95/100.

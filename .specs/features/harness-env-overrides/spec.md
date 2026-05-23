# Specification: Harness Environment Overrides & Root .env Support

This feature enables seamless configuration and runtime orchestration of the Harness AI Agent using standard environment variables from a `.env` file at the root of the workspace.

## Goals

1. **Zero-Configuration Startup:** Allow Harness to run without requiring a `.harness/config.json` file, dynamically building a robust default configuration for all major providers (Gemini, DeepSeek, OpenAI, Anthropic, Groq, Ollama).
2. **Root `.env` Support:** Relocate the `.env` template and variables from the hidden `bin/harness/` directory to the project root, making it the central source of truth.
3. **Environment-Driven Controls:** Allow overriding active providers and models dynamically using `HARNESS_PROVIDER` / `ACTIVE_PROVIDER` and `HARNESS_MODEL` / `ACTIVE_MODEL` in `.env` or system environment.

## BDD Scenarios

### Scenario 1: Root .env is Loaded Successfully
* **Given** a `.env` file exists at the root of the workspace containing `GEMINI_API_KEY="test-gemini-key"`
* **When** Harness starts
* **Then** the root `.env` is loaded and `os.Getenv("GEMINI_API_KEY")` is set to `"test-gemini-key"`

### Scenario 2: Zero-Config Startup with Env Keys
* **Given** `.harness/config.json` does not exist
* **And** `.env` at the root contains `DEEPSEEK_API_KEY="test-deepseek-key"`
* **When** Harness starts
* **Then** it does not crash or exit
* **And** it automatically defaults to `deepseek` as the active provider with `deepseek-chat` as the model

### Scenario 3: Explicit Provider & Model Override via Env
* **Given** a `.env` exists with `HARNESS_PROVIDER="openai"` and `HARNESS_MODEL="gpt-4o-mini"`
* **When** Harness loads its configuration
* **Then** the active provider is set to `openai`
* **And** the active model is configured to `gpt-4o-mini`

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-ENV-OVERRIDES"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-23T02:46:00Z"
evidence_checksum: "ffb546e"
```

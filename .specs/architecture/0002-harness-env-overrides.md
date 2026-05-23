# ADR-0002: Environment-Driven Configuration Fallback & Overrides

We have decided to update the Harness configuration loader to prioritize and fallback to standard environment variables (loaded from `.env`) when `.harness/config.json` is missing or incomplete. This removes the strict dependency on the initialization wizard (`harness init`) and aligns Harness with standard 12-factor app environment configuration principles.

## Context

Previously, the Harness runtime required a valid `.harness/config.json` file to run, falling back to a fatal exit with the error: "Config not found. Run 'harness init' first."
Additionally, the `.env` template and `.env` files were kept inside the nested `bin/harness/` directory, which is non-standard and highly confusing for developers who expect a single `.env` at the workspace root (`/home/kleberson/Documentos/skills/`).
Lastly, even if environment overrides like `HARNESS_PROVIDER` were set, they would only be applied if that provider was already defined in the local `.harness/config.json`.

## Decision

1. **Root .env Prioritization:** Move the `.env` and `.env.example` templates to the root directory `/home/kleberson/Documentos/skills/` so they are immediately visible to the developer.
2. **Config-Less Dynamic Fallback:** If `.harness/config.json` is missing, Harness will not crash. Instead, it will automatically build a default internal configuration containing all standard LLM providers (Gemini, DeepSeek, OpenAI, Anthropic, Groq, Ollama) and select an active provider based on which API key is present in the `.env`.
3. **Environment-Driven Overrides:** Ensure `loadAppCfg` merges configured providers with standard default templates, so environment variables (`HARNESS_PROVIDER`, `ACTIVE_PROVIDER`, `HARNESS_MODEL`, `ACTIVE_MODEL`) can dynamically select and configure any standard provider even if it was not explicitly declared in the `.harness/config.json` file.
4. **Resilient Variable Resolution:** Allow environment variables defined in `.env` at the root to seamlessly propagate into LLM requests without hardcoded template constraints.

## Status

Accepted

## Consequences

* **Improved Developer Experience (DX):** Zero-config start is now supported. If `.env` is populated with API keys, `make harness-start` just works.
* **Standard Conformity:** Aligns the project with modern web and agentic best practices where `.env` acts as the single source of truth for runtime secrets and options.
* **Backward Compatibility:** Existing custom configurations in `.harness/config.json` remain fully supported and functional.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-ENV-OVERRIDES"
phase: "ALIGN"
status: "COMPLETED"
last_update: "2026-05-23T02:45:00Z"
evidence_checksum: "ffb546e"
```

# 🏛️ API Architect Skill

> Skill for designing, governing, and architecting interoperable, secure, and resilient API ecosystems.

---

## 📖 Overview

The **API Architect** skill (formerly `api-designer`) elevates the API creation process from simple "contract design" to a complete architecture. It integrates modern design patterns (REST, GraphQL, tRPC) with critical layers of governance, security, and resilience.

Inspired by industry standards and reinforced by API governance practices, this skill ensures that every endpoint is consistent across the entire ecosystem.

---

## 🚀 Key Features

- **API Style Guide**: Definition of response standards (envelopes), errors, and naming conventions to ensure the API is predictable.
- **Multi-Protocol**: Native support for deciding between REST, GraphQL, and tRPC based on the technical context.
- **Security (OWASP)**: Systematic auditing against key API security risks (BOLA, IDOR, Mass Assignment).
- **Pragmatic Resilience**: Rate Limiting, Timeouts, and Circuit Breaker patterns integrated into the design.
- **Contract-First**: Total focus on specification (OpenAPI, SDL, Zod) before implementation.

---

## 📂 Skill Structure

```
api-architect/
├── SKILL.md                 # Technical definition and workflow
├── README.md                # This documentation
├── CHANGELOG.md             # Version history
├── references/              # Standards and reference guides
│   ├── api-style.md         # Style Guide (Envelopes, Errors, Pagination)
│   ├── rest-best-practices.md
│   ├── api-security-guide.md
│   └── api-rate-limiting.md
└── resources/
    └── implementation-playbook.md
```

---

## 🛠️ How to Use

1. **Invoke the skill** when planning a new feature that requires data exposure over the network.
2. Follow the **4 Phases** (Strategy, Design, Resilience, Validate) to ensure no critical detail is forgotten.
3. Use the **references** to define the format of responses and errors at the beginning of modeling.

---

## 📄 License

This module is part of the [Skills Hub](https://github.com/KlebersonCollab/skills) and is under the MIT license.

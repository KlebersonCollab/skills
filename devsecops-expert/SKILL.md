---
name: "devsecops-expert"
version: "1.0.0"
description: "Skill focada em DevSecOps, Segurança da Informação, SAST, DAST e práticas de Hardening em infraestrutura e código."
uses:
  - "architecture"
  - "git-workflow"
references:
  - "sdd"
---

# 🛡️ DevSecOps Expert Skill

## 1. Descrição e Propósito
A skill `devsecops-expert` eleva a qualidade de segurança de software dentro do ecossistema, orientando o agente a auditar dependências, validar configuração e adotar práticas "Secure-by-Design".

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros?
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

Antes de implementar ou validar soluções que envolvam infraestrutura, chaves, rotas públicas ou pipelines:
- **Verificar Secret Leaks**: Nunca comite chaves hardcoded, senhas ou tokens.
- **Validar Superfície de Ataque**: Toda porta ou endpoint exposto deve usar RBAC/Autenticação.
- **SAST Inicial**: Garantir que as dependências não possuem CVEs abertas antes de rodar os scripts.

## Goal
Garantir e blindar o ciclo de desenvolvimento contínuo (DevSecOps), fornecendo auditoria estática (SAST) e dinâmica (DAST) em código e infra.

## Output Structure
A execução desta skill produz os seguintes artefatos:
| Artefato | Localização | Descrição |
|----------|-------------|-----------|
| **Validation Report** | `validation-report.md` | Relatório de segurança. |

## Quality Rules
- Tolerância zero para vazamento de credenciais.
- Toda modificação estrutural precisa passar por SAST.

## Prohibited
- NUNCA suba chaves no versionamento.
- NUNCA ignore falhas de injeção ou CVEs expostas em logs.

## 3. Padrões de Segurança (Mandatory)
1. **Zero Trust Architecture**: Nenhum microserviço ou componente confia em outro nativamente sem tokenização/mutual TLS.
2. **Dependency Audit**: Exigir uso de ferramentas de varredura (ex: `bandit`, `trivy` ou `npm audit`).
3. **OWASP Top 10**: Prevenção ativa contra injeções SQL/NoSQL, XSS, CSRF e quebra de controle de acesso (BOLA).

## 4. Guia de Revisão Adversarial (Security Gate)
- O `harness-expert` pode chamar esta skill no final de uma refatoração arquitetural para "atacar" as decisões lógicas criadas (Threat Modeling / STRIDE).

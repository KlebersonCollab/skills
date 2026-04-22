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

## 2. 🔒 Prerequisites (Mandatory)
Antes de implementar ou validar soluções que envolvam infraestrutura, chaves, rotas públicas ou pipelines:
- **Verificar Secret Leaks**: Nunca comite chaves hardcoded, senhas ou tokens.
- **Validar Superfície de Ataque**: Toda porta ou endpoint exposto deve usar RBAC/Autenticação.
- **SAST Inicial**: Garantir que as dependências não possuem CVEs abertas antes de rodar os scripts.

## 3. Padrões de Segurança (Mandatory)
1. **Zero Trust Architecture**: Nenhum microserviço ou componente confia em outro nativamente sem tokenização/mutual TLS.
2. **Dependency Audit**: Exigir uso de ferramentas de varredura (ex: `bandit`, `trivy` ou `npm audit`).
3. **OWASP Top 10**: Prevenção ativa contra injeções SQL/NoSQL, XSS, CSRF e quebra de controle de acesso (BOLA).

## 4. Guia de Revisão Adversarial (Security Gate)
- O `harness-expert` pode chamar esta skill no final de uma refatoração arquitetural para "atacar" as decisões lógicas criadas (Threat Modeling / STRIDE).

# Skill Factory — Core Framework

> Every skill deserves a solid foundation. Consistency breeds quality.

[![Versão](https://img.shields.io/badge/Versão-1.0.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-2-brightgreen)](#-sub-skills)

---

## 📖 Visão Geral

Skill dedicada à **criação padronizada de novas skills** no Skills Hub. O Skill Factory automatiza o scaffolding, validação e registro, garantindo que toda nova skill siga os mesmos padrões de qualidade e estrutura.

Em vez de criar arquivos manualmente e risco de inconsistência, o agente executa esta skill que:
1. **Gera** todos os arquivos necessários a partir de templates padronizados.
2. **Valida** a conformidade da skill contra o checklist do hub.
3. **Registra** a skill no `README.md` raiz.

> **Princípio**: Se a skill SDD garante que o código é especificado antes de ser escrito, o Skill Factory garante que toda skill é estruturada antes de ser definida.

---

## ⚙️ Auto-Sizing

| Mode | Quando Usar | Arquivos Gerados |
|------|-------------|-----------------|
| **Quick** | Skill simples, sem sub-skills | `SKILL.md`, `README.md`, `CHANGELOG.md` |
| **Standard** | Skill com sub-skills | Todos acima + `<skill>-<sub>.skill.md` (N) |

---

## 🧩 Sub-skills

| Sub-skill | Arquivo | Responsabilidade |
|-----------|---------|------------------|
| **Bootstrap** | [skill-factory-bootstrap.skill.md](skill-factory-bootstrap.skill.md) | Gera o scaffolding completo da nova skill a partir dos templates padronizados. |
| **Validator** | [skill-factory-validator.skill.md](skill-factory-validator.skill.md) | Audita a skill criada contra o checklist de conformidade e emite relatório COMPLIANT/NON-COMPLIANT. |

---

## 🚀 Como Usar

### Exemplo 1: Criar Skill Simples (Quick Mode)

**Input:**
```
Criar uma nova skill chamada "deep-research" para pesquisa avançada na web.
Categoria: research.
Sem sub-skills.
```

**O agente executa:**
1. `skill-factory-bootstrap` com `skill_name=deep-research`, `description="Pesquisa avançada na web com síntese de resultados"`, `category=research`.
2. `skill-factory-validator` para auditar `deep-research/`.
3. Atualiza `README.md` raiz com a nova entrada na tabela.

**Resultado:**
```
deep-research/
├── SKILL.md          # Definição com placeholders para preenchimento
├── README.md         # Documentação com badges e estrutura
└── CHANGELOG.md      # Entry [1.0.0] com data de hoje
```

---

### Exemplo 2: Criar Skill com Sub-skills (Standard Mode)

**Input:**
```
Criar uma nova skill chamada "infra-automation" para automação de infraestrutura.
Categoria: automation.
Sub-skills: provisioner, monitor, rollback.
```

**O agente executa:**
1. `skill-factory-bootstrap` com `sub_skills=["provisioner", "monitor", "rollback"]`.
2. `skill-factory-validator` para auditar `infra-automation/`.
3. Atualiza `README.md` raiz.

**Resultado:**
```
infra-automation/
├── SKILL.md
├── README.md
├── CHANGELOG.md
├── infra-automation-provisioner.skill.md
├── infra-automation-monitor.skill.md
└── infra-automation-rollback.skill.md
```

---

## ✅ Checklist de Conformidade

Toda skill criada pelo framework é validada contra este checklist:

| # | Check | Descrição |
|---|-------|-----------|
| 1 | **Structural** | Todos os arquivos obrigatórios existem |
| 2 | **Frontmatter** | YAML válido com `name`, `version`, `description`, `category` |
| 3 | **Content** | Seções obrigatórias presentes (`Goal`, `Output Structure`, `Quality Rules`, `Prohibited`) |
| 4 | **Naming** | Diretório, arquivos e frontmatter seguem padrão de nomenclatura |
| 5 | **Registry** | Skill registrada na tabela do `README.md` raiz |

---

## 📐 Versão Policy

- Toda nova skill inicia na versão **`1.0.0`**.
- O changelog segue o formato [Keep a Changelog](https://keepachangelog.com/pt-BR/).
- O versionamento segue [Semantic Versioning](https://semver.org/lang/pt-BR/).

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.

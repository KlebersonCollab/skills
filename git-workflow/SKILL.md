---
name: git-workflow
version: 1.0.0
description: "Padrões de fluxo de trabalho Git, incluindo estratégias de ramificação, convenções de commit e integração com SDD."
category: development-workflow
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
0. **Mode Check**: Verificar o modo operacional atual (`.hub-mode`) e aplicar as diretrizes da skill `token-distiller`.
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---

## Goal

Estabelecer padrões de controle de versão, estratégias de ramificação (branching) e desenvolvimento colaborativo que garantam rastreabilidade, facilitem revisões de código e mantenham a integridade do histórico do projeto, especialmente quando operando sob o framework SDD.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Clean History** | Git Log | Histórico de commits linear e legível. |
| **Pull Request** | Markdown | PR documentado seguindo o template do projeto. |
| **Git Config** | `.gitmessage` | Configuração local de template de commit. |

---

## Estratégias de Ramificação (Branching)

### 1. GitHub Flow (Simples e Recomendado)
Ideal para deploy contínuo e equipes de pequeno a médio porte.
- `main` está sempre pronta para produção.
- Crie branches `feature/` a partir da `main`.
- Abra Pull Request para revisão.
- Após aprovação e testes, faça o merge para `main`.

### 2. Trunk-Based Development (Alta Velocidade)
Para equipes com CI/CD robusto e Feature Flags.
- Commits curtos e frequentes diretamente na `main` (ou branches de vida curtíssima).
- CI deve passar antes de qualquer integração.

### 3. GitFlow (Ciclo de Release Enterprise)
Para releases agendados e projetos complexos.
- `main` (produção), `develop` (integração).
- Branches de `feature`, `release` e `hotfix`.

---

## Conventional Commits

Formato obrigatório: `<type>(<scope>): <subject>`

| Tipo | Uso | Exemplo |
|------|-----|---------|
| `feat` | Nova funcionalidade | `feat(auth): add SSO support` |
| `fix` | Correção de bug | `fix(api): handle connection timeout` |
| `docs` | Documentação | `docs(readme): add setup steps` |
| `refactor` | Refatoração de código | `refactor(db): optimize query logic` |
| `test` | Testes | `test(unit): add auth validation tests` |

*Veja o guia completo em: [references/conventional-commits.md](references/conventional-commits.md)*

---

## Git no Workflow SDD (Mandatório)

A integração do Git com o ciclo SDD é fundamental para a rastreabilidade:

1.  **Fase Discovery**: Antes de iniciar, garanta que sua `main` local está sincronizada com o remoto (`git pull origin main`).
2.  **Fase Implementation**: 
    - **Branches Atômicas**: Cada feature/fix deve ter sua própria branch seguindo o padrão `feature/<id>-<name>` ou `fix/<id>-<name>`.
    - **Commits por Task**: Cada tarefa concluída em `tasks.md` deve gerar pelo menos um commit claro.
    - **Linguagem**: Todas as mensagens de commit **DEVEM** ser escritas em **Inglês**, mesmo que a documentação técnica esteja em Português.
    - **Atomicidade**: Não misture alterações de tarefas diferentes no mesmo commit.
3.  **Fase Review**:
    - **Linearidade**: Prefira `rebase` para manter o histórico limpo antes do merge.
    - **Evidence**: O link do Pull Request deve ser incluído no `validation-report.md`.

### Diagrama de Ciclo de Commit SDD
```mermaid
graph TD
    A[Início da Tarefa no tasks.md] --> B[Implementação do Código]
    B --> C[Execução de Testes Locais]
    C --> D{Passou?}
    D -- Não --> B
    D -- Sim --> E[Commit: type(scope): description #ID]
    E --> F[Atualizar tasks.md como Concluído]
    F --> G{Mais tarefas?}
    G -- Sim --> B
    G -- Não --> H[Abrir Pull Request]
```

---

## Mandatos de Integração: Merge vs Rebase

### Use REBASE para:
- Atualizar sua branch local com as mudanças mais recentes da `main`.
- Limpar commits intermediários (squash) antes de enviar para revisão.
- Manter um histórico linear e sem "noise" de merge commits desnecessários.

### Use MERGE para:
- Integrar uma branch de feature concluída na `main`.
- Quando a branch é compartilhada com outros desenvolvedores (evitar force push em branches públicas).

---

## Quality Rules

- **English-Only**: Todas as mensagens de commit devem ser escritas em Inglês.
- **Conventional-Commits**: Uso obrigatório do padrão Conventional Commits.
- **Linear-History**: Preferência por `rebase` para manter o histórico limpo.
- **Atomic-Commits**: Commits pequenos e focados em uma única mudança.

---

## Prohibited

- **NUNCA** realizar commits diretos na `main` sem revisão (exceto modo Quick SDD).
- **NUNCA** commitar segredos, chaves de API ou arquivos `.env`.
- **NUNCA** usar mensagens de commit vagas (ex: "update", "fix").
- **NUNCA** enviar Pull Requests gigantes (>500 linhas).
- **NUNCA** escrever mensagens de commit em qualquer idioma que não seja Inglês.

---

## Referências e Exemplos
- [Guia de Conventional Commits](references/conventional-commits.md)
- [Estratégias Detalhadas de Branching](references/branching-strategies.md)
- [Template de Mensagem de Git](examples/gitmessage.template)
- [Template de Pull Request](examples/pull-request.md)

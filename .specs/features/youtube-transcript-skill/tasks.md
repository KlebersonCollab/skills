# Tasks: YouTube Transcript Skill

Este arquivo rastreia o progresso da implementação da skill `youtube-transcript`.

---

## Phase 1: Bootstrap & Scaffolding
- [x] **Task 1: Scaffolding Inicial**
  - [x] Executar `skill-factory-bootstrap` para criar o diretório `youtube-transcript/` e arquivos base.
  - [x] Validar se a estrutura de diretórios foi criada corretamente.

## Phase 2: Core Development
- [x] **Task 2: Definição da Skill (SKILL.md)**
  - [x] Implementar a seção `Instructions` com o fluxo de CLI e fallback para Whisper.
  - [x] Integrar o script Python inline para limpeza e deduplicação de VTT.
  - [x] Configurar a gestão de dependências via `uv`.

- [x] **Task 3: Documentação e Metadados**
  - [x] Escrever o `README.md` detalhando o uso e as dependências externas.
  - [x] Atualizar o `CHANGELOG.md` com a versão inicial `1.0.0`.

## Phase 3: Validation & Quality
- [x] **Task 4: Auditoria de Conformidade**
  - [x] Executar `skill-factory-validator` na skill criada.
  - [x] Corrigir eventuais falhas apontadas pelo validador.

- [x] **Task 5: Testes Funcionais**
  - [x] Validar a extração de legendas manuais (Lógica validada).
  - [x] Validar a extração de legendas automáticas e o script de limpeza (Lógica validada com mock).
  - [x] Validar o fallback para Whisper (IA) via uv run protocol.

## Phase 4: Finalization (Review & Persistence)
- [x] **Task 6: Registro Global**
  - [x] Adicionar a nova skill na tabela de Skills Disponíveis do `README.md` raiz.
  - [x] Incrementar o contador de skills no badge do `README.md` raiz.

- [x] **Task 7: Relatórios e Estado**
  - [x] Gerar `validation-report.md` e `verification-report.md`.
  - [x] Atualizar `STATE.md`, `ROADMAP.md` e marcar tarefas como concluídas.

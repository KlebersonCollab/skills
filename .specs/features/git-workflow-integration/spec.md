# Spec: Git Workflow Skill Integration

## Goal
Integrar a skill `git-workflow` ao Hub de Skills, traduzindo seu conteúdo para Português Brasileiro, alinhando-a ao padrão `skill-factory` e incorporando mandatos específicos de SDD e fluxo de trabalho do projeto.

## Requisitos Funcionais
- **FR1: Tradução Completa**: Todo o conteúdo da skill original deve ser traduzido para Português Brasileiro (PT-BR).
- **FR2: Conformidade Skill Factory**: O arquivo `SKILL.md` deve conter o frontmatter correto e o hook `🔒 Prerequisites (Mandatory)`.
- **FR3: Integração SDD**: Adicionar seções que expliquem como o Git Workflow se conecta com as fases do SDD (ex: commits atômicos por tarefa).
- **FR4: Mandatos de Utilização**: Definir regras claras de quando usar `merge` vs `rebase` conforme a cultura do Hub.
- **FR5: Registro no Hub**: A skill deve ser validada e registrada no `README.md` raiz.
- **FR6: Idioma de Commit**: Especificar que todas as mensagens de commit devem ser escritas em Inglês.

## Critérios de Aceitação (BDD)

### Cenário 1: Estrutura da Skill
- **Given** que o diretório `git-workflow/` foi criado.
- **When** eu verifico o arquivo `SKILL.md`.
- **Then** ele deve conter o frontmatter com `name: git-workflow`, `version: 1.0.0` e a descrição em PT-BR.
- **And** deve conter a seção `🔒 Prerequisites (Mandatory)`.

### Cenário 2: Qualidade do Conteúdo
- **Given** o conteúdo original da ECC.
- **When** eu realizo a tradução e adaptação.
- **Then** as explicações de Branching Strategies e Conventional Commits devem estar claras em PT-BR.
- **And** deve haver uma seção específica sobre "Git no Ciclo SDD".

### Cenário 3: Registro e Validação
- **Given** que a skill foi implementada.
- **When** eu executo o `validator` (simulado ou manual conforme disponibilidade).
- **Then** a skill deve ser adicionada à tabela de skills no `README.md` raiz.

## Restrições
- Não usar placeholders.
- Seguir rigorosamente a regra de responder em Português Brasileiro (exceto exemplos de commit, que devem ser em Inglês).

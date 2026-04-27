# Specification: HB-CLI Advanced Stats & Cost Tracking (hb stats)

## Context
O projeto possui um binário nativo Go (`hb`) que gerencia a governança do Hub. Atualmente, o rastreamento de custos e uso de tokens é feito de forma ad-hoc ou manual. Inspirado na funcionalidade madura do projeto em `src/`, este recurso visa trazer visibilidade determinística sobre o consumo de recursos e performance das sessões.

## Functional Requirements
- **FR-1: Rastreamento de Tokens**: O `hb` deve ser capaz de registrar tokens de input, output e cache (hit/miss).
- **FR-2: Cálculo de Custos**: O `hb` deve calcular o custo estimado em USD baseado no modelo utilizado.
- **FR-3: Métricas de Código**: O `hb` deve rastrear linhas adicionadas e removidas durante a sessão (via integração com git).
- **FR-4: Relatório de Sessão**: O comando `hb stats` deve exibir um resumo visual e formatado do uso atual.
- **FR-5: Persistência de Sessão**: Os dados devem ser persistidos em um arquivo JSON durante a sessão atômica.
- **FR-6: Integração com Handoff**: Os dados de custo devem ser incluídos no processo de handoff para garantir continuidade.

## Acceptance Criteria (AC)
### AC-1: Comando `hb stats` básico
- **Given** que uma sessão foi iniciada com `hb session start`.
- **When** eu executo `hb stats`.
- **Then** o sistema deve exibir:
    - Total de tokens (Input/Output/Cache).
    - Custo estimado em USD.
    - Duração da sessão (Wall clock).
    - Linhas de código alteradas.

### AC-2: Persistência Robusta
- **Given** uma sessão ativa com dados acumulados.
- **When** eu encerro a sessão com `hb session stop`.
- **Then** os dados finais devem ser exibidos no sumário de encerramento.
- **And** o arquivo temporário de estatísticas deve ser arquivado ou limpo conforme a política.

### AC-3: Diferenciação de Modelos
- **Given** o uso de múltiplos modelos (ex: Gemini Pro, Flash).
- **When** eu consulto as estatísticas.
- **Then** o custo deve ser calculado separadamente por modelo antes de ser totalizado.

## Constraints
- Escrito em **Go 1.25+**.
- Sem dependências de banco de dados externo (usar arquivos locais em `.specs/project/`).
- Output formatado com cores (usando `fatih/color`).

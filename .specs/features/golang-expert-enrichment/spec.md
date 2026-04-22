# Spec: Golang Expert Skill Enrichment (ECC Patterns)

## 1. Visão Geral
Esta especificação detalha o enriquecimento da skill `golang-expert` com padrões idiomáticos e melhores práticas extraídas da skill externa `golang-patterns` do ECC (Everything Claude Code). O objetivo é elevar a qualidade das entregas em Go para um nível premium, focando em composição, simplicidade e performance.

## 2. Requisitos Funcionais (FR)

- **FR-1**: Integrar o padrão "Accept Interfaces, Return Structs" nas diretrizes de design.
- **FR-2**: Implementar a filosofia "Make the Zero Value Useful" na construção de tipos.
- **FR-3**: Documentar e exemplificar o "Functional Options Pattern" para configurações.
- **FR-4**: Estabelecer a regra de interfaces definidas pelo consumidor (Consumer-Defined Interfaces).
- **FR-5**: Adicionar suporte e exemplos para `errgroup` na orquestração de concorrência.
- **FR-6**: Implementar padrões de Graceful Shutdown para resiliência de sistemas.
- **FR-7**: Proibir explicitamente o uso de `context.Context` dentro de structs.
- **FR-8**: Expandir o checklist de performance com pré-alocação e `strings.Builder`.
- **FR-9**: Criar a skill especializada `golang-testing-expert` baseada no ECC.
- **FR-10**: Implementar padrões de Table-Driven Tests, Mocks, Benchmarks e Fuzzing na nova skill.
- **FR-11**: Integrar as duas skills via referências cruzadas.

## 3. Critérios de Aceitação (AC) - BDD

### Cenário 1: Enriquecimento de Padrões de Design
**Given** que o agente está utilizando a skill `golang-expert`
**When** ele projetar uma nova API ou componente
**Then** ele deve aplicar o padrão "Accept Interfaces, Return Structs" e definir interfaces no lado do consumidor.
**And** os tipos criados devem ser utilizáveis em seu valor zero quando possível.

### Cenário 2: Configuração de Componentes
**Given** a necessidade de configurar uma struct complexa (ex: Servidor, Cliente)
**When** o agente implementar a inicialização
**Then** ele deve utilizar o "Functional Options Pattern" para fornecer flexibilidade e defaults seguros.

### Cenário 3: Concorrência e Resiliência
**Given** um fluxo de processamento paralelo com goroutines
**When** o agente implementar a coordenação
**Then** ele deve preferir `errgroup` para capturar o primeiro erro e cancelar as demais goroutines do grupo.
**And** deve garantir que o sistema responda a sinais de interrupção (SIGINT/SIGTERM) para shutdown gracioso.

### Cenário 4: Prevenção de Anti-padrões
**Given** a manipulação de contextos
**When** o agente definir estruturas de dados
**Then** ele NUNCA deve incluir `context.Context` como campo de struct.

### Cenário 5: Excelência em Testes (TDD)
**Given** a necessidade de criar uma suíte de testes robusta
**When** o agente invocar a skill `golang-testing-expert`
**Then** ele deve seguir o ciclo Red-Green-Refactor.
**And** utilizar Table-Driven Tests com subtests (`t.Run`) e helpers (`t.Helper`, `t.Cleanup`).
**And** implementar Benchmarks e Fuzz tests para componentes críticos.

## 4. Restrições Técnicas
- Todo o conteúdo deve ser traduzido para Português Brasileiro (PT-BR).
- Exemplos de código devem ser concisos e idiomáticos.
- A estrutura de pastas de referências (`references/`) deve ser mantida.

## 5. Auditoria de Qualidade
- A skill deve passar no `validate_skills.py`.
- O score no `validation-report.md` deve ser >= 95.

# Spec: Token Optimization & Dual Mode (Token-Distiller)

## 1. Visão Geral
Implementar um sistema de dois modos de operação para o Hub de Skills, inspirado no conceito **Caveman**. O objetivo é permitir que o agente alterne entre uma comunicação ultra-concisa (Low Token) para tarefas simples e uma comunicação detalhada/analítica (High Token) para tarefas complexas, otimizando o consumo de tokens e a velocidade de resposta sem comprometer a integridade técnica.

## 2. Requisitos Funcionais

### RF-01: Modos de Operação
- **M-01.1 (Low Token / Caveman):**
    - Estilo linguístico: Fragmentos, sem artigos, sem saudações, sinônimos curtos.
    - Foco: Velocidade e execução direta.
    - Uso sugerido: Tarefas "Quick" e "Small", correções simples, atualizações de tarefas.
- **M-01.2 (High Token / Premium):**
    - Estilo linguístico: Gramaticalmente completo, analítico, explicativo.
    - Foco: Precisão, rastreabilidade e documentação exaustiva.
    - Uso sugerido: Tarefas "Medium", "Large" e "Complex", planejamento arquitetural, revisões críticas.

### RF-02: Alternância de Modos (Toggle)
- O sistema deve permitir a troca de modo via comando de barra (Slash Command) ou instrução natural.
- Comandos: `/mode low` (ou `/caveman`) e `/mode high` (ou `/premium`).
- O modo padrão deve ser detectado automaticamente com base no "Task Sizing" do SDD, mas pode ser sobrescrito pelo usuário.

### RF-03: Compressão de Contexto (Context Distiller)
- Implementar uma funcionalidade de compressão para arquivos de estado e memória (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`).
- O agente deve ser capaz de "minificar" o conteúdo histórico para economizar tokens de entrada em novas sessões.

### RF-04: Integração com SDD
- O modo "Low Token" deve ser compatível com as fases de implementação do SDD.
- O modo "High Token" deve ser o padrão para as fases de especificação e planejamento (Discovery/Specify).

## 3. Critérios de Aceitação (ACs) - BDD

### Cenário 1: Ativação do Modo Low Token
**Given** que estou operando no modo padrão (Premium)
**When** eu executo o comando `/mode low`
**Then** o agente deve confirmar a mudança em estilo Caveman (ex: "Low mode active. Token die. Speed go brrr.")
**And** as respostas subsequentes devem seguir as regras de compressão linguística do Caveman.

### Cenário 2: Auto-Sizing de Modo
**Given** que início uma tarefa classificada como "Quick"
**When** o processo de bootstrap da sessão ocorre
**Then** o agente deve sugerir ou ativar automaticamente o modo "Low Token" para economizar recursos.

### Cenário 3: Compressão de Arquivos de Estado
**Given** que o arquivo `STATE.md` ultrapassou 200 linhas
**When** eu solicitar a compressão (`/compress state`)
**Then** o agente deve criar uma versão minificada do arquivo, preservando todos os IDs de tarefas e status técnicos, mas reduzindo a prosa explicativa.

## 4. Restrições Técnicas
- A precisão técnica e IDs de tarefas **NUNCA** devem ser alterados ou omitidos no modo Low Token.
- Os diagramas Mermaid devem permanecer intactos em ambos os modos.
- A segurança e avisos críticos devem ignorar o modo Low Token e usar prosa clara (Auto-Clarity rule).

## 5. Próximos Passos (Planejamento)
- Definir a arquitetura do `token-distiller`.
- Criar a skill de compressão de arquivos.
- Atualizar o `onboarding-navigator` para suportar a seleção de modo no boot.

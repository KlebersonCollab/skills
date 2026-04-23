# Spec: Token Optimization & Dual Mode (Token-Distiller)

## 1. Visão Geral
Implementar sistema dois modos operação para Hub Skills, inspirado no conceito **Caveman**. objetivo é permitir que agente alterne entre comunicação ultra-concisa (Low Token) para tarefas simples e comunicação detalhada/analítica (High Token) para tarefas complexas, otimizando consumo tokens e velocidade resposta sem comprometer integridade técnica.

## 2. Requisitos Funcionais

### RF-01: Modos de Operação
- **M-01.1 (Low Token / Caveman):**
 - Estilo linguístico: Fragmentos, sem artigos, sem saudações, sinônimos curtos.
 - Foco: Velocidade e execução direta.
 - Uso sugerido: Tarefas "Quick" e "Small", correções simples, atualizações tarefas.
- **M-01.2 (High Token / Premium):**
 - Estilo linguístico: Gramaticalmente completo, analítico, explicativo.
 - Foco: Precisão, rastreabilidade e documentação exaustiva.
 - Uso sugerido: Tarefas "Medium", "Large" e "Complex", planejamento arquitetural, revisões críticas.

### RF-02: Alternância de Modos (Toggle)
- sistema deve permitir troca modo via comando barra (Slash Command) ou instrução natural.
- Comandos: `/mode low` (ou `/caveman`) e `/mode high` (ou `/premium`).
- modo padrão deve ser detectado automaticamente com base no "Task Sizing" SDD, mas pode ser sobrescrito pelo usuário.

### RF-03: Compressão de Contexto (Context Distiller)
- Implementar funcionalidade compressão para arquivos estado e memória (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`).
- agente deve ser capaz "minificar" conteúdo histórico para economizar tokens entrada em novas sessões.

### RF-04: Integração com SDD
- modo "Low Token" deve ser compatível com fases implementação SDD.
- modo "High Token" deve ser padrão para fases especificação e planejamento (Discovery/Specify).

## 3. Critérios de Aceitação (ACs) - BDD

### Cenário 1: Ativação do Modo Low Token
**Given** que estou operando no modo padrão (Premium)
**When** eu executo comando `/mode low`
**Then** agente deve confirmar mudança em estilo Caveman (ex: "Low mode active. Token die. Speed go brrr.")
**And** respostas subsequentes devem seguir regras compressão linguística Caveman.

### Cenário 2: Auto-Sizing de Modo
**Given** que início tarefa classificada como "Quick"
**When** processo bootstrap sessão ocorre
**Then** agente deve sugerir ou ativar automaticamente modo "Low Token" para economizar recursos.

### Cenário 3: Compressão de Arquivos de Estado
**Given** que arquivo `STATE.md` ultrapassou 200 linhas
**When** eu solicitar compressão (`/compress state`)
**Then** agente deve criar versão minificada arquivo, preservando todos IDs tarefas e status técnicos, mas reduzindo prosa explicativa.

## 4. Restrições Técnicas
- precisão técnica e IDs tarefas **NUNCA** devem ser alterados ou omitidos no modo Low Token.
- diagramas Mermaid devem permanecer intactos em ambos modos.
- segurança e avisos críticos devem ignorar modo Low Token e usar prosa clara (Auto-Clarity rule).

## 5. Próximos Passos (Planejamento)
- Definir arquitetura `token-distiller`.
- Criar skill compressão arquivos.
- Atualizar `onboarding-navigator` para suportar seleção modo no boot.
# Mapeamento Brownfield (Projetos Existentes)

**Gatilho:** "Mapear codebase", "Analisar código existente", "Documentar arquitetura atual"
**Propósito:** Entender a estrutura existente do projeto antes de adicionar novas features.

## Processo
Antes de iniciar, utilize ferramentas de análise de código (ast-grep, ripgrep) para explorar sistematicamente:
1. Estrutura de diretórios
2. Stack tecnológica a partir dos manifestos de dependência
3. Padrões extraídos de exemplos de código representativos
4. Integrações externas e dependências
5. Pontos de atenção (Tech debt, bugs, riscos de segurança)

## Artefatos Gerados (em `.specs/codebase/`)

### 1. STACK.md
Documenta a stack tecnológica.
- **Tamanho Limite:** ~1.200 palavras
- **Conteúdo:** Linguagem, Framework, Runtime, Banco de Dados, Bibliotecas de Teste.

### 2. ARCHITECTURE.md
Documenta padrões arquiteturais e fluxo de dados.
- **Tamanho Limite:** ~2.400 palavras
- **Conteúdo:** Diagrama estrutural (Mermaid), padrões identificados, fluxo de dados principal.

### 3. CONVENTIONS.md
Documenta estilo de código e convenções de nomenclatura.
- **Tamanho Limite:** ~1.800 palavras
- **Conteúdo:** Nomes de arquivos, funções, tratamento de erros e tipagem baseados em evidências do código.

### 4. STRUCTURE.md
Documenta o layout de diretórios.
- **Tamanho Limite:** ~1.200 palavras
- **Conteúdo:** Árvore de diretórios (máx 3 níveis), organização de módulos.

### 5. TESTING.md
Documenta a infraestrutura de testes.
- **Tamanho Limite:** ~2.400 palavras
- **Conteúdo:** Frameworks de teste, organização, comandos de execução, cobertura e suporte a paralelismo.

### 6. INTEGRATIONS.md
Documenta integrações com serviços externos.
- **Conteúdo:** APIs, Webhooks, Jobs em background e suas formas de autenticação.

### 7. CONCERNS.md
Superfície de avisos acionáveis sobre a codebase.
- **Tamanho Limite:** ~3.000 palavras
- **Conteúdo:** Dívida técnica (Tech Debt), bugs conhecidos, falhas de segurança, gargalos de performance e fragilidades. *Obrigatório sempre incluir arquivo/caminho como evidência.*

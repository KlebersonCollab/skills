# Tasks: Token Optimization & Dual Mode

## Meta 1: Infraestrutura e Skill Core
- [x] **Task 1.1**: Criar diretório `skills/token-distiller/` e inicializar `SKILL.md`. `id: T1.1`
- [x] **Task 1.2**: Definir diretrizes "Caveman" (Low Token) na skill baseada no repositório de referência. `id: T1.2`
- [x] **Task 1.3**: Definir diretrizes "Premium" (High Token) na skill para tarefas analíticas. `id: T1.3`
- [x] **Task 1.4**: Configurar arquivo de persistência de modo `.hub-mode` na raiz. `id: T1.4`

## Meta 2: Motor de Compressão (Python)
- [x] **Task 2.1**: Implementar `scripts/utils/compressor.py` para minificação de Markdown. `id: T2.1`
- [x] **Task 2.2**: Criar suite de testes para o compressor (garantir preservação de técnica). `id: T2.2`
- [x] **Task 2.3**: Implementar CLI wrapper para facilitar o uso: `python scripts/hub.py compress <file>`. `id: T2.3`

## Meta 3: Integração e Automação
- [x] **Task 3.1**: Atualizar `onboarding-navigator` para injetar o prompt correto no bootstrap. `id: T3.1`
- [x] **Task 3.2**: Adicionar comandos de barra `/mode` e `/compress` no gerenciador de comandos. `id: T3.2`
- [x] **Task 3.3**: Automatizar a compressão de `STATE.md` quando o arquivo exceder um limite de tokens. `id: T3.3`

## Meta 4: Validação e Encerramento
- [x] **Task 4.1**: Realizar benchmark de economia de tokens em uma sessão real. `id: T4.1`
- [x] **Task 4.2**: Gerar `validation-report.md` com evidências dos sensores do contrato. `id: T4.2`
- [x] **Task 4.3**: Realizar merge para a branch principal e atualizar documentação do Hub. `id: T4.3`

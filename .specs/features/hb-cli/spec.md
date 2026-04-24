# Spec: HB-CLI - Command: `audit`

## 1. Visão Geral
O `hb audit` é a ferramenta de integridade definitiva para o Hub de Skills. Ele substitui os scripts Python lentos por um binário Go de alta performance que valida a estrutura, metadados e governança de todas as 25+ skills em milissegundos.

## 2. Requisitos Funcionais (RF)
- **RF01 (Skill Scan)**: O comando deve escanear recursivamente o diretório raiz em busca de diretórios que contenham um `SKILL.md`.
- **RF02 (Structure Validation)**: Validar a presença obrigatória de arquivos de governança: `SKILL.md`, `tasks.md`, `MEMORY.md` e `LEARNINGS.md`.
- **RF03 (Metadata Audit)**: Validar o YAML Frontmatter do `SKILL.md` (campos obrigatórios: `name`, `version`, `description`, `category`).
- **RF04 (Version Consistency)**: Garantir que a versão da skill segue o padrão SemVer.
- **RF05 (SDD State Check)**: Verificar se o `tasks.md` da skill está em um estado válido (ex: não estar vazio se houver progresso reportado).
- **RF06 (Prohibited Rules)**: Realizar busca por padrões proibidos (ex: secrets, paths absolutos de outros usuários).

## 3. Requisitos Não-Funcionais (RNF)
- **RNF01 (Performance)**: O scan completo de 25+ skills deve levar menos de 200ms.
- **RNF02 (Portabilidade)**: Deve ser compilado como um binário estático para macOS.
- **RNF03 (Saída Visual)**: Utilizar cores e tabelas claras para reportar erros (padrão CLI moderno).
- **RNF04 (Zero Dependências)**: O binário final não deve exigir Python, Node ou outras runtimes no host.

## 4. Critérios de Aceitação (AC)
- [ ] O comando `hb audit` retorna Exit Code 0 se todas as skills passarem nos testes.
- [ ] O comando retorna Exit Code 1 se houver qualquer violação de integridade.
- [ ] Exibe um sumário final: "X skills audited, Y passed, Z failed".
- [ ] Ignora diretórios técnicos: `.git`, `.venv`, `.specs`, `node_modules`, `artifacts`.

## 5. Distribuição e Portabilidade (RF07)
- **Multi-OS**: O CLI deve ser distribuído para macOS (arm64/amd64), Linux (amd64) e Windows (amd64).
- **Single Binary**: A ferramenta deve ser um binário único autossuficiente.
- **CI/CD Integration**: Os binários devem ser gerados automaticamente via GitHub Actions em cada release.

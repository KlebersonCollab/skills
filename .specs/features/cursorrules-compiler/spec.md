# PRD: Cursorrules Compiler

## 1. Contexto e Visão
O Hub atualmente distribui skills para Claude, Gemini e agentes genéricos via pastas ocultas (`.claude/`, etc). A IDE **Cursor** é hoje uma das ferramentas mais utilizadas e ela lê contexto através de um arquivo `.cursorrules` na raiz do projeto. Esta feature visa criar um compilador que agregue os Mandatos Globais e as Skills em um único arquivo `.cursorrules` altamente otimizado.

## 2. Escopo (In Scope)
- Script Python (`scripts/generate_cursorrules.py`) para ler `GLOBAL_MANDATES.md` e os arquivos `SKILL.md`.
- Formatação XML/Markdown otimizada para o limite de contexto do Cursor.
- Geração automática durante o build (`Makefile` e GitHub Actions).

## 3. Requisitos Técnicos
1. **Compilador**: O script deve extrair as seções vitais (Goal, Workflow, Rules) ignorando changelogs e introduções genéricas para economizar tokens.
2. **Make**: Adicionar comando `make dist-cursor`.
3. **Distribuição**: O arquivo compilado será gerado na pasta `dist_staging/` e zipado (ou apenas disponibilizado cru no Release).

## 4. Próximos Passos (Implementação)
- [ ] Criar o script de compilação.
- [ ] Testar a injeção de contexto em um projeto vazio no Cursor.
- [ ] Integrar no `Makefile` (comando `release`).

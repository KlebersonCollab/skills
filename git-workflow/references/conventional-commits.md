# Guia de Conventional Commits

O padrão **Conventional Commits** é uma convenção leve sobre mensagens de commit. **Mandato do Projeto**: Todas as mensagens devem ser escritas em **Inglês**.

## Estrutura da Mensagem

```
<tipo>(<escopo>): <assunto>

[corpo opcional]

[rodapé opcional]
```

### Tipos (Mandatórios)

| Tipo | Descrição |
|------|-----------|
| **feat** | Uma nova funcionalidade para o usuário, não uma nova funcionalidade para o script de build. |
| **fix** | Uma correção de bug para o usuário, não uma correção para o script de build. |
| **docs** | Mudanças na documentação. |
| **style** | Mudanças que não afetam o significado do código (espaço em branco, formatação, falta de ponto e vírgula, etc). |
| **refactor** | Uma mudança de código que nem corrige um bug nem adiciona uma funcionalidade. |
| **perf** | Uma mudança de código que melhora o desempenho. |
| **test** | Adicionando testes ausentes ou corrigindo testes existentes. |
| **build** | Mudanças que afetam o sistema de build ou dependências externas (ex: gulp, broccoli, npm). |
| **ci** | Mudanças em nossos arquivos e scripts de configuração de CI (ex: Travis, Circle, BrowserStack, SauceLabs). |
| **chore** | Outras mudanças que não modificam os arquivos `src` ou de teste. |
| **revert** | Reverte um commit anterior. |

### Escopo (Opcional)
O escopo pode ser qualquer coisa que especifique o local da mudança do commit. Por exemplo: `auth`, `api`, `ui`, `database`, etc.

### Assunto (Mandatório)
O assunto contém uma descrição sucinta da mudança:
- Use o imperativo, tempo presente: "change" não "changed" nem "changes".
- Não capitalize a primeira letra.
- Sem ponto final (.) no final.

## Exemplos de Mensagens

### Commit Simples
```bash
feat(api): add oauth2 support for external partners
```

### Commit com Corpo e Rodapé
```bash
fix(ui): resolve race condition in login form

Added a loading state to prevent multiple submissions while the request
is in flight.

Closes #456
```

### Breaking Change
```bash
feat(db):! migrate to postgresql

BREAKING CHANGE: The database layer now requires a PostgreSQL connection string.
```

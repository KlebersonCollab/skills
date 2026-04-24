# Contract: HB-CLI Validation

## 1. Sensores de Validação
Para garantir que o `hb audit` é confiável, utilizaremos os seguintes sensores:

### S1: Autovalidação (Dogfooding)
O comando `hb audit` deve ser capaz de auditar a si mesmo e ao repositório de skills atual e passar com sucesso.

### S2: Testes de Regressão (Mock Skills)
Criar uma pasta temporária `test_fixtures/` com:
- Uma skill válida.
- Uma skill sem `SKILL.md` (deve falhar).
- Uma skill com YAML inválido (deve falhar).
- Uma skill com versão "1.0" em vez de "1.0.0" (deve falhar).

### S3: Benchmark de Performance
O tempo de execução deve ser medido.
- **Meta**: < 200ms para o repositório completo.

## 2. Matriz de Aceitação

| Entrada | Comportamento Esperado | Status Code |
|---------|------------------------|-------------|
| Repo íntegro | Lista skills e dá OK | 0 |
| Skill s/ SKILL.md | Erro: Missing SKILL.md | 1 |
| Frontmatter corrompido | Erro: Invalid YAML | 1 |
| Pasta ignorada (.git) | Não deve processar | 0 |

## 3. Definição de Pronto (DoD)
- [ ] Código passa em `go fmt` e `go vet`.
- [ ] Binário compilado com sucesso para a arquitetura local.
- [ ] Todos os casos de `test_fixtures` passam.
- [ ] Relatório visual segue o padrão premium do Hub.

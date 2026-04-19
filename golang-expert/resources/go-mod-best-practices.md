# Melhores Práticas: Go Modules

Guia prático para gerenciamento de dependências e versionamento em projetos Go.

## 1. Comandos Essenciais

- `go mod init <module-name>`: Inicializa um novo módulo. O nome deve ser a URL do repositório (ex: `github.com/user/repo`).
- `go mod tidy`: Remove dependências não utilizadas e adiciona as necessárias. **Execute sempre antes de commitar.**
- `go mod vendor`: Cria uma cópia local de todas as dependências em uma pasta `vendor/`. Útil para builds offline ou ambientes restritos.
- `go list -m all`: Lista todas as dependências diretas e indiretas.

## 2. Versionamento e Semantic Versioning (SemVer)

Go Modules seguem o SemVer (`vX.Y.Z`):
- **Patch (Z)**: Correções de bugs.
- **Minor (Y)**: Novas funcionalidades retrocompatíveis.
- **Major (X)**: Mudanças que quebram a compatibilidade.

**Atenção**: Versões Major >= 2 devem incluir o sufixo no nome do módulo:
```go
module github.com/user/repo/v2
```

## 3. O arquivo `go.sum`
NUNCA edite o `go.sum` manualmente. Ele contém os hashes criptográficos das dependências para garantir que o código baixado hoje seja idêntico ao baixado amanhã (Segurança e Reprodutibilidade).

## 4. Substituições Locais (Replace)
Para testar mudanças em uma dependência localmente antes de publicá-la:
```bash
go mod edit -replace github.com/dependencia/projeto=../caminho/local
```
**Lembre-se**: Remova o `replace` antes de fazer o push para o repositório principal.

## 5. Estratégia de Commits
- Commit o `go.mod` e o `go.sum`.
- O uso da pasta `vendor/` é opcional, mas se utilizada, deve ser commitada integralmente.

## Dicas Pro
1. **Mantenha dependências atualizadas**: Use `go get -u ./...` para atualizar dependências (com cautela).
2. **Evite Dependências Desnecessárias**: Go preza pela simplicidade. Às vezes, 10 linhas de código próprio são melhores que uma dependência externa pesada.
3. **Internal Packages**: Use a pasta `internal/` para pacotes que não devem ser exportados para outros módulos.

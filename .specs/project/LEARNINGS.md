# Project Incremental Wisdom (Learnings)

## Automação & Scripting
- **Zipping Hidden Folders**: Ao criar ZIPs de pastas ocultas (como `.claude`), é fundamental entrar no diretório pai (`cd dist_staging`) antes de executar o comando `zip` para garantir que o ZIP contenha a pasta como sua raiz, facilitando o uso pelo usuário final.
- **Dynamic Skill Discovery**: O uso de `find . -maxdepth 2 -name "SKILL.md"` é uma estratégia robusta para listar skills sem precisar de um arquivo de manifesto manual, tornando o projeto auto-expansível.

## CI/CD (GitHub Actions)
- A separação de artefatos por nome no `upload-artifact` v4 melhora muito a experiência do usuário.
- **GitHub Releases vs Artifacts**: Para distribuição de software final (como os ZIPs de skills), Releases são preferíveis por serem permanentes e permitirem links estáticos (`/releases/latest/download/`), o que simplifica a integração e documentação.
- **Tag-Driven Releases**: Automatizar a criação de releases a partir do push de tags (`v*`) garante que apenas versões validadas sejam distribuídas aos usuários.

# IDE Integration: Flutter with FVM

Para que sua IDE reconheça corretamente o SDK do Flutter gerenciado pelo FVM, siga estes passos.

## 1. VS Code (Recomendado)

Crie ou edite o arquivo `.vscode/settings.json` na raiz do projeto:

```json
{
  "dart.flutterSdkPath": ".fvm/flutter_sdk",
  "search.exclude": {
    "**/.fvm": true
  },
  "files.exclude": {
    "**/.fvm": true
  }
}
```

**Dica**: Adicione `.vscode/` ao seu `.gitignore` se você quiser manter as configurações pessoais fora do repositório, ou versione se quiser garantir que todos usem o SDK local.

## 2. Android Studio / IntelliJ

Para estas IDEs, você deve apontar o caminho do SDK do Flutter manualmente:

1. Abra o projeto no Android Studio.
2. Vá em `Preferences` (macOS) ou `Settings` (Windows).
3. Navegue até `Languages & Frameworks > Flutter`.
4. No campo `Flutter SDK path`, informe o caminho completo até a pasta `.fvm/flutter_sdk` do seu projeto.
   - Exemplo: `/Users/meu_user/Projetos/meu_app/.fvm/flutter_sdk`
5. Clique em `Apply`.

## 3. Terminal da IDE

Sempre que abrir o terminal integrado da sua IDE, lembre-se de usar o prefixo `fvm`.

Para facilitar, você pode criar um alias no seu `.zshrc` ou `.bashrc`:

```bash
alias fl="fvm flutter"
```

Agora você pode usar `fl run` em vez de `fvm flutter run`.

## 4. Troubleshooting no VS Code

Se o VS Code continuar avisando que o SDK não foi encontrado:

1. Execute `fvm use stable` no terminal para garantir que o link `.fvm/flutter_sdk` existe.
2. No VS Code, pressione `Cmd+Shift+P` (macOS) ou `Ctrl+Shift+P` (Windows).
3. Digite `Flutter: Change SDK`.
4. Selecione a opção que aponta para o caminho do FVM no seu projeto.

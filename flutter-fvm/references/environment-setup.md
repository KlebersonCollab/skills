# Environment Setup: Flutter with FVM

Este guia detalha como configurar o ambiente de desenvolvimento Flutter utilizando o FVM.

## 1. Instalação do FVM

O FVM pode ser instalado de várias formas:

### macOS / Linux
```bash
brew tap leoafarias/fvm
brew install fvm
```

### Windows
```powershell
choco install fvm
# Ou via Pub
dart pub global activate fvm
```

## 2. Configurando o PATH

Certifique-se de que o diretório de binários do Dart/Pub está no seu PATH se você instalou via Pub:

```bash
# macOS/Linux (zsh)
export PATH="$PATH":"$HOME/.pub-cache/bin"
```

## 3. Gerenciando Versões do SDK

### Instalar uma versão específica
```bash
fvm install 3.19.0
```

### Instalar a versão stable
```bash
fvm install stable
```

### Listar versões instaladas
```bash
fvm list
```

## 4. Fixando a Versão no Projeto

No diretório raiz do seu projeto Flutter, execute:

```bash
fvm use stable
```

Isso criará uma pasta `.fvm` com o arquivo `fvm_config.json` e um link simbólico `flutter_sdk` apontando para a versão selecionada.

## 5. Verificação de Saúde (FVM Doctor)

Execute o comando abaixo para verificar se tudo está configurado corretamente para o FVM:

```bash
fvm doctor
```

Para verificar o estado do Flutter via FVM:

```bash
fvm flutter doctor
```

## 6. .gitignore Recomendado

Adicione o seguinte ao seu `.gitignore` para não versionar o cache local do SDK:

```gitignore
.fvm/flutter_sdk
```

Mas **não** adicione o `fvm_config.json`, ele deve ser versionado para que o time use a mesma versão.

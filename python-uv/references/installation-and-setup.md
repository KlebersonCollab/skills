# Installation & Setup

Guia completo de instalação do UV e configuração de ambientes Python em todas as plataformas.

---

## Instalação

### macOS

```bash
# Homebrew (Recomendado)
brew install uv

# Standalone installer
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### Linux

```bash
# Standalone installer (Recomendado)
curl -LsSf https://astral.sh/uv/install.sh | sh

# Via pipx (se já tiver Python)
pipx install uv
```

### Windows

```powershell
# PowerShell (Recomendado)
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"

# Via pip (se já tiver Python)
pip install uv
```

### Verificação

```bash
uv --version
uvx --version
# Deve exibir algo como: uv 0.9.7
```

---

## Virtual Environment Setup

### Método 1: UV Project (Recomendado para novos projetos)

```bash
# UV cria e gerencia venv automaticamente
uv init meu-projeto
cd meu-projeto
uv sync  # cria .venv automaticamente
uv run python -c "print('Hello UV!')"
```

### Método 2: Manual com venv

```bash
# Criar venv manualmente
python -m venv .venv

# Ativar (macOS / Linux)
source .venv/bin/activate

# Ativar (Windows Git Bash)
. .venv/Scripts/activate

# Ativar (Windows CMD)
.venv\Scripts\activate.bat

# Ativar (Windows PowerShell)
.venv\Scripts\Activate.ps1

# Instalar pacotes com UV dentro do venv
uv pip install requests numpy pandas
```

### Método 3: UV sem ativar (Recomendado para automação)

```bash
# Executar diretamente sem ativar venv
uv run python script.py
uv run pytest
uv run ruff check .
```

---

## Configuração de PATH

Se UV não for encontrado após a instalação:

### macOS / Linux

```bash
# Adicionar ao ~/.zshrc ou ~/.bashrc
export PATH="$HOME/.local/bin:$PATH"

# Recarregar configuração
source ~/.zshrc  # ou source ~/.bashrc
```

### Windows

```powershell
# Verificar se está no PATH
$env:PATH -split ';' | Select-String "uv"

# Adicionar manualmente se necessário
$env:PATH += ";$env:LOCALAPPDATA\Programs\uv"
```

---

## Troubleshooting

### UV não encontrado após instalação

```bash
# Verificar localização
which uv           # macOS/Linux
where.exe uv       # Windows

# Se não encontrar, reinstalar
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### Venv não ativa

```bash
# Verificar se o script de ativação existe
ls .venv/bin/activate        # macOS/Linux
ls .venv/Scripts/activate    # Windows

# Se não existir, recriar
rm -rf .venv
uv venv .venv
# ou
python -m venv .venv
```

### Permissões negadas

```bash
# macOS/Linux
chmod +x ~/.local/bin/uv
chmod +x ~/.local/bin/uvx
chmod -R u+w ~/.cache/uv/
```

---

## Quick Reference Card

```bash
# Instalação
curl -LsSf https://astral.sh/uv/install.sh | sh    # Linux/Mac
brew install uv                                      # macOS Homebrew

# Novo projeto
uv init meu-projeto && cd meu-projeto

# Dependências
uv add requests pydantic
uv add --dev pytest ruff mypy
uv sync

# Executar
uv run python main.py
uv run pytest

# Ferramentas globais
uv tool install ruff
uv tool install mypy

# Versão Python
uv python install 3.12
uv python pin 3.12
```

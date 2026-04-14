# Inline Script Metadata (PEP 723)

Guia completo para scripts Python self-contained com dependências definidas em comentários, eliminando a necessidade de `pyproject.toml` para scripts únicos.

---

## O que é

PEP 723 permite definir dependências diretamente em comentários do script Python. UV lê estes metadados e instala automaticamente as dependências ao executar.

### Sintaxe

```python
# /// script
# dependencies = [
#   "pacote-a>=1.0.0",
#   "pacote-b",
# ]
# requires-python = ">=3.11"
# ///
```

**Regras de sintaxe:**
- `# /// script` — marcador de abertura (exatamente assim)
- `# dependencies = [...]` — lista de dependências em formato TOML
- `# requires-python = "..."` — versão Python (opcional)
- `# ///` — marcador de fechamento (exatamente assim)

---

## Exemplos Práticos

### Web Scraper

```python
# /// script
# dependencies = [
#   "requests>=2.28.0",
#   "beautifulsoup4>=4.12.0",
# ]
# ///

import requests
from bs4 import BeautifulSoup

def scrape(url: str) -> str:
    response = requests.get(url)
    soup = BeautifulSoup(response.content, "html.parser")
    return soup.title.string if soup.title else "No title"

if __name__ == "__main__":
    title = scrape("https://example.com")
    print(f"Page title: {title}")
```

```bash
uv run scraper.py
```

### API Server (FastAPI)

```python
# /// script
# requires-python = ">=3.11"
# dependencies = [
#   "fastapi>=0.100.0",
#   "uvicorn[standard]>=0.24.0",
# ]
# ///

from fastapi import FastAPI
import uvicorn

app = FastAPI(title="Quick API")

@app.get("/")
def root() -> dict[str, str]:
    return {"message": "Hello from inline script!"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
```

```bash
uv run server.py
# Acesse http://localhost:8000
```

### CLI Tool com Rich

```python
# /// script
# dependencies = [
#   "typer>=0.9.0",
#   "rich>=13.0.0",
# ]
# ///

import typer
from rich.console import Console
from pathlib import Path

app = typer.Typer()
console = Console()

@app.command()
def process(
    input_dir: Path = typer.Argument(..., help="Input directory"),
    pattern: str = typer.Option("*.txt", help="File pattern"),
) -> None:
    """Process files matching pattern."""
    files = list(input_dir.glob(pattern))
    console.print(f"[green]Found {len(files)} files[/green]")
    for f in files:
        console.print(f"  → {f.name}")

if __name__ == "__main__":
    app()
```

```bash
uv run processor.py ./data --pattern "*.csv"
```

### Data Analysis com Polars

```python
# /// script
# requires-python = ">=3.10"
# dependencies = [
#   "polars>=0.19.0",
#   "matplotlib>=3.8.0",
# ]
# ///

import polars as pl
import matplotlib.pyplot as plt
import sys

def analyze(file_path: str) -> None:
    df = pl.read_csv(file_path)
    print(df.describe())

    # Salvar gráfico
    data_pd = df.to_pandas()
    data_pd.hist(figsize=(10, 6))
    plt.tight_layout()
    plt.savefig("analysis.png")
    print("Saved analysis.png")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: uv run analyze.py <file.csv>")
        sys.exit(1)
    analyze(sys.argv[1])
```

---

## Quando Usar vs Não Usar

### ✅ Use inline metadata para:
- Scripts de arquivo único
- Utilitários e automação rápida
- Exemplos compartilháveis que "just work"
- Scripts de análise de dados
- CLIs simples
- Prototipação rápida

### ❌ NÃO use para:
- Projetos multi-arquivo (use `pyproject.toml`)
- Aplicações com muitas dependências (>10-15 pacotes)
- Projetos com separação dev/test deps
- Aplicações de produção
- Pacotes para publicação no PyPI

---

## Boas Práticas

### Versionar dependências para reprodução

```python
# ✅ Bom — reproduzível
# dependencies = [
#   "requests==2.31.0",
#   "pandas==2.1.0",
# ]

# ✅ Bom — flexível mas seguro
# dependencies = [
#   "requests>=2.28.0,<3.0.0",
#   "pandas>=2.0.0",
# ]

# ⚠️ Arriscado — comportamento pode mudar
# dependencies = [
#   "requests",
#   "pandas",
# ]
```

### Organizar dependências

```python
# /// script
# dependencies = [
#   # Web
#   "requests>=2.28.0",
#   "beautifulsoup4>=4.12.0",
#   # Data
#   "pandas>=2.0.0",
#   "numpy>=1.24.0",
# ]
# ///
```

---

## Troubleshooting

### "Failed to parse inline script metadata"

Verificar sintaxe exata dos marcadores:

```python
# /// script        ← EXATAMENTE assim (com espaço)
# dependencies = [
#   "pacote",       ← Formato TOML válido
# ]
# ///               ← EXATAMENTE assim
```

### "Package not found"

```bash
# Verificar nome no PyPI
curl https://pypi.org/pypi/nome-do-pacote/json | head -5
```

### "Requires Python >=3.11 but found 3.10"

```bash
# Instalar versão necessária
uv python install 3.11

# Executar com versão específica
uv run --python 3.11 script.py
```

---

## Performance

| Execução | Tempo |
|----------|-------|
| Primeira vez (download + setup) | 5-30s (depende das deps) |
| Execuções subsequentes (cache) | <1s |
| Após `uv cache clean` | Volta ao tempo da primeira |

# Padrões Idiomáticos e Performance em Python

Este guia define os padrões de excelência para escrita de código Python robusto, eficiente e mantível, integrado ao workflow de alta performance do **UV**.

---

## 1. Princípios Fundamentais (The Zen of Python)

### Legibilidade em Primeiro Lugar
O código deve ser óbvio e fácil de entender. Se o código é difícil de explicar, é uma má ideia.

```python
# ✅ GOOD: Claro e legível
def obter_usuarios_ativos(usuarios: list[User]) -> list[User]:
    """Retorna apenas usuários ativos da lista fornecida."""
    return [u for u in usuarios if u.esta_ativo]

# ❌ BAD: "Clever" mas confuso
def get_act_u(u):
    return [x for x in u if x.a]
```

### Explícito é melhor que Implícito
Evite "mágica" e efeitos colaterais ocultos. Seja claro sobre o que o seu código faz.

```python
# ✅ GOOD: Configuração explícita
import logging
logging.basicConfig(level=logging.INFO)

# ❌ BAD: Efeitos colaterais escondidos
import modulo_magico # Configura logging internamente sem avisar
```

---

## 2. EAFP vs LBYL

Python prefere o estilo **EAFP** (*Easier to Ask Forgiveness than Permission*) em vez de **LBYL** (*Look Before You Leap*).

```python
# ✅ GOOD (EAFP): Tente e trate a exceção
try:
    valor = dicionario["chave"]
except KeyError:
    valor = padrao

# ❌ BAD (LBYL): Verifique antes de agir
if "chave" in dicionario:
    valor = dicionario["chave"]
else:
    valor = padrao
```

---

## 3. Modern Type Hints (Python 3.9+)

Utilize tipagem estática para melhorar a manutenção e permitir que ferramentas como `mypy` (via `uv run mypy`) detectem erros precocemente.

```python
# Python 3.9+ - Use tipos built-in diretamente
def processar_itens(itens: list[str]) -> dict[str, int]:
    return {item: len(item) for item in itens}

# Generic Types e Union (3.10+)
def primeiro_elemento[T](lista: list[T]) -> T | None:
    return lista[0] if lista else None
```

### Protocol-Based Duck Typing (Interfaces)
Use `Protocol` para definir contratos sem herança rígida.

```python
from typing import Protocol

class Renderizavel(Protocol):
    def render(self) -> str: ...

def exibir_tudo(itens: list[Renderizavel]) -> None:
    for item in itens:
        print(item.render())
```

---

## 4. Gestão de Exceções Avançada

### Exception Chaining
Sempre use `from e` ao relançar exceções para preservar o traceback original.

```python
def processar_dados(dados: str):
    try:
        resultado = json.loads(dados)
    except json.JSONDecodeError as e:
        raise ValueError("Dados JSON inválidos") from e
```

### Hierarquia de Exceções Customizada
Crie uma base para sua aplicação para facilitar o tratamento global.

```python
class AppError(Exception): """Base para todos os erros da app"""
class ValidationError(AppError): """Erro de validação de input"""
```

---

## 5. Performance e Eficiência de Memória

### Otimização com `__slots__`
Para classes com milhões de instâncias, use `__slots__` para reduzir drasticamente o uso de memória (evita a criação do `__dict__`).

```python
class Ponto:
    __slots__ = ['x', 'y']
    def __init__(self, x: float, y: float):
        self.x = x
        self.y = y
```

### Geradores para Grandes Volumes de Dados
Nunca retorne listas gigantescas se puder usar geradores (Lazy Evaluation).

```python
# ✅ GOOD: Processa linha por linha
def ler_arquivo_gigante(caminho: str):
    with open(caminho) as f:
        for linha in f:
            yield linha.strip()

# ❌ BAD: Carrega tudo na RAM
def ler_tudo(caminho: str):
    with open(caminho) as f:
        return f.readlines()
```

### Manipulação de Strings
Evite concatenação com `+` em loops (O(n²)). Use `.join()` ou `io.StringIO` (O(n)).

```python
# ✅ GOOD
resultado = "".join(lista_de_strings)
```

---

## 6. Data Classes com Validação

Use `dataclasses` para contêineres de dados, e `__post_init__` para validações simples. Para validações complexas, considere `Pydantic` (integrado via `uv add pydantic`).

```python
from dataclasses import dataclass

@dataclass
class Usuario:
    email: str
    idade: int

    def __post_init__(self):
        if "@" not in self.email:
            raise ValueError("Email inválido")
        if self.idade < 0:
            raise ValueError("Idade não pode ser negativa")
```

---

## 7. Anti-Padrões a Evitar

1. **Argumentos Default Mutáveis**: Nunca use `list` ou `dict` como default. Use `None`.
   ```python
   # ❌ BAD
   def add_item(item, lista=[]): ... 
   # ✅ GOOD
   def add_item(item, lista=None):
       if lista is None: lista = []
   ```
2. **Bare Except**: Nunca use `except:`. Capture exceções específicas ou `Exception`.
3. **Verificação de Tipo com `type()`**: Use `isinstance(obj, classe)`.
4. **Comparação com `None`**: Use `is None` em vez de `== None`.

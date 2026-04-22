# Guia de Testes Avançado com Pytest

Este guia estabelece os padrões de qualidade e testes para projetos Python, utilizando **pytest** como ferramenta principal executada via **UV**.

---

## 1. Metodologia TDD (Red-Green-Refactor)

O desenvolvimento deve seguir o ciclo atômico de testes:

1.  🔴 **RED**: Escreva um teste que falha para o comportamento desejado.
2.  🟢 **GREEN**: Escreva o código mínimo necessário para fazer o teste passar.
3.  🔵 **REFACTOR**: Melhore o código mantendo os testes verdes.

---

## 2. Configuração e Cobertura

### Metas de Cobertura
- **Alvo**: 80%+ de cobertura global.
- **Caminhos Críticos**: 100% de cobertura obrigatória.

### Execução via UV
Configure o `pyproject.toml` para incluir relatórios automáticos:

```toml
[tool.pytest.ini_options]
addopts = "--cov=meu_pacote --cov-report=term-missing --cov-report=html"
```

Comando de execução:
```bash
uv run pytest
```

---

## 3. Estrutura de Testes Profissional

Organize seus testes por categoria para facilitar a execução seletiva (usando markers).

```text
tests/
├── conftest.py          # Fixtures compartilhadas
├── unit/                # Testes de lógica pura (rápidos)
├── integration/         # Testes de DB, API, I/O
└── e2e/                 # Fluxos completos de usuário
```

---

## 4. Mastery de Fixtures

Use fixtures para gerenciar estado e recursos de forma limpa.

### Setup e Teardown (Yield)
```python
import pytest

@pytest.fixture
def db_session():
    # Setup
    session = Session.create()
    session.begin_transaction()
    
    yield session # Fornece para o teste
    
    # Teardown
    session.rollback()
    session.close()
```

### Escopos de Fixture
- `function` (padrão): Executa para cada teste.
- `module`: Executa uma vez por arquivo.
- `session`: Executa uma vez por suite completa.

---

## 5. Parametrização

Evite duplicar testes para inputs diferentes. Use `@pytest.mark.parametrize`.

```python
@pytest.mark.parametrize("email, esperado", [
    ("valido@teste.com", True),
    ("invalido", False),
    ("", False),
])
def test_validacao_email(email, esperado):
    assert is_valid_email(email) is esperado
```

---

## 6. Mocking e Isolamento

Utilize `unittest.mock` (ou `pytest-mock`) para isolar dependências externas.

### Mocking de Funções/APIs
```python
from unittest.mock import patch

@patch("meu_pacote.servico_externo.chamar_api")
def test_processamento_com_mock(mock_api):
    mock_api.return_value = {"status": "ok"}
    resultado = processar()
    assert resultado == "sucesso"
    mock_api.assert_called_once()
```

### Mocking de Exceções
```python
def test_erro_na_api(mock_api):
    mock_api.side_effect = ConnectionError("Timeout")
    with pytest.raises(ConnectionError):
        executar_chamada()
```

---

## 7. Testes Assíncronos

Para projetos com `asyncio` (FastAPI, etc), use `pytest-asyncio`.

```python
@pytest.mark.asyncio
async def test_endpoint_async(cliente_async):
    res = await cliente_async.get("/status")
    assert res.status_code == 200
```

---

## 8. Boas Práticas (DOs and DON'Ts)

✅ **DO:**
- **Nomes descritivos**: `test_login_com_senha_errada_deve_falhar`.
- **Um comportamento por teste**: Mantenha os testes atômicos.
- **Use `tmp_path`**: Fixture nativa do pytest para manipular arquivos temporários com segurança.
- **Mocks com `autospec=True`**: Garante que o mock tenha a mesma interface que o objeto real.

❌ **DON'T:**
- **Não teste implementação**: Foque no comportamento (input/output).
- **Não compartilhe estado**: Testes devem ser independentes e poder rodar em qualquer ordem.
- **Não ignore falhas**: Um teste falhando é um bloqueio imediato.
- **Não use `print`**: Use logs ou o output do pytest em caso de falha.

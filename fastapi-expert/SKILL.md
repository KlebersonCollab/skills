---
name: fastapi-expert
version: 2.2.0
description: "Expert level FastAPI implementation focused on performance, clean architecture, and Python-UV integration."
category: development-python
---

# FastAPI Expert: API Excellence

> "Fast, easy to learn, fast to code, ready for production." - This skill ensures your FastAPI code actually lives up to that promise.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
0. **Mode Check**: Verificar o modo operacional atual (`.hub-mode`) e aplicar as diretrizes da skill `token-distiller`.
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
## Goal

Prover um framework de decisão e implementação para APIs de alto desempenho, garantindo o uso correto de Injeção de Dependências, Tipagem Estrita (Annotated) e integração nativa com o ecossistema `python-uv`.

---

## Workflow (6 Phases)

### Phase 0: ARCHITECTURE_DESIGN
Definição da estrutura do projeto (Clean Architecture).
- **Rule**: Decidir entre Repository Pattern ou Service Layer antes de codar.
- **Reference**: Ver [Architecture Guide](references/architecture.md).

### Phase 1: SCHEMA_DESIGN
Definição dos contratos de dados usando Pydantic V2.
- **Rule**: NUNCA usar `RootModel`. Preferir `TypeAdapter` ou modelos estruturados.
- **Mandate**: Todo campo deve ter `Field(description=...)` para documentação automática.

### Phase 2: DEPENDENCY_ARCHITECTURE
Mapeamento de dependências e segurança (JWT).
- **Rule**: Usar exclusivamente `Annotated[Type, Depends(func)]`.
- **Logic**: Criar aliases para dependências reutilizáveis.

### Phase 3: IMPLEMENTATION & ERROR_HANDLING
Escrita dos endpoints e handlers de exceção.
- **Rule**: Usar `async def` apenas para operações verdadeiramente I/O async.
- **Mandate**: Implementar handlers globais para exceções de domínio.

### Phase 4: VALIDATION & PERF
Auditoria de performance e documentação.
- **Check**: Validar se não há código bloqueante dentro de `async def`.
- **Mandate**: Perform performance audits using standardized benchmarking tools and document results.

### Phase 5: TESTING
Implementação de testes unitários e de integração.
- **Rule**: Garantir 100% de cobertura nos serviços críticos com Pytest.

---

## Key Patterns (The FastAPI Standard)

### 1. Annotated Style (Mandatory)
```python
from typing import Annotated
from fastapi import Depends, Path

# DO THIS
async def read_item(item_id: Annotated[int, Path(title="The ID of the item")]):
    ...

# NOT THIS
async def read_item(item_id: int = Path(...)):
    ...
```

### 2. UV Integration
Sempre inicializar e gerenciar dependências via `python-uv`:
```bash
uv add fastapi pydantic-settings
uv run fastapi dev main.py
```

---

## Quality Rules

- **Clean Code**: Seguir SOLID. Endpoints devem ser "magros" (logic in services/dependencies).
- **Security**: Utilizar `OAuth2PasswordBearer` e escopos para controle de acesso.
- **Performance**: Pydantic V2 (Rust core) deve ser a única fonte de serialização.

---

## Prohibited

- NUNCA usar `...` (Ellipsis) em modelos Pydantic ou parâmetros obrigatórios.
- NUNCA misturar lógica de banco de dados diretamente no endpoint (usar Dependências).
- NUNCA usar bibliotecas de serialização depreciadas (ujson/orjson).

## Output Structure

A execução desta skill deve resultar em APIs que seguem a seguinte estrutura de arquivos recomendada:

| Artefato | Localização | Descrição |
|----------|-------------|-----------|
| **Entrypoint** | `src/main.py` | Inicialização do FastAPI e inclusão de routers. |
| **Schemas** | `src/schemas/` | Modelos Pydantic V2 para request/response. |
| **Endpoints** | `src/api/` | Routers organizados por domínio (tags). |
| **Dependencies** | `src/dependencies.py` | Fábrica de dependências injetáveis. |
| **Config** | `src/config.py` | Gestão de ambiente via `pydantic-settings`. |
| **Migrations** | `migrations/` | Scripts de evolução do banco de dados (Alembic). |

## Detailed References

Para um mergulho profundo em cada tópico, consulte os guias especializados:

| Guia | Tópico |
|------|--------|
| [Architecture](references/architecture.md) | Repositories, Services e Lifespan. |
| [API Design](references/api_design.md) | Padrões REST, Status Codes e Naming. |
| [Security](references/security.md) | JWT, OAuth2 e Password Hashing. |
| [Error Handling](references/error_handling.md) | Exceções de Domínio e Handlers Globais. |
| [Performance](references/performance.md) | Async vs Sync e Pydantic Rust core. |
| [Testing](references/testing.md) | Pytest, AsyncClient e Integration Tests. |
| [Migrations](references/migrations.md) | Alembic Assíncrono e versionamento de banco. |


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.359394Z"
evidence_checksum: "NONE"
```

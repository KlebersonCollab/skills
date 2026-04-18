# FastAPI Error Handling Standard

Um tratamento de erros consistente é vital para a experiência do desenvolvedor (frontend) e para a segurança da API.

## 1. Custom Exceptions

Defina exceções de domínio para separar erros de lógica de erros de infraestrutura.

```python
class DomainException(Exception):
    def __init__(self, message: str, status_code: int = 400):
        self.message = message
        self.status_code = status_code

class UserNotFound(DomainException):
    def __init__(self):
        super().__init__("Usuário não encontrado", status_code=404)
```

## 2. Global Exception Handlers

Use `@app.exception_handler` para capturar exceções e retornar um formato JSON padronizado.

```python
from fastapi import Request
from fastapi.responses import JSONResponse

@app.exception_handler(DomainException)
async def domain_exception_handler(request: Request, exc: DomainException):
    return JSONResponse(
        status_code=exc.status_code,
        content={
            "error": exc.__class__.__name__,
            "message": exc.message,
            "path": request.url.path
        },
    )
```

## 3. Validation Errors (Override)

Personalize o erro 422 (Pydantic) para torná-lo mais amigável.

```python
from fastapi.exceptions import RequestValidationError

@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request, exc):
    return JSONResponse(
        status_code=422,
        content={"detail": exc.errors(), "message": "Dados inválidos"},
    )
```

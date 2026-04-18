# FastAPI Security Reference

## 1. Dependency Injection for Security

O FastAPI permite injetar segurança diretamente na assinatura da função.

```python
from typing import Annotated
from fastapi import Depends, Security
from fastapi.security import OAuth2PasswordBearer

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")

# Reutilizável em qualquer endpoint
TokenDep = Annotated[str, Depends(oauth2_scheme)]

@app.get("/items/")
async def read_items(token: TokenDep):
    return {"token": token}
```

## 2. Scopes e Autorização

Use `Security` em vez de `Depends` quando precisar de escopos:

```python
async def get_current_user(
    security_scopes: SecurityScopes, 
    token: Annotated[str, Depends(oauth2_scheme)]
):
    # Lógica de validação de escopo aqui
    ...
```

## 3. JWT (JSON Web Tokens) Implementation

Utilize `python-jose` para manipulação de tokens.

```python
from datetime import datetime, timedelta
from jose import jwt

SECRET_KEY = "sua_chave_secreta"
ALGORITHM = "HS256"

def create_access_token(data: dict, expires_delta: timedelta | None = None):
    to_encode = data.copy()
    expire = datetime.utcnow() + (expires_delta or timedelta(minutes=15))
    to_encode.update({"exp": expire})
    return jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
```

## 4. Password Hashing (Passlib)

NUNCA armazene senhas em texto plano. Use `passlib[bcrypt]`.

```python
from passlib.context import CryptContext

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

def hash_password(password: str):
    return pwd_context.hash(password)

def verify_password(plain_password, hashed_password):
    return pwd_context.verify(plain_password, hashed_password)
```

## 5. Best Practices
- **Token Expiration**: Sempre defina um tempo de expiração curto (ex: 15-30 min).
- **Refresh Tokens**: Para sessões longas, implemente Refresh Tokens armazenados de forma segura.
- **Pydantic-Settings**: Nunca coloque chaves de API no código. Use `BaseSettings`.
- **HTTPS Redirect**: Em produção, sempre use `HTTPSRedirectMiddleware`.
- **CORS**: Configure o `CORSMiddleware` de forma restrita, nunca use `allow_origins=["*"]` em produção.

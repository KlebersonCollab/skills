# Exemplo de Especificação REST API (v1)

Este exemplo descreve uma API de gerenciamento de recursos genérica.

## Base URL
`https://api.projeto.com/v1`

---

## Recursos: Usuários (/users)

### 1. Listar Usuários
* **Endpoint**: `/users`
* **Método**: `GET`
* **Descrição**: Retorna uma lista paginada de usuários.
* **Query Params**:
    * `page` (int, default: 1)
    * `limit` (int, default: 10)
* **Status Codes**:
    * `200 OK`: Sucesso.
    * `401 Unauthorized`: Token ausente ou inválido.

---

### 2. Criar Novo Usuário
* **Endpoint**: `/users`
* **Método**: `POST`
* **Body (JSON)**:
  ```json
  {
    "name": "João Silva",
    "email": "joao.silva@email.com",
    "role": "editor"
  }
  ```
* **Status Codes**:
    * `201 Created`: Recurso criado com sucesso.
    * `400 Bad Request`: Erro de validação nos campos.
    * `409 Conflict`: E-mail já cadastrado.

---

## Padronização de Erros (Error Schema)

Todas as respostas de erro devem seguir o formato abaixo para facilitar o consumo por clientes.

```json
{
  "error": {
    "code": "VAL-001",
    "message": "Mensagem descritiva para o desenvolvedor",
    "details": [
      {
        "field": "email",
        "issue": "E-mail inválido"
      }
    ],
    "trace_id": "abc-123-xyz"
  }
}
```

### Códigos de Erro Comuns
| Código | Descrição | Status HTTP |
| :--- | :--- | :--- |
| `AUTH-001` | Token expirado | 401 |
| `FORB-001` | Permissão insuficiente | 403 |
| `NOTF-001` | Recurso não encontrado | 404 |
| `SRVR-001` | Erro interno inesperado | 500 |

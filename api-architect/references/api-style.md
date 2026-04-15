# API Style Guide: Padrões de Interface e Consistência

Este guia define os padrões obrigatórios para todas as APIs projetadas pela skill **API Architect**, garantindo previsibilidade e facilidade de consumo.

---

## 1. Envelopes de Resposta (JSON)

### Sucesso
Toda resposta de sucesso deve retornar os dados dentro de um objeto `data`.

```json
{
  "data": {
    "id": "uuid-123",
    "name": "Kleberson Romero",
    "email": "dev@example.com"
  }
}
```

### Listas e Paginação
Use o padrão de metadados para coleções.

```json
{
  "data": [ ... ],
  "meta": {
    "total": 100,
    "page": 1,
    "limit": 10,
    "has_next": true
  }
}
```

---

## 2. Tratamento de Erros

Toda resposta de erro (4xx, 5xx) deve seguir o formato abaixo para facilitar o tratamento no frontend/clientes.

```json
{
  "error": {
    "code": "ENTITY_NOT_FOUND",
    "message": "O usuário com ID 123 não foi encontrado.",
    "details": {
      "resource": "User",
      "id": "123"
    }
  }
}
```

- **code**: String em `SCREAMING_SNAKE_CASE` (único para cada tipo de erro).
- **message**: Mensagem amigável para humanos.
- **details**: (Opcional) Objeto com detalhes técnicos ou campos de validação.

---

## 3. Nomenclatura e Tipos

- **Campos**: Use `snake_case` para chaves JSON (consistência com a maioria das bibliotecas de persistência).
- **IDs**: Sempre use strings (UUIDs, ULIDs ou HashIDs). Nunca exponha IDs sequenciais de banco de dados.
- **Datas**: Sempre use o formato ISO 8601 (`YYYY-MM-DDTHH:mm:ssZ`).
- **Booleanos**: Devem começar com verbos de estado (ex: `is_active`, `has_permission`).

---

## 4. Verbos HTTP e Status Codes

| Verbo | Ação | Status Sucesso | Status Erro Comum |
|-------|------|----------------|-------------------|
| **GET** | Recuperar recurso | 200 OK | 404 Not Found |
| **POST** | Criar recurso | 201 Created | 400 Bad Request |
| **PUT** | Substituir recurso | 200 OK | 400 Bad Request |
| **PATCH** | Atualizar parcialmente | 200 OK | 400 Bad Request |
| **DELETE** | Remover recurso | 204 No Content | 404 Not Found |

---

## 5. Versionamento

- **Via URL**: Prioritário para APIs públicas (ex: `/v1/users`).
- **Via Headers**: Aceitável para APIs internas ou versões menores (ex: `X-API-Version: 1.2`).

---

## 6. Paginação (Cursor vs Offset)

- **Offset**: Simples, para pequenas coleções (ex: `?page=1&limit=10`).
- **Cursor**: Recomendado para feeds ou grandes coleções (ex: `?after=cursor_id`).

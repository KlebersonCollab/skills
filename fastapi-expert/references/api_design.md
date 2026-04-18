# REST API Design Guide

Padrões de design para garantir que a API seja intuitiva e escalável.

## 1. Nomenclatura de Recursos
- Use **Substantivos no Plural**: `/users`, `/items`, `/orders`.
- NUNCA use verbos no caminho: `/getUsers` (Errado) -> `/users` (Certo).

## 2. Versionamento
- Sempre use versionamento no caminho: `/v1/users`, `/v2/items`.
- Isso permite evoluir a API sem quebrar clientes antigos.

## 3. Status Codes Mandatórios
| Code | Significado | Quando usar |
|------|-------------|-------------|
| 200 | OK | Sucesso em GET ou PUT. |
| 201 | Created | Sucesso em POST (Criação). |
| 204 | No Content | Sucesso em DELETE (Sem corpo). |
| 400 | Bad Request | Erro de lógica do cliente. |
| 401 | Unauthorized | Falta de autenticação. |
| 403 | Forbidden | Autenticado, mas sem permissão (Scopes). |
| 404 | Not Found | Recurso inexistente. |

## 4. Filtragem e Paginação
Use Query Parameters para busca e paginação:
- `GET /users?limit=10&offset=20`
- `GET /items?status=active`

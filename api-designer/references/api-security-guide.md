# API Security Guide — OWASP API Top 10

Guia de proteção e auditoria de segurança para APIs.

---

## 1. O que é o OWASP API Top 10?

O OWASP (Open Web Application Security Project) mantém uma lista das 10 vulnerabilidades mais críticas em APIs.

---

## 2. Principais Riscos e Como Prevenir

### 2.1 API1: Broken Object Level Authorization (BOLA / IDOR)
O atacante manipula o ID de um recurso na URL para acessar dados de outro usuário.
- **Exemplo**: `GET /api/orders/123` → Muda para `GET /api/orders/124`.
- **Prevenção**: Sempre valide se o usuário logado possui permissão explícita para o ID solicitado no banco de dados. **Nunca** confie apenas no ID da URL.

### 2.2 API2: Broken Authentication
Implementação fraca de autenticação (JWT fraco, tokens que não expiram).
- **Prevenção**: Use algoritmos fortes (RS256), defina expiração curta para tokens, use HTTPS sempre e implemente rotação de segredos.

### 2.3 API3: Broken Object Property Level Authorization (Mass Assignment)
O cliente envia campos que não deveria poder alterar.
- **Exemplo**: Enviar `"is_admin": true` no POST de criação de usuário.
- **Prevenção**: Use DTOs (Data Transfer Objects) ou schemas de entrada rigorosos que listam apenas os campos permitidos (`allowlist`).

### 2.4 API4: Unrestricted Resource Consumption (Rate Limit)
Falta de limites de CPU, memória ou taxa de requisições.
- **Prevenção**: Implemente Rate Limiting, limite o tamanho do payload (request body) e use paginação obrigatória.

### 2.5 API5: Broken Function Level Authorization
Usuário comum acessando endpoints administrativos.
- **Exemplo**: Usuário acessando `/api/admin/delete-all`.
- **Prevenção**: Implemente RBAC (Role-Based Access Control) ou ABAC e valide em cada rota.

---

## 3. Segurança de Dados

- **Sensitive Data Exposure**: Nunca retorne senhas, tokens ou dados pessoais (PII) desnecessários na resposta da API.
- **Error Handling**: Evite expor stack traces ou detalhes do servidor em erros. Use mensagens genéricas para o usuário e logs detalhados internos.

---

## 4. Checklist de Segurança para Design de API

- [ ] HTTPS é obrigatório em todos os ambientes?
- [ ] O sistema valida permissão de posse (Ownership) para cada ID de recurso?
- [ ] Existem limites de Rate Limit configurados?
- [ ] O schema de entrada proíbe campos extras (No Mass Assignment)?
- [ ] Tokens JWT usam algoritmos seguros e expiração curta?
- [ ] Dados sensíveis estão sendo filtrados da resposta?
- [ ] A API usa cabeçalhos de segurança (CORS, HSTS, X-Content-Type-Options)?

---

## 5. Ferramentas de Auditoria

- **[OWASP ZAP](https://www.zaproxy.org/)**: Scanner de segurança web.
- **[Postman API Security](https://www.postman.com/api-security/)**: Testes de segurança integrados ao Postman.
- **[42Crunch](https://42crunch.com/)**: Auditoria de especificações OpenAPI.

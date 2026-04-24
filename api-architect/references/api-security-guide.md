# API Security Guide — OWASP API Top 10

Security protection and audit guide for APIs.

---

## 1. What is the OWASP API Top 10?

The OWASP (Open Web Application Security Project) maintains a list of the 10 most critical vulnerabilities in APIs.

---

## 2. Main Risks and How to Prevent Them

### 2.1 API1: Broken Object Level Authorization (BOLA / IDOR)
The attacker manipulates a resource ID in the URL to access another user's data.
- **Example**: `GET /api/orders/123` → Changes to `GET /api/orders/124`.
- **Prevention**: Always validate if the logged-in user has explicit permission for the requested ID in the database. **Never** trust only the URL ID.

### 2.2 API2: Broken Authentication
Weak authentication implementation (weak JWT, non-expiring tokens).
- **Prevention**: Use strong algorithms (RS256), set short expiration for tokens, always use HTTPS, and implement secret rotation.

### 2.3 API3: Broken Object Property Level Authorization (Mass Assignment)
The client sends fields they should not be able to change.
- **Example**: Sending `"is_admin": true` in the user creation POST.
- **Prevention**: Use DTOs (Data Transfer Objects) or rigorous input schemas that list only allowed fields (`allowlist`).

### 2.4 API4: Unrestricted Resource Consumption (Rate Limit)
Lack of limits on CPU, memory, or request rate.
- **Prevention**: Implement Rate Limiting, limit payload size (request body), and use mandatory pagination.

### 2.5 API5: Broken Function Level Authorization
Common user accessing administrative endpoints.
- **Example**: User accessing `/api/admin/delete-all`.
- **Prevention**: Implement RBAC (Role-Based Access Control) or ABAC and validate on every route.

---

## 3. Data Security

- **Sensitive Data Exposure**: Never return passwords, tokens, or unnecessary personal data (PII) in the API response.
- **Error Handling**: Avoid exposing stack traces or server details in errors. Use generic messages for the user and detailed internal logs.

---

## 4. Security Checklist for API Design

- [ ] Is HTTPS mandatory in all environments?
- [ ] Does the system validate ownership for each resource ID?
- [ ] Are Rate Limit limits configured?
- [ ] Does the input schema prohibit extra fields (No Mass Assignment)?
- [ ] Do JWT tokens use secure algorithms and short expiration?
- [ ] Are sensitive data being filtered from the response?
- [ ] Does the API use security headers (CORS, HSTS, X-Content-Type-Options)?

---

## 5. Audit Tools

- **[OWASP ZAP](https://www.zaproxy.org/)**: Web security scanner.
- **[Postman API Security](https://www.postman.com/api-security/)**: Integrated security tests in Postman.
- **[42Crunch](https://42crunch.com/)**: Audit of OpenAPI specifications.

# Technical Plan: CI Pipeline

## 🏗️ Arquitetura
Pipeline linear executado em containers Ubuntu.

```mermaid
graph LR
    Push[Git Push] --> Setup[Setup Environment]
    Setup --> Dependencies[Install Dependencies]
    Dependencies --> Validate[Make Validate]
    Validate --> Verify[Make Verify-Vers]
    Verify --> Done[Success]
```

## 🛠️ Design de Implementação

### Configuração do Workflow (`.github/workflows/validate.yml`):
- **OS**: `ubuntu-latest`.
- **Python**: `3.12`.
- **Ferramentas**: `astral-sh/setup-uv`.

### Passos:
1.  **Checkout**: Clonar o repositório.
2.  **Setup UV**: Instalar o gerenciador `uv`.
3.  **Run Validate**: `make validate` (usa `uv run`).
4.  **Run Verify Versions**: `make verify-vers`.

## 🏁 Milestones
1.  [ ] Criar diretório `.github/workflows`.
2.  [ ] Escrever arquivo YAML do workflow.
3.  [ ] Simular execução local (se possível) ou testar via commit.

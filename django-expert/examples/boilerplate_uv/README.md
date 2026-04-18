# Django Expert Boilerplate

Este é um boilerplate profissional seguindo a skill `django-expert`.

## Requisitos
- [UV](https://github.com/astral-sh/uv)

## Como rodar

1. **Instale as dependências:**
   ```bash
   uv sync
   ```

2. **Configure o ambiente:**
   Crie um arquivo `.env` baseado no `.env.example`.

3. **Rode as migrações:**
   ```bash
   uv run python manage.py migrate
   ```

4. **Inicie o servidor:**
   ```bash
   uv run python manage.py runserver
   ```

## Funcionalidades Inclusas
- **HTMX:** Reatividade server-side sem JS pesado.
- **UV:** Gerenciamento de pacotes ultra-rápido.
- **Celery Ready:** Configurado para tarefas em background.
- **Security:** Configurações via variáveis de ambiente.

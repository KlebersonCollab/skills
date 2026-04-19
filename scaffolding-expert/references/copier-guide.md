# Guia de Referência: Copier (uvx copier)

O Copier é a ferramenta padrão do Hub para geração dinâmica de projetos e arquivos (Scaffolding).

## 🚀 Comandos Básicos

### Inicializar a partir de um template
```bash
uvx copier copy https://github.com/usuario/template-repo.git ./meu-projeto
```

### Atualizar um projeto já gerado (Update)
```bash
uvx copier update
```

## 🛠️ Melhores Práticas

1. **Local-First**: Para criar novos componentes ou sub-módulos, prefira templates locais em `./artifacts/templates/`.
2. **Respostas Salvas**: O arquivo `copier-answers.yml` é gerado automaticamente. Nunca o delete, pois ele permite atualizações futuras.
3. **Flags Úteis**:
   - `--vcs-ref`: Especificar uma tag ou branch do template.
   - `--data`: Passar variáveis via CLI sem o prompt interativo.
   - `--overwrite`: Sobrescrever arquivos existentes sem perguntar.

## 📁 Estrutura de um Template Copier
```text
template-repo/
├── copier.yml          # Configuração e perguntas
├── README.md           # Documentação do template
└── {{ project_name }}/  # Estrutura de arquivos (Jinja2)
```

## ⚠️ Integração com o Hub
Sempre que usar o `scaffolding-expert`, verifique se o template segue as Global Mandates (especialmente a inclusão de pastas `.gemini/` e `.specs/`).

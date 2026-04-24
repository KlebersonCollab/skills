# Reference Guide: Copier (uvx copier)

Copier is the Hub's standard tool for dynamic project and file generation (Scaffolding).

## 🚀 Basic Commands

### Initialize from a template
```bash
uvx copier copy https://github.com/user/template-repo.git ./my-project
```

### Update a generated project
```bash
uvx copier update
```

## 🛠️ Best Practices

1. **Local-First**: To create new components or sub-modules, prefer local templates in `./artifacts/templates/`.
2. **Saved Answers**: The `copier-answers.yml` file is automatically generated. Never delete it, as it enables future updates.
3. **Useful Flags**:
   - `--vcs-ref`: Specify a tag or branch of the template.
   - `--data`: Pass variables via CLI without the interactive prompt.
   - `--overwrite`: Overwrite existing files without asking.

## 📁 Copier Template Structure
```text
template-repo/
├── copier.yml          # Configuration and questions
├── README.md           # Template documentation
└── {{ project_name }}/  # File structure (Jinja2)
```

## ⚠️ Hub Integration
Whenever using `scaffolding-expert`, verify if the template follows the Global Mandates (especially the inclusion of `.gemini/` and `.specs/` folders).

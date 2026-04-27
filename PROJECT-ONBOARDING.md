# AI Agent Skills Hub - Project Governance

Este documento explica como utilizar a governança deste Hub em **novos projetos**.

## 1. Instalação Rápida (One-Liner)

Para habilitar a governança do HB-CLI em qualquer projeto, execute:

```bash
curl -sSL https://raw.githubusercontent.com/KlebersonCollab/skills/main/scripts/install-hb.sh | bash
```

Isso instalará o `hb` em `~/.local/bin/hb`. Certifique-se de que este diretório está no seu PATH.

## 2. Bootstrapping do Projeto

Após instalar, inicialize a governança no diretório raiz do seu novo projeto:

```bash
hb sdd bootstrap
```

Isso criará a estrutura de Tríade de Memória em `.specs/project/`:
- `STATE.md`: Progresso e foco atual.
- `MEMORY.md`: Contextos e decisões de longo prazo.
- `LEARNINGS.md`: Aprendizados técnicos e padrões.

## 3. Integração com GitHub Actions

Para que seu projeto seja auditado automaticamente seguindo os padrões do Hub, crie o arquivo `.github/workflows/governance.yml` no seu projeto:

```yaml
name: Project Governance Audit

on:
  push:
    branches: [ main ]
  pull_request:

jobs:
  audit:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      # Usa a Action oficial do Skills Hub para instalar o HB-CLI
      - uses: KlebersonCollab/skills/.github/actions/hb-setup@main
      
      - name: Run Integrity Audit
        run: hb audit
        
      - name: Check SDD Compliance
        run: hb sdd audit
```

## 4. Uso de Skills do Hub

Você pode "equipar" seu projeto com inteligência específica do Hub:

```bash
# Instala o suporte a Python com todas as regras de governança
hb install python-uv --remote
```

## 5. Templates Disponíveis

Você pode usar nossos templates base como ponto de partida:
- `templates/go-service`: Estrutura base para serviços Go com governança nativa.

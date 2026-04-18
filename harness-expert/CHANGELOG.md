# Changelog — Harness Expert

Todas as mudanças notáveis desta skill serão documentadas neste arquivo.

O formato segue o padrão [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/), e o versionamento segue [Semantic Versioning](https://semver.org/lang/pt-BR/).

---

## [1.1.0] — 2026-04-18

### Alterado
- **Renomeação Global**: Transição oficial de "Hardness" para **Harness Engineering** para alinhar com padrões da indústria.
- **Enhanced Rehydration**: Adicionado passo de **Bootstrap** para validação do ambiente operacional e instalação automática de dependências.

### Adicionado
- **Harness Principles**: Documentação de Feed Forward, Feedback e Sensores no README.

## [1.0.0] — 2026-04-15

### Adicionado
- **SKILL.md**: Definição principal da skill de Harness Engineering.
- **README.md**: Documentação detalhada e visão geral.
- **harness-expert-sync.skill.md**: Sub-skill para sincronização automática de estado.
- **harness-expert-rehydrate.skill.md**: Sub-skill para injeção de contexto operacional.
- **harness-expert-compress.skill.md**: Sub-skill para compactação de contexto (Context Compressor).
- **Resources**: Implementado `compressor.py` para análise e compactação de tarefas.
- **Exemplos**: Adicionado script de demonstração de sincronização.

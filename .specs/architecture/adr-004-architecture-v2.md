# ADR 004: Architecture V2 - Modular Expert Skills

## Status
Accepted

## Context
O Hub começou como uma coleção de arquivos markdown, mas está evoluindo para um ecossistema de ferramentas automatizadas (scripts, templates, validadores). Precisamos formalizar a transição para a "Arquitetura V2".

## Decision
Transição de "Static Knowledge" para "Active Expert Skills".

## Rationale
- **Encapsulamento**: Cada skill deve ser auto-contida (conhecimento + automação).
- **Modularidade**: Uso de `references/` e `resources/` para evitar arquivos gigantes.
- **Validação Programática**: O hub deve ser capaz de se auto-validar via scripts (ex: `validate_skills.py`).

## Consequences
- Reforço da obrigatoriedade do `Makefile` para todas as operações.
- Skills devem ser atualizadas para incluir exemplos reais e recursos de implementação (playbooks).

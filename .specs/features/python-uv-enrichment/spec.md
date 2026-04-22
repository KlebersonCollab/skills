# Feature Spec: Python Expert Enrichment (python-uv)

## Background
A skill `python-uv` foca atualmente no gerenciamento de ambiente e dependências usando UV. Para elevar o nível do Hub, precisamos integrar padrões de desenvolvimento idiomático (Python Patterns) e técnicas avançadas de teste (Python Testing) baseadas em padrões de mercado (ECC).

## Requisitos de Negócio (ECC Enrichment)
- **R1**: Integrar padrões idiomáticos Python (EAFP, Modern Type Hints 3.9+, Protocol).
- **R2**: Implementar guia completo de Testes (TDD, Fixtures, Mocking, Async Testing).
- **R3**: Adicionar padrões de performance (Generators, `__slots__`, String Optimization).
- **R4**: Padronizar a estrutura de projetos modernos.
- **R5**: Manter a compatibilidade com UV como motor principal.

## Requisitos Técnicos
- **RT1**: Criar arquivos de referência em `references/patterns.md` e `references/testing.md`.
- **RT2**: Atualizar o `SKILL.md` para integrar estes novos conhecimentos.
- **RT3**: Tradução completa para Português Brasileiro (PT-BR).
- **RT4**: Seguir o padrão de documentação do Hub (Concept -> Bad/Good Snippets -> Checklist).

## Critérios de Aceitação (ACs)
- [x] O `SKILL.md` menciona os novos pilares de Patterns e Testing.
- [x] Existem exemplos de "Bad" vs "Good" para padrões idiomáticos.
- [x] O ciclo TDD (Red-Green-Refactor) está documentado.
- [x] Padrões de Mocking e Fixtures com Pytest estão exemplificados.
- [x] A documentação está 100% em PT-BR.
- [x] A estrutura da skill segue o padrão do Hub (References, Examples, etc.).

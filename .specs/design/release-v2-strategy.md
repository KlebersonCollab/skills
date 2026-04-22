# Design Specification: Release V2 "Bunker Strategy"

## Visão Geral
Esta especificação define o processo de liberação (Release) da versão 2.0 (Hub Enterprise Upgrade) utilizando a abordagem "Bunker" (Auditoria Estrita SDD). O objetivo é garantir que 100% da base de código, das documentações e das regras de governança do projeto passem por validação técnica rigorosa antes do empacotamento.

## Fluxo de Execução do `Makefile` (O "Super Exit Gate")

A consolidação se dará pela criação de um novo comando no `Makefile` raiz, provisoriamente chamado de `make release-check`. Este comando agregará a seguinte cadeia de auditorias:

### 1. Security & Linting (DevSecOps)
- Execução do **Ruff** (linter e formatter) em modo de verificação rigorosa para todo o código Python.
- *Ferramentas*: `uv run ruff check .` e `uv run ruff format --check .`

### 2. Unit & Integration Testing
- Execução da recém-criada suíte de testes usando pytest, garantindo que a cobertura técnica (especialmente dos `/scripts/`) não tenha sofrido regressões.
- *Ferramentas*: `uv run pytest tests/`

### 3. Structural Integrity & Governance
- Verificação documental das 21 skills do Hub, confirmando a presença de *Frontmatters* corretos e o *Mandatory Hook* do SDD (`🔒 Prerequisites (Mandatory)`).
- *Ferramentas*: `python scripts/validate_skills.py` (ou script correspondente de validação).

### 4. Knowledge Graph Synchronization
- Regeração do arquivo `KNOWLEDGE-MAP.mermaid` para assegurar que ele reflete com exatidão as alterações estruturais mais recentes do catálogo.
- *Ferramentas*: `python scripts/generate_knowledge_map.py`

## Consolidação do Changelog
- Apenas após a obtenção do status "100% verde" (Zero regressões ou erros de linting), prosseguiremos com a compilação do arquivo `CHANGELOG-HUB.md`.
- As notas de release mesclarão as informações validadas pelas automações com curadoria contextual, delineando claramente o atingimento do status "Enterprise" (Metas 1 a 5 concluídas).

## Próximos Passos (Handoff para Implementação)
1. **Atualizar `Makefile`**: Inserir o target `release-check`.
2. **Executar a Auditoria (`make release-check`)**: Rodar a bateria. Se quebrar, corrigir o débito técnico antes de seguir.
3. **Gerar Notas de Release**: Atualizar a documentação oficial, versionar via Git Tag (e.g. `v2.0.0`) e disparar no ecossistema.

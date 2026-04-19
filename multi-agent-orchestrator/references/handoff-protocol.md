# Protocolo de Handoff entre Agentes

O arquivo `handoff.md` é o contrato de transição entre personas (ex: Arquiteto para Engenheiro). Ele deve garantir que o próximo agente tenha contexto total sem redundância.

## Estrutura Obrigatória

### 1. Resumo do Estado (State Summary)
- **O que foi feito**: Lista sucinta de tarefas concluídas.
- **Bloqueios**: Impedimentos encontrados e como foram resolvidos ou contornados.

### 2. Artefatos Gerados
- Links para novos arquivos, specs (PRD/RFC), ADRs e diagramas Mermaid.

### 3. Missão para a Próxima Persona
- **Objetivo Imediato**: O que deve ser feito a seguir.
- **Critérios de Aceite**: Como o próximo agente saberá que concluiu sua parte.

### 4. Contexto Técnico (Deep Dive)
- Variáveis de ambiente necessárias.
- Decisões de design críticas que não estão em ADRs (decisões "táticas").

## Exemplo de Template

```markdown
# Handoff: [Persona Origem] -> [Persona Destino]

## 📝 Status Atual
- [X] Spec Finalizada em `.specs/features/auth-v2/prd.md`
- [X] ADR-005 aprovado para uso de JWT.

## 🔗 Artefatos
- Spec: `prd.md`
- Mockup: `ui-draft.mermaid`

## 🎯 Próximos Passos (Para [Persona Destino])
1. Implementar o repositório de usuários seguindo a interface `IUserRepository`.
2. Validar contra os testes em `tests/auth_test.py`.

## ⚠️ Observações
- O banco de dados de dev deve estar rodando na porta 5432.
```

# Contract: GAN-style Harness Integration

## 🤝 Acordo de Entrega
Eu, o Agente, me comprometo a implementar os padrões de GAN-style Harness nas skills `harness-expert` e `sdd` sem degradar ou sobrescrever os conhecimentos existentes. A implementação será modular e focada em habilitar loops de feedback de alta qualidade.

## 📡 Sensores de Validação (O que será testado)

1. **Integridade de Conteúdo**: O arquivo `harness-expert/SKILL.md` deve conter tanto a lógica de Sincronia de Estado (original) quanto o novo GAN Loop (adicional).
2. **Coesão SDD**: A skill `sdd` deve referenciar o GAN Loop como uma ferramenta de Review para alta complexidade.
3. **Conformidade de Hook**: Todas as skills afetadas devem manter o hook mandatório de 5 pontos (`Prerequisites`).
4. **Validadores Locais**: O comando `uv run scripts/validate_skills.py` deve retornar **PASS** para todas as skills modificadas.

## 🏁 Critério de Conclusão (Definition of Done)
- Skills atualizadas na raiz.
- Sincronização realizada para as pastas de agentes (`.gemini`, `.claude`, `.agent`).
- Relatório de auditoria atualizado refletindo o novo "Power-up" do Hub.

---
*Contrato firmado via SDD.*

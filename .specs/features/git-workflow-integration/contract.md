# Contract: Git Workflow Skill Integration

## Sensores de Validação

| ID | Sensor | Critério de Sucesso |
|----|--------|---------------------|
| S1 | Frontmatter Check | `git-workflow/SKILL.md` possui frontmatter válido com versão 1.0.0. |
| S2 | Language Check | 100% do conteúdo visível está em Português Brasileiro. |
| S3 | SDD Hook | A seção `🔒 Prerequisites (Mandatory)` está presente no topo da `SKILL.md`. |
| S4 | Registry Check | A skill aparece na tabela do `README.md` raiz. |
| S5 | Reference Integrity | Links para `references/` e `examples/` dentro da skill estão funcionando. |

## Score Rubric (GAN-style)
- **Originalidade (20%)**: Adaptação para a cultura do projeto vs cópia literal.
- **Craft (30%)**: Qualidade do Markdown, diagramas Mermaid e clareza técnica.
- **Funcionalidade (30%)**: Utilidade prática dos workflows sugeridos para o SDD.
- **Design (20%)**: Organização visual e estrutura de arquivos conforme `skill-factory`.

**Nota de Corte**: 8.5/10

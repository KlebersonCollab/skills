# 📝 Tasks: Hub Enterprise Upgrade

**Status**: [Phase 2] Planning Completed

## 🎯 Meta 1: Hub Test Suite (Prioridade 1)
- [x] 1.1 Configurar ambiente `pytest` e fixtures raiz para testes Python. `[Size: Small]`
- [x] 1.2 Implementar testes unitários focados na pasta `/scripts/`. `[Size: Medium]`

## 🎯 Meta 2: Automated Knowledge Distiller
- [x] 2.1 Criar parser focado nos atributos `uses` e `references` das skills. `[Size: Medium]`
- [x] 2.2 Integrar o parser com geração e formatação do mapa Mermaid automaticamente. `[Size: Medium]`

## 🎯 Meta 3: DevSecOps Skill
- [x] 3.1 Executar `skill-factory` para iniciar o scaffold da skill DevSecOps. `[Size: Quick]`
- [x] 3.2 Documentar guias e mandatos focados em SAST/DAST no `SKILL.md`. `[Size: Medium]`

## 🎯 Meta 4: Skill Installer CLI
- [ ] 4.1 Especificar o contrato CLI (ex: `hub install <skill>`). `[Size: Medium]`
- [ ] 4.2 Lógica de download granular ou extração controlada via Git Sparse Checkout/Raw. `[Size: Large]`

## 🎯 Meta 5: Web UI Dashboard
- [x] 5.1 Bootstrapping com framework Streamlit (Python nativo). `[Size: Small]`
- [x] 5.2 Estruturar parser que reuse `scripts/utils.py` para injetar documentação e Mermaid no Streamlit. `[Size: Medium]`

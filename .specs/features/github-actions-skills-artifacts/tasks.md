# Tasks: GitHub Actions Skills Artifacts

Este arquivo rastreia as tarefas atômicas para a implementação da geração de artefatos de skills.

## 1. Preparation
- [x] Criar diretório da feature e arquivos iniciais (spec, plan, tasks) <!-- id: 0 -->
- [x] Validar a lista de skills na raiz do projeto <!-- id: 1 -->

## 2. Scripting & Tooling
- [x] Criar script `scripts/generate-skills-dist.sh` para gerenciar o mapeamento e cópia <!-- id: 2 -->
- [x] Criar script `scripts/verify-dist.sh` para validar a estrutura final do `dist/` <!-- id: 3 -->

## 3. GitHub Actions Integration
- [x] Criar `.github/workflows/generate-skills-artifacts.yml` <!-- id: 4 -->
- [x] Configurar trigger para `push` na `main` e `workflow_dispatch` <!-- id: 5 -->
- [x] Configurar steps para rodar os scripts de geração e upload de artefatos <!-- id: 6 -->

## 4. Verification & Persistence
- [x] Executar o script de geração localmente e validar o conteúdo dos ZIPs <!-- id: 7 -->
- [x] Validar o workflow (se possível simular ou revisar lógica) <!-- id: 8 -->
- [x] Realizar a Fase 4 do SDD (Review & Persistence) <!-- id: 9 -->
- [x] Atualizar ROADMAP.md e PROJECT.md com a nova feature <!-- id: 10 -->

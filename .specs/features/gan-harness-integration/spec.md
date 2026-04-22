# Spec: GAN-style Harness Integration

## 🎯 Visão Geral
Integrar os padrões de **Harness Engineering** inspirados em GAN (Generative Adversarial Networks) ao Hub de Skills. O objetivo é elevar a qualidade das entregas (especialmente Frontend e Mobile) através de um ciclo de feedback entre um "Gerador" e um "Avaliador" implacável.

## 👥 Personas
- **Implementador**: Usa a skill `harness-expert` para gerenciar o loop de feedback.
- **Avaliador**: Agente de QA (via Playwright) que audita a aplicação viva.

## 📋 Critérios de Aceitação (ACs)

### AC 1: Enriquecimento da `harness-expert`
- [ ] Adicionar seção `## 🔄 GAN Feedback Loop` detalhando as fases de Planejamento, Geração e Avaliação.
- [ ] Incluir a **Rubrica de Avaliação** (Design, Originalidade, Craft, Funcionalidade) como padrão mandatório.
- [ ] Adicionar suporte a **Sensores de Qualidade** (Playwright/Screenshot).
- [ ] Garantir que o conhecimento existente sobre Sincronia de Estado não seja removido.

### AC 2: Atualização do Framework `sdd`
- [ ] Atualizar a `Fase 4: REVIEW` para mencionar o uso opcional (mas recomendado para Large/Complex) do loop GAN.
- [ ] Vincular a saída do Reviewer ao `validation-report.md` com scores baseados na rubrica GAN.

### AC 3: Coesão Visual e de Governança
- [ ] Garantir que o novo hook mandatório de 5 pontos (Prerequisites) seja preservado.
- [ ] Manter a consistência do idioma (PT-BR).

## 🚫 Restrições
- **Preservação de Dados**: Não sobrescrever descrições de categorias ou versões sem motivo.
- **Modularidade**: As ideias GAN devem ser tratadas como uma "Estratégia Avançada", não como o único modo de operação, para evitar overhead em tarefas simples (Quick/Small).

---
*Documento gerado sob o framework SDD.*

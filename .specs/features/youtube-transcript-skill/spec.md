# Specification: YouTube Transcript Skill

A skill projetada para automatizar a extração de transcrições de vídeos do YouTube, priorizando legendas existentes e utilizando IA (Whisper) como fallback, com limpeza automática de texto.

---

## 1. Requisitos Funcionais (FR)

- **FR-1 (Detecção de Legendas):** Deve listar legendas disponíveis sem baixar o vídeo.
- **FR-2 (Download de Legendas):** Deve baixar legendas manuais (prioridade) ou automáticas em formato VTT.
- **FR-3 (Fallback para IA):** Deve baixar áudio e transcrever via Whisper se não houver legendas.
- **FR-4 (Processamento de Texto):** Deve converter VTT para TXT puro, removendo tags HTML e deduplicando linhas.
- **FR-5 (Gestão de Dependências):** Deve verificar a presença de `yt-dlp`, `ffmpeg` e configurar ambiente `uv` para `whisper`.

## 2. Critérios de Aceitação (AC - BDD)

### Cenário 1: Extração de Legendas Existentes
**Given** que o usuário fornece uma URL válida do YouTube com legendas manuais ou automáticas.
**When** a skill é executada.
**Then** deve baixar o arquivo VTT correspondente, convertê-lo para TXT limpo e salvar com o título do vídeo.

### Cenário 2: Fallback para Transcrição via IA
**Given** que o vídeo do YouTube NÃO possui legendas disponíveis.
**When** a skill detecta a ausência de legendas.
**Then** deve perguntar ao usuário se deseja realizar a transcrição via Whisper (IA), baixar apenas o áudio e gerar o TXT.

### Cenário 3: Limpeza e Deduplicação
**Given** um arquivo de legenda VTT com tags HTML e linhas repetidas (comum em legendas automáticas).
**When** o processo de limpeza é acionado.
**Then** o arquivo de saída deve conter apenas o texto limpo, sem metadados de tempo ou duplicidades.

## 3. Restrições Técnicas (Constraints)

- **C-1:** Utilizar `uv` para gerenciar as dependências de Python (whisper, openai-whisper).
- **C-2:** Utilizar `yt-dlp` via CLI para extração de metadados e legendas.
- **C-3:** Salvar arquivos temporários em diretório oculto ou temporário do sistema.
- **C-4:** O output final deve ser sempre um arquivo `.txt` UTF-8.

## 4. Definição de "Done" (DoD)

- [ ] Skill registrada via `skill-factory`.
- [ ] Documentação (`README.md`, `SKILL.md`) completa e validada.
- [ ] Testes de validação para os três cenários BDD realizados.
- [ ] Relatório de verificação (`verification-report.md`) gerado.

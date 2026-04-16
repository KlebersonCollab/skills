---
name: youtube-transcript
version: 1.0.0
description: "Skill para automatizar a extração de transcrições de vídeos do YouTube com fallback para IA (Whisper)."
category: automation
---

# YouTube Transcript

> Skill para automatizar a extração de transcrições de vídeos do YouTube com fallback para IA (Whisper).

---

## Goal

O objetivo principal desta skill é fornecer uma forma rápida e robusta de obter o conteúdo textual de vídeos do YouTube, priorizando legendas existentes (manuais ou automáticas) para economizar recursos e tempo, e utilizando modelos de IA (Whisper) apenas quando necessário.

---

## Protocol

### Phase 1: Context & Metadata
1.  **Extract Info:** Executar `yt-dlp --get-title --list-subs "{{url}}"` para obter o título do vídeo e mapear as legendas disponíveis.
2.  **Sanitize Title:** Limpar o título de caracteres especiais para uso como nome de arquivo.

### Phase 2: Selection Strategy
Seguir a ordem de prioridade para obter a transcrição:

1.  **Manual Subtitles:** Se houver legendas manuais (`--write-sub`), baixá-las no formato VTT.
2.  **Automatic Subtitles:** Se não houver manuais, mas houver automáticas (`--write-auto-sub`), baixá-las no formato VTT.
3.  **AI Fallback (Whisper):** Se nenhuma legenda existir:
    - Solicitar confirmação ao usuário para baixar apenas o áudio (`bestaudio`).
    - Executar transcrição via Whisper utilizando `uv run --with openai-whisper whisper audio_file.mp3 --model base`.

### Phase 3: Text Processing (Cleanup)
Independentemente da fonte (VTT ou Whisper), processar o texto final para garantir limpeza absoluta utilizando o script Python inline:

```python
import sys, re

def clean_vtt(content):
    # Remove VTT Header and Timestamps
    lines = content.split('\n')
    cleaned = []
    for line in lines:
        if '-->' in line or line.strip() == 'WEBVTT' or line.strip().isdigit() or 'Kind:' in line or 'Language:' in line:
            continue
        # Remove HTML tags and extra whitespace
        line = re.sub(r'<[^>]+>', '', line).strip()
        if line: cleaned.append(line)
    
    # Advanced Deduplication for rolling auto-subs
    final = []
    for line in cleaned:
        if not final:
            final.append(line)
            continue
        
        # If current line starts with previous line, replace previous with current (more complete)
        if line.startswith(final[-1]):
            final[-1] = line
        else:
            final.append(line)
    return ' '.join(final)

if __name__ == "__main__":
    with open(sys.argv[1], 'r', encoding='utf-8') as f:
        print(clean_vtt(f.read()))
```

Executar via: `uv run python -c "..." input.vtt > output.txt`

### Phase 4: Delivery & Cleanup
1.  Salvar o arquivo final como `{video_title}.txt`.
2.  Remover arquivos temporários (`.vtt`, `.mp3`, `.wav`) criados durante o processo.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos:

- **Transcrição (.txt):** Arquivo de texto puro com o conteúdo do vídeo limpo e formatado.
- **Metadados:** Título do vídeo integrado ao nome do arquivo.

---

## Quality Rules

- **Accuracy First:** Priorizar sempre legendas manuais sobre automáticas ou geradas por IA.
- **Clean Text:** O resultado final não deve conter tags HTML, timestamps VTT ou redundâncias de legendas automáticas.
- **Dependency Check:** Validar `yt-dlp` e `ffmpeg`. Para `whisper`, utilizar `uv run --with`.

## Prohibited

- NUNCA baixar o vídeo completo se apenas o áudio ou legendas forem necessários.
- NUNCA sobrescrever arquivos existentes sem confirmação.
- NUNCA ignorar falhas de download sem reportar o motivo claro.

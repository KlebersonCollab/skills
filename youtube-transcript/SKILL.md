---
name: youtube-transcript
version: 1.0.0
description: "Skill to automate the extraction of transcripts from YouTube videos with AI fallback (Whisper)."
category: automation
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# YouTube Transcript

> Skill to automate the extraction of transcripts from YouTube videos with AI fallback (Whisper).

---

## Goal

The primary goal of this skill is to provide a fast and robust way to obtain the textual content of YouTube videos, prioritizing existing captions (manual or automatic) to save resources and time, and using AI models (Whisper) only when necessary.

---

## Protocol

### Phase 1: Context & Metadata
1.  **Extract Info:** Execute `yt-dlp --get-title --list-subs "{{url}}"` to get the video title and map available captions.
2.  **Sanitize Title:** Clean the title of special characters for use as a filename.

### Phase 2: Selection Strategy
Follow the priority order to obtain the transcript:

1.  **Manual Subtitles:** If manual captions exist (`--write-sub`), download them in VTT format.
2.  **Automatic Subtitles:** If no manual ones exist, but automatic ones do (`--write-auto-sub`), download them in VTT format.
3.  **AI Fallback (Whisper):** If no captions exist:
    - Request confirmation from the user to download only the audio (`bestaudio`).
    - Execute transcription via Whisper using `uv run --with openai-whisper whisper audio_file.mp3 --model base`.

### Phase 3: Text Processing (Cleanup)
Regardless of the source (VTT or Whisper), process the final text to ensure absolute cleanliness using the inline Python script:

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

Execute via: `uv run python -c "..." input.vtt > output.txt`

### Phase 4: Delivery & Cleanup
1.  Save the final file as `{video_title}.txt`.
2.  Remove temporary files (`.vtt`, `.mp3`, `.wav`) created during the process.

---

## Output Structure

Execution of this skill results in the following artifacts:

- **Transcript (.txt):** Plain text file with cleaned and formatted video content.
- **Metadata:** Video title integrated into the filename.

---

## Quality Rules

- **Accuracy First:** Always prioritize manual captions over automatic or AI-generated ones.
- **Clean Text:** The final result must not contain HTML tags, VTT timestamps, or automatic caption redundancies.
- **Dependency Check:** Validate `yt-dlp` and `ffmpeg`. For `whisper`, use `uv run --with`.

## Prohibited

- NEVER download the full video if only audio or captions are needed.
- NEVER overwrite existing files without confirmation.
- NEVER ignore download failures without reporting a clear reason.

# YouTube Transcript

> Skill para automatizar a extração de transcrições de vídeos do YouTube com fallback para IA (Whisper).

[![Versão](https://img.shields.io/badge/Versão-1.0.0-blue)](#changelog)
[![Status](https://img.shields.io/badge/Status-Beta-yellow)](#)

---

## 📖 Visão Geral

A skill **YouTube Transcript** é uma ferramenta de orquestração projetada para obter o conteúdo textual de vídeos do YouTube da forma mais eficiente possível. Ela automatiza o fluxo de decisão entre legendas existentes e transcrição sintética via IA, garantindo sempre um arquivo de saída limpo e pronto para consumo.

### 🚀 Diferenciais
- **Eficiência de Recursos:** Não baixa o vídeo completo; apenas metadados, legendas ou áudio.
- **Limpeza Inteligente:** Algoritmo Python para remover tags HTML e deduplicar palavras repetidas comuns em legendas automáticas.
- **Fallback Local:** Utiliza o modelo OpenAI Whisper localmente (via \`uv\`) quando o vídeo não possui legendas originais.

---

## ⚙️ Pré-requisitos

Para o pleno funcionamento, certifique-se de ter as seguintes ferramentas instaladas:

1.  **[yt-dlp](https://github.com/yt-dlp/yt-dlp):** Para extração de metadados e legendas.
    - Instalação: \`brew install yt-dlp\` ou \`uv add yt-dlp\`
2.  **[ffmpeg](https://ffmpeg.org/):** Necessário para manipulação de áudio e fallback para Whisper.
    - Instalação: \`brew install ffmpeg\`
3.  **[uv](https://github.com/astral-sh/uv):** Gerenciador de dependências Python utilizado para rodar o Whisper sem poluir seu ambiente global.

---

## 🛠️ Como Usar

A skill é invocada passando a URL do vídeo desejado.

### Fluxo Automático
1.  **Mapeamento:** A skill verifica se existem legendas manuais (Ex: PT-BR, EN).
2.  **Download:** Se encontradas, baixa o arquivo VTT.
3.  **Fallback:** Se não houver legendas, a skill solicitará permissão para baixar o áudio e rodar o **Whisper (IA)**.
4.  **Limpeza:** O texto passa por um script de sanitização para gerar o \`.txt\` final.

### Exemplo de Output
\`\`\`text
Nome do Vídeo: "Como fazer café.txt"
Conteúdo:
Olá, hoje vamos aprender a fazer café de uma forma simples e rápida...
\`\`\`

---

## 🧩 Detalhes Técnicos

- **Modelo Whisper:** Por padrão, utiliza o modelo \`base\` para equilíbrio entre velocidade e precisão.
- **Formato Final:** Sempre UTF-8, sem metadados de tempo (timestamps).

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.

# Exemplo: Limpeza de Transcrição para IA

Este documento demonstra a diferença entre uma transcrição bruta (gerada automaticamente) e uma versão limpa, otimizada para processamento por LLMs.

## 1. Transcrição Bruta (Raw)
> "então pessoal ah hoje a gente vai falar sobre é flutter né e como que o fvm ajuda a gente a manter as versões é do sdk certinhas né então tipo você roda fvm install e aí ele instala pra você..."

**Problemas:**
- Vícios de linguagem ("ah", "né", "tipo").
- Falta de pontuação.
- Repetições desnecessárias.

## 2. Transcrição Limpa (Cleaned)
> "Hoje discutiremos o Flutter e como o FVM (Flutter Version Management) auxilia na manutenção das versões do SDK. Ao executar o comando `fvm install`, a ferramenta instala a versão especificada automaticamente."

**Melhorias:**
- Remoção de ruídos.
- Aplicação de norma culta/técnica.
- Pontuação adequada.
- Termos técnicos destacados.

## 💡 Por que limpar?
1. **Redução de Tokens**: Versões limpas são até 30% menores.
2. **Clareza de Intenção**: Facilita para a IA extrair pontos-chave sem se perder em "filler words".
3. **Melhor Summarização**: O resumo gerado será muito mais preciso e profissional.

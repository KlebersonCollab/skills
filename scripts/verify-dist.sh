#!/bin/bash
# Script para verificar a integridade dos artefatos gerados.
set -e

echo "🔍 Verificando estrutura dos artefatos gerados..."

ARTIFACTS=("claude-skills.zip" "gemini-skills.zip" "agent-skills.zip")
AGENTS=(".claude" ".gemini" ".agent")
INSTR_FILES=("CLAUDE.md" "GEMINI.md" "AGENT.md")

if [ ! -d "artifacts" ]; then
    echo "❌ Erro: Diretório 'artifacts' não encontrado!"
    exit 1
fi

for i in "${!ARTIFACTS[@]}"; do
    FILE="artifacts/${ARTIFACTS[$i]}"
    AGENT_DIR="${AGENTS[$i]}"
    INSTR_FILE="${INSTR_FILES[$i]}"
    
    echo "--- Validando $FILE ---"
    
    if [ ! -f "$FILE" ]; then
        echo "❌ Erro: $FILE não encontrado!"
        exit 1
    fi
    
    # Lista arquivos no zip para checar estrutura
    FILES_IN_ZIP=$(unzip -l "$FILE")
    
    # Verifica presença da pasta do agente (ex: .claude/)
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/"; then
        echo "❌ Erro: Diretório $AGENT_DIR/ não encontrado em $FILE!"
        exit 1
    fi
    
    # Verifica presença do arquivo de instruções na raiz do pacote (dentro da pasta do agente)
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/$INSTR_FILE"; then
        echo "❌ Erro: $INSTR_FILE não encontrado em $AGENT_DIR/ dentro de $FILE!"
        exit 1
    fi
    
    # Verifica presença da pasta de skills
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/skills/"; then
        echo "❌ Erro: Pasta 'skills/' não encontrada em $AGENT_DIR/ dentro de $FILE!"
        exit 1
    fi
    
    # Verifica se pelo menos uma skill conhecida está presente (ex: api-architect)
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/skills/api-architect/"; then
        echo "❌ Erro: Skill 'api-architect' não encontrada em $FILE!"
        exit 1
    fi
    
    echo "✅ Estrutura de $FILE validada com sucesso!"
done

echo "🎉 Todos os artefatos foram verificados com sucesso!"

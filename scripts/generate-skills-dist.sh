#!/bin/bash
# Script para gerar artefatos (.zip) de skills para diferentes agentes.
set -e

# Configuração
DIST_BASE="dist_staging"
OUTPUT_DIR="artifacts"

# Limpeza
rm -rf "$DIST_BASE" "$OUTPUT_DIR"
mkdir -p "$DIST_BASE" "$OUTPUT_DIR"

# Configuração dos Agentes (Nome:ArquivoInstrucao:PastaDestino)
AGENTS=(
    "claude:.claude/CLAUDE.md:.claude"
    "gemini:.gemini/GEMINI.md:.gemini"
    "agent:.agent/AGENT.md:.agent"
)

# Identificação dinâmica das skills (pastas na raiz com SKILL.md)
SKILLS=$(find . -maxdepth 2 -name "SKILL.md" -not -path '*/.*' -exec dirname {} \; | sed 's|^\./||')

echo "🚀 Identificando skills para distribuição..."
for SKILL in $SKILLS; do
    echo "  - $SKILL"
done

for AGENT_CONFIG in "${AGENTS[@]}"; do
    IFS=':' read -r NAME INST_FILE TARGET_FOLDER <<< "$AGENT_CONFIG"
    
    echo "📦 Preparando artefato para $NAME..."
    
    STAGING="$DIST_BASE/$TARGET_FOLDER"
    mkdir -p "$STAGING/skills"
    
    # 1. Copiar arquivo de instruções do agente
    if [ -f "$INST_FILE" ]; then
        cp "$INST_FILE" "$STAGING/"
        echo "   ✅ Copiado: $INST_FILE"
    else
        echo "   ⚠️  Aviso: $INST_FILE não encontrado!"
    fi
    
    # 2. Copiar todas as skills identificadas
    for SKILL in $SKILLS; do
        cp -r "$SKILL" "$STAGING/skills/"
    done
    echo "   ✅ Copiadas todas as skills para $STAGING/skills/"
    
    # 3. Gerar ZIP
    # Entramos no dist_staging para que o ZIP contenha a pasta alvo (ex: .claude/)
    (cd "$DIST_BASE" && zip -rq "../$OUTPUT_DIR/$NAME-skills.zip" "$TARGET_FOLDER")
    echo "   🎁 Gerado: $OUTPUT_DIR/$NAME-skills.zip"
done

echo "✅ Todos os artefatos foram gerados em $OUTPUT_DIR/"

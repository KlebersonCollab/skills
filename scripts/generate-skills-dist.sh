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
    (cd "$DIST_BASE" && zip -rq "../$OUTPUT_DIR/$NAME-skills.zip" "$TARGET_FOLDER")
    echo "   🎁 Gerado: $OUTPUT_DIR/$NAME-skills.zip"
done

# --- Lógica Especial para Antigravity ---
echo "📦 Preparando artefato especial: Antigravity..."
ANTIGRAVITY_STAGING="$DIST_BASE/antigravity/.gemini"
mkdir -p "$ANTIGRAVITY_STAGING/skills"
mkdir -p "$ANTIGRAVITY_STAGING/rules"

# 1. Copiar GEMINI.md para a pasta rules
if [ -f ".gemini/GEMINI.md" ]; then
    cp ".gemini/GEMINI.md" "$ANTIGRAVITY_STAGING/rules/"
    echo "   ✅ Copiado: .gemini/GEMINI.md -> rules/"
fi

# 2. Copiar todas as skills
for SKILL in $SKILLS; do
    cp -r "$SKILL" "$ANTIGRAVITY_STAGING/skills/"
done
echo "   ✅ Copiadas todas as skills para $ANTIGRAVITY_STAGING/skills/"

# 3. Gerar Antigravity.zip (partindo da pasta antigravity/)
(cd "$DIST_BASE/antigravity" && zip -rq "../../$OUTPUT_DIR/Antigravity.zip" ".gemini")
echo "   🎁 Gerado: $OUTPUT_DIR/Antigravity.zip"

echo "✅ Todos os artefatos foram gerados em $OUTPUT_DIR/"

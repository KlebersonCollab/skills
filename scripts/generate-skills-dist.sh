#!/bin/bash
# Script to generate skill artifacts (.zip) for different agents.
set -e

# Configuration
DIST_BASE="dist_staging"
OUTPUT_DIR="artifacts"

# Cleanup
rm -rf "$DIST_BASE" "$OUTPUT_DIR"
mkdir -p "$DIST_BASE" "$OUTPUT_DIR"

# Agent Configuration (Name:InstructionFile:TargetFolder)
AGENTS=(
    "claude:.claude/CLAUDE.md:.claude"
    "gemini:.gemini/GEMINI.md:.gemini"
    "agent:.agent/AGENT.md:.agent"
)

# Dynamic identification of skills (root folders with SKILL.md)
SKILLS=$(find . -maxdepth 2 -name "SKILL.md" -not -path '*/.*' -exec dirname {} \; | sed 's|^\./||')

echo "🚀 Identifying skills for distribution..."
for SKILL in $SKILLS; do
    echo "  - $SKILL"
done

for AGENT_CONFIG in "${AGENTS[@]}"; do
    IFS=':' read -r NAME INST_FILE TARGET_FOLDER <<< "$AGENT_CONFIG"
    
    echo "📦 Preparing artifact for $NAME..."
    
    STAGING="$DIST_BASE/$TARGET_FOLDER"
    mkdir -p "$STAGING/skills"
    
    # 1. Copy agent instruction file
    if [ -f "$INST_FILE" ]; then
        cp "$INST_FILE" "$STAGING/"
        echo "   ✅ Copied: $INST_FILE"
    else
        echo "   ⚠️  Warning: $INST_FILE not found!"
    fi
    
    # 2. Copy all identified skills
    for SKILL in $SKILLS; do
        cp -r "$SKILL" "$STAGING/skills/"
    done
    echo "   ✅ Copied all skills to $STAGING/skills/"

    # 3. Copy Infrastructure (scripts, Makefile, config)
    echo "   🛠️  Copying infrastructure..."
    cp -r scripts "$STAGING/"
    cp Makefile "$STAGING/"
    cp pyproject.toml "$STAGING/"
    if [ -f "uv.lock" ]; then cp "uv.lock" "$STAGING/"; fi
    if [ -f ".hub-mode" ]; then cp ".hub-mode" "$STAGING/"; fi
    
    # 4. Generate ZIP
    (cd "$DIST_BASE" && zip -rq "../$OUTPUT_DIR/$NAME-skills.zip" "$TARGET_FOLDER")
    echo "   🎁 Generated: $OUTPUT_DIR/$NAME-skills.zip"
done

# --- Special Logic for Antigravity ---
echo "📦 Preparing special artifact: Antigravity..."
ANTIGRAVITY_STAGING="$DIST_BASE/antigravity/.gemini"
mkdir -p "$ANTIGRAVITY_STAGING/skills"
mkdir -p "$ANTIGRAVITY_STAGING/rules"

# 1. Copy GEMINI.md to rules folder
if [ -f ".gemini/GEMINI.md" ]; then
    cp ".gemini/GEMINI.md" "$ANTIGRAVITY_STAGING/rules/"
    echo "   ✅ Copied: .gemini/GEMINI.md -> rules/"
fi

# 2. Copy all skills
for SKILL in $SKILLS; do
    cp -r "$SKILL" "$ANTIGRAVITY_STAGING/skills/"
done
echo "   ✅ Copied all skills to $ANTIGRAVITY_STAGING/skills/"

# 3. Copy Infrastructure for Antigravity
echo "   🛠️  Copying infrastructure for Antigravity..."
cp -r scripts "$ANTIGRAVITY_STAGING/../"
cp Makefile "$ANTIGRAVITY_STAGING/../"
cp pyproject.toml "$ANTIGRAVITY_STAGING/../"
if [ -f "uv.lock" ]; then cp "uv.lock" "$ANTIGRAVITY_STAGING/../"; fi
if [ -f ".hub-mode" ]; then cp ".hub-mode" "$ANTIGRAVITY_STAGING/../"; fi

# 4. Generate Antigravity.zip (starting from antigravity/ folder)
(cd "$DIST_BASE/antigravity" && zip -rq "../../$OUTPUT_DIR/Antigravity.zip" ".gemini")
echo "   🎁 Generated: $OUTPUT_DIR/Antigravity.zip"

echo "✅ All artifacts have been generated in $OUTPUT_DIR/"

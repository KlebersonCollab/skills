#!/bin/bash
# Script to verify the integrity of generated artifacts.
set -e

echo "🔍 Verifying structure of generated artifacts..."

ARTIFACTS=("claude-skills.zip" "gemini-skills.zip" "agent-skills.zip")
AGENTS=(".claude" ".gemini" ".agent")
INSTR_FILES=("CLAUDE.md" "GEMINI.md" "AGENT.md")

if [ ! -d "artifacts" ]; then
    echo "❌ Error: 'artifacts' directory not found!"
    exit 1
fi

for i in "${!ARTIFACTS[@]}"; do
    FILE="artifacts/${ARTIFACTS[$i]}"
    AGENT_DIR="${AGENTS[$i]}"
    INSTR_FILE="${INSTR_FILES[$i]}"
    
    echo "--- Validating $FILE ---"
    
    if [ ! -f "$FILE" ]; then
        echo "❌ Error: $FILE not found!"
        exit 1
    fi
    
    # List files in zip to check structure
    FILES_IN_ZIP=$(unzip -l "$FILE")
    
    # Verify presence of agent directory (e.g., .claude/)
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/"; then
        echo "❌ Error: Directory $AGENT_DIR/ not found in $FILE!"
        exit 1
    fi
    
    # Verify presence of instruction file at the package root (inside agent folder)
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/$INSTR_FILE"; then
        echo "❌ Error: $INSTR_FILE not found in $AGENT_DIR/ inside $FILE!"
        exit 1
    fi
    
    # Verify presence of skills folder
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/skills/"; then
        echo "❌ Error: 'skills/' folder not found in $AGENT_DIR/ inside $FILE!"
        exit 1
    fi
    
    # Verify if at least one known skill is present (e.g., api-architect)
    if ! echo "$FILES_IN_ZIP" | grep -q "$AGENT_DIR/skills/api-architect/"; then
        echo "❌ Error: Skill 'api-architect' not found in $FILE!"
        exit 1
    fi
    
    echo "✅ Structure of $FILE successfully validated!"
done

echo "🎉 All artifacts verified successfully!"

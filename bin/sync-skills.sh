#!/bin/bash

# Lista de skills para sincronizar conforme solicitação do usuário
SKILLS=(
    "architecture"
    "benchmark-expert"
    "brainstorming"
    "clean-code-mentor"
    "django-expert"
    "fastapi-expert"
    "flutter-fvm"
    "git-workflow"
    "observability-expert"
    "python-uv"
    "sdd"
    "skill-factory"
    "token-distiller"
    "youtube-transcript"
)

# Destinos de governança
DESTINATIONS=(
    ".agents/skills"
    ".gemini/skills"
)

# Diretório raiz do projeto
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "🚀 Iniciando sincronização de skills..."

for dest in "${DESTINATIONS[@]}"; do
    TARGET_DIR="$ROOT_DIR/$dest"
    echo ""
    echo "📂 Destino: $dest"
    mkdir -p "$TARGET_DIR"
    
    for skill in "${SKILLS[@]}"; do
        SOURCE_DIR="$ROOT_DIR/$skill"
        if [ -d "$SOURCE_DIR" ]; then
            printf "  %-25s " "Sincronizando $skill..."
            # rsync -av --delete --exclude '.git' preserva permissões e sincroniza apenas mudanças
            rsync -a --delete --exclude '.git' "$SOURCE_DIR/" "$TARGET_DIR/$skill/"
            echo "DONE"
        else
            echo "  ⚠️  Skill não encontrada: $skill"
        fi
    done
done

# Sincronização do GLOBAL_MANDATES.md para arquivos de governança
MANDATES_SOURCE="$ROOT_DIR/.specs/codebase/GLOBAL_MANDATES.md"
MANDATES_TARGETS=(
    ".gemini/GEMINI.md"
    ".gemini/rules/agent.md"
    ".claude/CLAUDE.md"
    ".agents/rules/agent.md"
    ".agents/AGENTS.md"
)

if [ -f "$MANDATES_SOURCE" ]; then
    echo ""
    echo "📜 Sincronizando GLOBAL_MANDATES.md..."
    for target in "${MANDATES_TARGETS[@]}"; do
        TARGET_PATH="$ROOT_DIR/$target"
        mkdir -p "$(dirname "$TARGET_PATH")"
        cp "$MANDATES_SOURCE" "$TARGET_PATH"
        echo "  - Copiado para: $target"
    done
else
    echo "⚠️  GLOBAL_MANDATES.md não encontrado em $MANDATES_SOURCE"
fi

echo ""
echo "✅ Todas as skills foram sincronizadas com sucesso!"

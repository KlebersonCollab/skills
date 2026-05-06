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

echo ""
echo "✅ Todas as skills foram sincronizadas com sucesso!"

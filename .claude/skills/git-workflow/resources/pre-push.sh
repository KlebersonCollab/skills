#!/usr/bin/env bash

# SDD Gated Workflow Hook: Enforce build, lint, and tests before pushing
echo "🚀 [SDD] Verificando integridade antes do push (Build, Lint, Tests)..."

# Função para rodar um comando apenas se ele existir no Makefile
check_make_target() {
    local target=$1
    if [ -f Makefile ] && grep -qE "^${target}:" Makefile; then
        echo "⏳ Executando 'make $target'..."
        make "$target" || { echo "❌ Falha no 'make $target'. Push abortado."; exit 1; }
    fi
}

# Função para rodar um comando apenas se ele existir no package.json
check_npm_script() {
    local script=$1
    if [ -f package.json ] && grep -q "\"${script}\":" package.json; then
        echo "⏳ Executando 'npm run $script'..."
        npm run "$script" || { echo "❌ Falha no 'npm run $script'. Push abortado."; exit 1; }
    fi
}

# 1. Verifica Build
check_make_target "build"
check_npm_script "build"

# 2. Verifica Lint
check_make_target "lint"
check_npm_script "lint"

# 3. Verifica Testes
check_make_target "test"
check_npm_script "test"

# 4. Auditoria SDD (se existir localmente)
check_make_target "audit"

echo "✅ Verificações de Qualidade e Governança concluídas com sucesso. Iniciando push..."
exit 0

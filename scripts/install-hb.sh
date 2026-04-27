#!/bin/bash

# HB-CLI Installer
# Usage: curl -sSL https://raw.githubusercontent.com/KlebersonCollab/skills/main/scripts/install-hb.sh | bash

set -e

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; fi
if [ "$ARCH" = "aarch64" ]; then ARCH="arm64"; fi

URL="https://github.com/KlebersonCollab/skills/releases/latest/download/hb-${OS}-${ARCH}"

INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

echo "🚀 Downloading HB-CLI for ${OS}/${ARCH}..."
curl -L "$URL" -o "$INSTALL_DIR/hb"
chmod +x "$INSTALL_DIR/hb"

echo "✨ HB-CLI installed at $INSTALL_DIR/hb"
echo "👉 Make sure $INSTALL_DIR is in your PATH."
echo "💡 Try running: hb --help"

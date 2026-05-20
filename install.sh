#!/bin/bash

GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m'

# Kita kunci di Node.js v22.14.0 (LTS yang kompatibel dengan better-sqlite3)
NODE_VERSION="v22.14.0"
LOCAL_ENV_DIR=".horizon-env"

echo -e "${BLUE}📦 Installing Horizon dependencies...${NC}\n"

# 1. Download Portable Node.js jika belum ada di dalam folder proyek
if [ ! -d "$LOCAL_ENV_DIR" ]; then
    echo -e "${BLUE}Setting up isolated Node.js environment (${NODE_VERSION})...${NC}"
    
    # Deteksi sistem operasi (OS) dan arsitektur agar unduhan akurat
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    if [ "$ARCH" = "x86_64" ]; then ARCH="x64"; fi
    if [ "$ARCH" = "aarch64" ]; then ARCH="arm64"; fi

    FILE_NAME="node-${NODE_VERSION}-${OS}-${ARCH}.tar.xz"
    URL="https://nodejs.org/dist/${NODE_VERSION}/${FILE_NAME}"

    mkdir -p "$LOCAL_ENV_DIR"
    echo -e "Downloading Node.js from $URL..."
    # Unduh dan ekstrak langsung ke folder .horizon-env
    curl -fsSL "$URL" | tar -xJ -C "$LOCAL_ENV_DIR" --strip-components=1
fi

# 2. Arahkan terminal untuk memprioritaskan Node.js dan NPM dari folder lokal proyek
export PATH="$PWD/$LOCAL_ENV_DIR/bin:$PATH"

# 3. Install pnpm secara lokal (di dalam .horizon-env) jika belum ada
if ! command -v pnpm &> /dev/null; then
    echo -e "${BLUE}Installing local pnpm...${NC}"
    npm install -g pnpm
fi

echo -e "${GREEN}Installing UI dependencies...${NC}"
cd ui || exit
pnpm install
cd ..

echo -e "\n${GREEN}✅ Installation complete!${NC}"
echo -e "${BLUE}Run ./start.sh to start the application${NC}\n"
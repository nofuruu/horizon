#!/bin/bash

GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}📦 Installing Horizon dependencies...${NC}\n"


echo -e "${GREEN}Installing UI dependencies...${NC}"
cd ui
pnpm install
cd ..


echo -e "\n${GREEN}✅ Installation complete!${NC}"
echo -e "${BLUE}Run ./start.sh to start the application${NC}\n"

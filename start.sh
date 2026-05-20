#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Starting Horizon Monitoring Dashboard...${NC}\n"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go first.${NC}"
    exit 1
fi

# Check if Node/pnpm is installed
if ! command -v pnpm &> /dev/null; then
    echo -e "${RED}❌ pnpm is not installed. Please install pnpm first.${NC}"
    exit 1
fi

# Function to cleanup on exit
cleanup() {
    echo -e "\n${BLUE}🛑 Shutting down services...${NC}"
    kill $API_PID $UI_PID 2>/dev/null
    exit
}

trap cleanup SIGINT SIGTERM

# Start API in background
echo -e "${GREEN}📡 Starting API server...${NC}"
cd api
go run main.go &
API_PID=$!
cd ..

# Wait a bit for API to start
sleep 2

# Start UI in background
echo -e "${GREEN}🎨 Starting UI server...${NC}"
cd ui
pnpm dev &
UI_PID=$!
cd ..

echo -e "\n${GREEN}✅ Horizon is running!${NC}"
echo -e "${BLUE}📊 UI:  http://localhost:5173${NC}"
echo -e "${BLUE}🔌 API: http://localhost:8080${NC}"
echo -e "\n${BLUE}Press Ctrl+C to stop all services${NC}\n"

# Wait for both processes
wait $API_PID $UI_PID

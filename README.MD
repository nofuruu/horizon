# Horizon 
Real-time Linux system monitoring dashboard dengan GO backend dan svelte


----
## Requirements 

**Backend**
- Go 1.25 (karena menggunakan framework gin)
- Gin Web Framework
- gopsutil (system metrics)

**UI/client**
- sveltekit
- TailwindCSS
- LayerChart 
- Shadcn-Svelte 

---
## Get Started

### 1. Backend Setup (Go)
> Make sure Go 1.25+ is installed

```bash
cd api
go mod init horizon
go get -u github.com/gin-gonic/gin
go get -u github.com/shirou/gopsutil/v3

# Run the server
go run main.go
```
API will be available at `http://localhost:8080`

### 2. Frontend Setup (Svelte)

```bash
cd ui
pnpm install 

# Development mode
pnpm run dev
```
Frontend will be available at `http://localhost:5173`

## API Endpoints

- `GET /health` - Health check
- `GET /api/system` - System information
- `GET /api/cpu` - CPU metrics
- `GET /api/memory` - Memory metrics
- `GET /api/disk` - Disk metrics
- `GET /api/metrics` - All metrics (combined)

## Development

Run backend and frontend simultaneously:

```bash
# Terminal 1: Start the backend
cd api && go run main.go

# Terminal 2: Start the frontend
cd ui && pnpm dev
```

Then open your browser at `http://localhost:5173`

----

## Features

- Real-time CPU monitoring
- Real-time Memory monitoring
- Real-time Disk monitoring
- System information display
- Historical charts (40 seconds data)
- Auto-refresh every 2 seconds
- Dark mode support
---

## Screenshots

| Section | |
|---|---|
| **Dashboard** | ![Dashboard](assets/dashboard.png) |
| **CPU Usage** | ![CPU Usage](assets/cpu_example.png) |
| **System Information** | ![System Info](assets/example.png) |
| **Process Management** | ![Process Termination](assets/terminate.png) |

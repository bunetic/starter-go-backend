# Go Backend Starter

A simple Go backend starter template with frontend demo showing health checks and ping-pong communication.

## File Structure

This template contains the files needed to deploy a Go backend with frontend on Bunetic:

- **`main.go`** - Go HTTP server with health and ping endpoints
- **`go.mod`** - Go module definition (uses only standard library)
- **`index.html`** - Frontend demo with health check and ping-pong buttons
- **`go.sum`** - Required for Go module detection by Dokku (even if empty)
- **`Procfile`** - Tells Bunetic how to start the Go application (`web: ./bin/bunetic-go-backend`)

## Features

- **Health Check Endpoint** (`/health`) - Shows server status, uptime, and timestamp
- **Ping Endpoint** (`/ping`) - Returns "pong" with server info and timestamp
- **Frontend Demo** - Simple HTML interface to test both endpoints
- **Response Time Measurement** - Shows exact milliseconds for each request
- **CORS Enabled** - Ready for cross-origin requests
- **Zero Dependencies** - Uses only Go standard library

## API Endpoints

### GET /health
Returns server health status:
```json
{
  "status": "healthy",
  "timestamp": "2025-09-29T12:00:00Z",
  "uptime": "5m30s"
}
```

### GET /ping
Returns pong response:
```json
{
  "message": "pong",
  "timestamp": "2025-09-29T12:00:00Z",
  "server_id": "go-backend-v1"
}
```

## Why These Files Are Needed

- **Go requires a module** - `go.mod` and `go.sum` define the module for Dokku detection
- **Binary compilation** - Dokku compiles `main.go` into `./bin/bunetic-go-backend` executable
- **Frontend included** - `index.html` demonstrates the backend functionality  
- **Procfile defines startup** - Bunetic reads this to launch the compiled binary from `./bin/`



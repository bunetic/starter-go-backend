package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Response structures
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

type PingResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	ServerID  string    `json:"server_id"`
}

var startTime = time.Now()

// CORS middleware
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// Health endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	
	if r.Method == "OPTIONS" {
		return
	}

	uptime := time.Since(startTime)
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Uptime:    uptime.Round(time.Second).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Ping endpoint
func pingHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	
	if r.Method == "OPTIONS" {
		return
	}

	response := PingResponse{
		Message:   "pong",
		Timestamp: time.Now(),
		ServerID:  "go-backend-v1",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Serve static files (HTML frontend)
func staticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "index.html")
		return
	}
	http.NotFound(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Routes
	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ping", pingHandler)

	fmt.Printf("🚀 Go backend server starting on port %s\n", port)
	fmt.Printf("📊 Health endpoint: http://localhost:%s/health\n", port)
	fmt.Printf("🏓 Ping endpoint: http://localhost:%s/ping\n", port)
	fmt.Printf("🌐 Frontend: http://localhost:%s/\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

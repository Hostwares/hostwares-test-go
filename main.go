package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

func envStatus(key string) string {
	if os.Getenv(key) != "" {
		return "✅ Set"
	}
	return "❌ Not set"
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		appName := os.Getenv("APP_NAME")
		if appName == "" {
			appName = "Hostwares Test Go"
		}

		response := map[string]interface{}{
			"status":     "🚀 Running on Hostwares!",
			"app_name":   appName,
			"framework":  "Go net/http (stdlib)",
			"go_version": runtime.Version(),
			"environment": map[string]string{
				"APP_NAME":     appName,
				"DATABASE_URL": envStatus("DATABASE_URL"),
				"JWT_SECRET":   envStatus("JWT_SECRET"),
				"SMTP_HOST":    envStatus("SMTP_HOST"),
				"S3_BUCKET":    envStatus("S3_BUCKET"),
			},
			"deployed_at": time.Now().UTC().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})

	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

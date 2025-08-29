package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	// "github.com/mdarify1337/backend-go/backend/controllers"
	"github.com/mdarify1337/backend-go/backend/migrations"
	"github.com/mdarify1337/backend-go/backend/services"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow your frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // fixed missing colon
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			log.Println("[CORS] Preflight request handled")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		log.Printf("[CORS] Passing request %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Get env variables from docker-compose
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	// Connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	log.Println("[DB] Connecting to:", dsn)

	// Connect to Postgres
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("[DB] Failed to connect:", err)
	}
	defer db.Close()
	log.Println("[DB] Connection established")

	// Run migrations
	if err := migrations.RunAll(db); err != nil {
		log.Fatal("[DB] Migration failed:", err)
	}
	log.Println("[DB] ✅ All tables are ready")

	// Simple API route
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[API] /ping hit")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"pong"}`))
	})

	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Printf("[API] %s %s\n", r.Method, r.URL.Path)
	// 	if r.Method != http.MethodGet {
	// 		log.Println("[API] Method not allowed on /")
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	log.Println("[API] Fetching users from DB (not implemented)")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte(`{"users":[]}`))
	// })
	services.UserRoutes(mux, db)
	// services.UserRoutes(mux, db)

	// CreateUser endpoint

	// Wrap with CORS
	handler := enableCORS(mux)
	log.Println("🚀 Go backend running on port 3001")
	log.Fatal(http.ListenAndServe(":3001", handler))
}

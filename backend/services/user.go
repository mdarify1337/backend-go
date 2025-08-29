package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/mdarify1337/backend-go/backend/controllers"
)

func UserRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/CreateUser",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /CreateUser %s request\n", r.Method)

			if r.Method == http.MethodOptions {
				log.Println("[API] Preflight request on /CreateUser")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			if r.Method != http.MethodPost {
				log.Println("[API] Invalid method on /CreateUser")
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			log.Println("[API] Handling user creation")
			w.Header().Set("Content-Type", "application/json")
			controllers.CreateUser(db, w, r)
		})
	mux.HandleFunc("/GetUsers",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /GetUsers %s request\n", r.Method)
			if r.Method != http.MethodGet {
				log.Println("[API] Invalid method on /GetUsers")
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			log.Println("[API] Fetching users from DB")
			w.Header().Set("Content-Type", "application/json")
			controllers.GetUsers(db, w, r)
		})
}

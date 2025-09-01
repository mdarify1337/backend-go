package services

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}

			log.Println("[API] Handling user creation")
			w.Header().Set("Content-Type", "application/json")
			controllers.CreateUser(db, w, r)
		},
	)
	mux.HandleFunc("/GetUsers",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /GetUsers %s request\n", r.Method)
			if r.Method != http.MethodGet {
				log.Println("[API] Invalid method on /GetUsers")
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}
			log.Println("[API] Fetching users from DB")
			w.Header().Set("Content-Type", "application/json")
			controllers.GetUsers(db, w, r)
		},
	)
	mux.HandleFunc("/UpdateUser",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /UpdateUser %s request\n", r.Method)

			if r.Method == http.MethodOptions {
				log.Println("[API] Preflight request on /UpdateUser")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			if r.Method != http.MethodPut {
				log.Println("[API] Invalid method on /UpdateUser")
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}

			log.Println("[API] Handling user update")
			w.Header().Set("Content-Type", "application/json")
			controllers.UpdateUser(db, w, r)
		},
	)

	mux.HandleFunc("/GetUser",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /GetUser %s request\n", r.Method)
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}
			controllers.GetUser(db, w, r)
		},
	)

	mux.HandleFunc("/DeleteUser",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /DeleteUser %s request\n", r.Method)
			if r.Method != http.MethodDelete {
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}
			// Extract user ID from query parameters
			idStr := r.URL.Query().Get("id")
			if idStr == "" {
				http.Error(w, "Missing user ID", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
			controllers.DeleteUser(db, w, r, id)
		},
	)

	mux.HandleFunc("/SignInUser",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /SignInUser %s request\n", r.Method)
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}
			controllers.SignInUser(db, w, r)
		})
}

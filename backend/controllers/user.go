package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mdarify1337/backend-go/backend/models"
)

// CreateUser inserts a new user into the DB
func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Timestamps
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	// Insert into DB
	query := `
		INSERT INTO users (username, email, password, first_name, 
		last_name, created_at, updated_at, picture)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`
	err := db.QueryRow(query,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.UpdatedAt,
		user.Picture,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, fmt.Sprintf("DB insert error: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with created user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	fmt.Println("✅ User saved:", user)
}

func GetUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, 
		username, email, 
		password, first_name, 
		last_name, created_at, 
		updated_at, picture 
		FROM users;
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB query error: %v", err),
			http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password,
			&user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt,
			&user.Picture); err != nil {
			http.Error(w, fmt.Sprintf("Row scan error: %v", err),
				http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Rows error: %v", err),
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

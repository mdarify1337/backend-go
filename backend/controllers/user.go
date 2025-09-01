package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update timestamp
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	// Update DB record
	query := `
		UPDATE users 
		SET username=$1, email=$2, password=$3, 
		    first_name=$4, last_name=$5, 
		    updated_at=$6, picture=$7
		WHERE id=$8;
	`
	result, err := db.Exec(query,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
		user.Picture,
		user.ID,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB update error: %v", err),
			http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("RowsAffected error: %v", err),
			http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No user found with given ID", http.StatusNotFound)
		return
	}

	// Respond with updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	fmt.Println("✅ User updated:", user)
}

func GetUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Extract ID from query param
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

	var user models.User
	query := `SELECT id, username, email, password, first_name, last_name, created_at, updated_at, picture 
	          FROM users WHERE id=$1;`

	err = db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password,
		&user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt, &user.Picture)

	if err == sql.ErrNoRows {
		http.Error(w, "No user found with given ID", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("DB query error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	fmt.Println("✅ User fetched:", user)
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("[Controller] DeleteUser called with id=%d\n", id)

	// Prepare the delete statement
	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		http.Error(w, "Failed to prepare delete statement", http.StatusInternalServerError)
		log.Printf("Error preparing delete statement: %v", err)
		return
	}
	defer stmt.Close()

	// Execute the query
	result, err := stmt.Exec(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		log.Printf("Error executing delete: %v", err)
		return
	}

	// Check if a row was affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows", http.StatusInternalServerError)
		log.Printf("Error checking affected rows: %v", err)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		log.Printf("No user found with id=%d", id)
		return
	}

	// Respond with JSON
	response := map[string]string{"message": "User deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Printf("[Controller] User with id=%d deleted successfully\n", id)
}

func SignInUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, username, email, password, first_name, last_name, created_at, updated_at, picture 
	          FROM users WHERE email=$1 AND password=$2;`

	err := db.QueryRow(query, creds.Username, creds.Password).Scan(&user.ID, &user.Username, &user.Email,
		&user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt, &user.Picture)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("DB query error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	fmt.Println("✅ User signed in:", user)
}

package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mdarify1337/backend-go/backend/models"
)

type RequestContext struct {
	DB *sql.DB
	W  http.ResponseWriter
	R  *http.Request
}

func CreateProduct(data RequestContext) {
	var product models.Product
	if err := json.NewDecoder(data.R.Body).Decode(&product); err != nil {
		http.Error(data.W, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Timestamps
	product.CreatedAt = time.Now().Format(time.RFC3339)
	product.UpdatedAt = time.Now().Format(time.RFC3339)

	// Insert into DB
	query := `
		INSERT INTO products (name, description, price, quantity, 
		created_at, updated_at, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`
	err := data.DB.QueryRow(query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CreatedAt,
		product.UpdatedAt,
		product.UserID,
	).Scan(&product.ID)

	if err != nil {
		http.Error(data.W, fmt.Sprintf("DB insert error: %v", err),
			http.StatusInternalServerError)
		return
	}

	// Respond with created product
	data.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(data.W).Encode(product)
	fmt.Println("✅ Product saved:", product)
}

func GetProducts(data RequestContext) {
	query := `
			SELECT id, 
			name, 
			description, 
			price, 
			quantity, 
			created_at, 
			updated_at, 
			user_id FROM products;
		`

	rows, err := data.DB.Query(query)
	if err != nil {
		http.Error(data.W, fmt.Sprintf("DB query error: %v", err),
			http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name,
			&product.Description, &product.Price,
			&product.Quantity, &product.CreatedAt,
			&product.UpdatedAt, &product.UserID); err != nil {
			http.Error(data.W, fmt.Sprintf("Row scan error: %v", err),
				http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		http.Error(data.W, fmt.Sprintf("Rows error: %v", err),
			http.StatusInternalServerError)
		return
	}

	data.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(data.W).Encode(products)
}

func GetProductByID(data RequestContext) {
	// Extract ID from query parameter (?id=5)
	idStr := data.R.URL.Query().Get("id")
	if idStr == "" {
		http.Error(data.W, "Missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(data.W, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Query DB for a single product
	query := `
        SELECT id, 
		name, 
		description, 
		price, 
		quantity, 
		created_at, 
		updated_at, 
		user_id
        FROM products
        WHERE id = $1;
    `

	var product models.Product
	err = data.DB.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Description,
		&product.Price, &product.Quantity, &product.CreatedAt,
		&product.UpdatedAt, &product.UserID,
	)

	if err == sql.ErrNoRows {
		http.Error(data.W, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(data.W, fmt.Sprintf("DB query error: %v", err),
			http.StatusInternalServerError)
		return
	}

	// Respond with the product
	data.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(data.W).Encode(product)
}

func UpdateProduct(data RequestContext) {
	var product models.Product
	if err := json.NewDecoder(data.R.Body).Decode(&product); err != nil {
		http.Error(data.W, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure ID is provided
	if product.ID == 0 {
		http.Error(data.W, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Update timestamp
	product.UpdatedAt = time.Now().Format(time.RFC3339)

	// Update DB record
	query := `
		UPDATE products
		SET name=$1, description=$2, price=$3, quantity=$4, updated_at=$5, user_id=$6
		WHERE id=$7;
	`
	result, err := data.DB.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.UpdatedAt,
		product.UserID,
		product.ID,
	)
	if err != nil {
		http.Error(data.W, fmt.Sprintf("DB update error: %v", err),
			http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(data.W, fmt.Sprintf("RowsAffected error: %v", err),
			http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(data.W, "No product found with given ID", http.StatusNotFound)
		return
	}

	// Respond with updated product
	data.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(data.W).Encode(product)
	fmt.Println("✅ Product updated:", product)
}

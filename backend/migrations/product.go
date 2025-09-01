package migrations

import (
	"database/sql"
	"fmt"
)

func CreateProductTable(db *sql.DB) error {

	query := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			price DECIMAL(10, 2) NOT NULL,
			quantity INT NOT NULL,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
			user_id INT REFERENCES users(id) ON DELETE CASCADE
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create products table: %w",
			err)
	}
	return nil
}

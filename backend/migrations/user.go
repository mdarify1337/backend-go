// backend/migrations/users.go
package migrations

import (
	"database/sql"
	"fmt"
)

func CreateUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(255) NOT NULL,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		picture TEXT
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	return nil
}

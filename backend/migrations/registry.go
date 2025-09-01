// backend/migrations/registry.go
package migrations

import "database/sql"

func RunAll(db *sql.DB) error {
	if err := CreateUsersTable(db); err != nil {
		return err
	}
	// add more like: if err := CreateMeetingsTable(db);
	// err != nil { return err }
	return nil
}

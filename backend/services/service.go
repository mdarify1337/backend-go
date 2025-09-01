package services

import (
	"database/sql"
	"net/http"
)

func RunAllServices(mux *http.ServeMux, db *sql.DB) {
	UserRoutes(mux, db)
	ProductRoutes(mux, db)
}

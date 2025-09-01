package services

import (
	"database/sql"
	"github.com/mdarify1337/backend-go/backend/controllers"
	"log"
	"net/http"
)

func ProductRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/CreateProduct",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /CreateProduct %s request\n", r.Method)

			if r.Method == http.MethodOptions {
				log.Println("[API] Preflight request on /CreateProduct")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			if r.Method != http.MethodPost {
				log.Println("[API] Invalid method on /CreateProduct")
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}

			log.Println("[API] Handling product creation")
			w.Header().Set("Content-Type", "application/json")
			controllers.CreateProduct(controllers.RequestContext{
				DB: db,
				W:  w,
				R:  r,
			})
		},
	)
	mux.HandleFunc("/GetProducts",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /GetProducts %s request\n", r.Method)
			if r.Method != http.MethodGet {
				log.Println("[API] Invalid method on /GetProducts")
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}
			log.Println("[API] Fetching products from DB")
			w.Header().Set("Content-Type", "application/json")
			controllers.GetProducts(controllers.RequestContext{
				DB: db,
				W:  w,
				R:  r,
			})
		},
	)

	mux.HandleFunc("/GetProductByID/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /GetProductByID %s request\n", r.Method)
			if r.Method != http.MethodGet {
				log.Println("[API] Invalid method on /GetProductByID")
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}
			log.Println("[API] Fetching product by ID from DB")
			w.Header().Set("Content-Type", "application/json")
			controllers.GetProductByID(controllers.RequestContext{
				DB: db,
				W:  w,
				R:  r,
			})
		},
	)

	mux.HandleFunc("/UpdateProduct/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API] /UpdateProduct %s request\n", r.Method)

			if r.Method == http.MethodOptions {
				log.Println("[API] Preflight request on /UpdateProduct")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			if r.Method != http.MethodPut {
				log.Println("[API] Invalid method on /UpdateProduct")
				http.Error(w, "Method not allowed",
					http.StatusMethodNotAllowed)
				return
			}

			log.Println("[API] Handling product update")
			w.Header().Set("Content-Type", "application/json")
			controllers.UpdateProduct(controllers.RequestContext{
				DB: db,
				W:  w,
				R:  r,
			})
		},
	)
}

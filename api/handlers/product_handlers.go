package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"api/models"
)

type ProductHandlers struct {
	DB *sql.DB
}

func NewProductHandlers(db *sql.DB) *ProductHandlers {
	return &ProductHandlers{DB: db}
}

func (h *ProductHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{}

	rows, err := h.DB.Query("SELECT name, scu, link, image_link, description, id_product FROM products")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Name, &p.SKU, &p.Link, &p.ImageLink, &p.Description, &p.IDProduct); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		products = append(products, p)
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var p models.Product

	err := h.DB.QueryRow("SELECT name, scu, link, image_link, description, id_product FROM products WHERE id_product = $1",
		vars["id"]).Scan(&p.Name, &p.SKU, &p.Link, &p.ImageLink, &p.Description, &p.IDProduct)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Product not found")
		return
	} else if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (h *ProductHandlers) GetProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	products := []models.Product{}

	rows, err := h.DB.Query("SELECT name, scu, link, image_link, description, id_product FROM products WHERE name ILIKE $1",
		"%"+vars["name"]+"%")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Name, &p.SKU, &p.Link, &p.ImageLink, &p.Description, &p.IDProduct); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		products = append(products, p)
	}

	if len(products) == 0 {
		respondWithError(w, http.StatusNotFound, "No products found with this name")
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
} 
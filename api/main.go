package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Product struct {
	Name        string `json:"name"`
	SKU         string `json:"sku"`
	Link        string `json:"link"`
	ImageLink   string `json:"image_link"`
	Description string `json:"description"`
	IDProduct   int64  `json:"id_product"`
}

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

func (app *App) Initialize() {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		connectionString = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	}

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/products/{id:[0-9]+}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/products/name/{name}", app.getProductByName).Methods("GET")
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{}

	rows, err := app.DB.Query("SELECT name, scu, link, image_link, description, id_product FROM products")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Name, &p.SKU, &p.Link, &p.ImageLink, &p.Description, &p.IDProduct); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		products = append(products, p)
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var p Product

	err := app.DB.QueryRow("SELECT name, scu, link, image_link, description, id_product FROM products WHERE id_product = $1",
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

func (app *App) getProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	products := []Product{}

	// Используем ILIKE для поиска без учета регистра и с частичным совпадением
	rows, err := app.DB.Query("SELECT name, scu, link, image_link, description, id_product FROM products WHERE name ILIKE $1",
		"%"+vars["name"]+"%")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
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

func main() {
	app := App{}
	app.Initialize()

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", app.Router))
} 
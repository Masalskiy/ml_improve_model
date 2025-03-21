package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"api/handlers"
)

type App struct {
	Router          *mux.Router
	ProductHandlers *handlers.ProductHandlers
}

func (app *App) Initialize() {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		connectionString = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.ProductHandlers = handlers.NewProductHandlers(db)
	app.Router = mux.NewRouter()
	app.Router.Use(corsMiddleware)
	app.initializeRoutes()
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/products", app.ProductHandlers.GetProducts).Methods("GET", "OPTIONS")
	app.Router.HandleFunc("/products/{id:[0-9]+}", app.ProductHandlers.GetProduct).Methods("GET", "OPTIONS")
	app.Router.HandleFunc("/products/name/{name}", app.ProductHandlers.GetProductByName).Methods("GET", "OPTIONS")
}

func main() {
	app := App{}
	app.Initialize()

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", app.Router))
} 
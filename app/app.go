package app

import (
	"fmt"
	"github.com/dougmendes/grocerylist-backend/domain"
	"github.com/dougmendes/grocerylist-backend/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

func Start() {
	// Start the app
	router := mux.NewRouter()
	dbClient := getDbClient()
	productRepository := domain.NewProductRepositoryDb(dbClient)
	ph := ProductHandler{service.NewProductService(productRepository)}

	// Routes
	router.HandleFunc("/products", ph.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{product_id:[0-9]+}", ph.GetProductById).Methods("GET")

	// Start the server
	http.ListenAndServe("localhost:8080", router)

	// Close the db client

}

func getDbClient() *sqlx.DB {
	// Get the db client
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

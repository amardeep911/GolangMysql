package main

import (
	"log"
	"net/http"

	"github.com/amardeep/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// Start the server on port 9010
	port := "9010"
	addr := "localhost:" + port
	log.Printf("Server is starting on http://%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

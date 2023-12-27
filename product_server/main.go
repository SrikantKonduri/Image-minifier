package main

import (
	"log"
	"net/http"
	"product_server/databases"
	"product_server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db := databases.ConnectDB()
	db = db

	router := mux.NewRouter()
	router.HandleFunc("/products/add", handlers.HandleAddProduct)
	router.HandleFunc("/", handlers.HandleHome)

	log.Fatal(http.ListenAndServe(":8000", router))
}

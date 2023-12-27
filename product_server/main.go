package main

import (
	"fmt"
	"log"
	"net/http"
	"product_server/databases"
	"product_server/handlers"
	messagequeue "product_server/message_queue"
	"product_server/utils"

	"github.com/gorilla/mux"
)

func main() {
	ch, ctx, err := messagequeue.ConnectToQueue()

	if err != nil {
		utils.FailOnError(err, "Cannot connect to queue")
	}
	fmt.Println("[+] Connnected to Message Queue")
	db := databases.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/products/add", handlers.HandleAddProduct(db, ch, ctx)).Methods("POST")
	router.HandleFunc("/", handlers.HandleHome)

	log.Fatal(http.ListenAndServe(":8000", router))
}

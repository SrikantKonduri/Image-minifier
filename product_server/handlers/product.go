package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"product_server/databases"
	"strings"
	"time"

	messagequeue "product_server/message_queue"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RecProduct struct {
	UserID             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	ProductPrice       float32  `json:"product_price"`
}

type Response struct {
	Pid     int    `json:"pid"`
	Message string `json:"message"`
}

func HandleAddProduct(db *sql.DB, ch *amqp.Channel, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the JSON request body
		var req RecProduct
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		fmt.Println("D: ", decoder)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			fmt.Println("[x] ", err)
			return
		}

		userExist, err := databases.VerifyUser(db, req.UserID)

		if err != nil {
			http.Error(w, "Failed to verify user", http.StatusInternalServerError)
			return
		}
		if !userExist {
			http.Error(w, "User does not exist", http.StatusUnauthorized)
			fmt.Println("[+] User does not exist")
			return
		}

		product_images := strings.Join(req.ProductImages, ",")
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("Received message: %+v\n", req)
		var newProduct databases.Product = databases.Product{1, req.ProductName, req.ProductDescription, req.ProductPrice, product_images, product_images, currentTime, currentTime}

		pid, err := databases.AddProduct(db, newProduct)
		if err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			fmt.Println("[x] ", err)
			return
		}
		fmt.Println("Product Id: ", pid)
		err = messagequeue.ProduceMessage(pid, ch, ctx, os.Getenv("MSG_QUEUE_NAME"))
		if err != nil {
			http.Error(w, "Failed to Publish into queue", http.StatusInternalServerError)
			fmt.Println("[x] ", err)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(pid)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Producer Home</h1>"))
}

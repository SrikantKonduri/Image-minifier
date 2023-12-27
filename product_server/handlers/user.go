package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"product_server/databases"
)

type RecUser struct {
	UserName      string `json:"user_name"`
	UserMobile    string `json:"user_mobile"`
	UserLatitude  string `json:"latitude"`
	UserLongitude string `json:"longitude"`
}

func HandleAddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the JSON request body
		var req RecUser
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		fmt.Println("D: ", decoder)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			fmt.Println("[x] ", err)
			return
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("Received message: %+v\n", req)
		var newUser databases.User = databases.User{1, req.UserName, req.UserMobile, req.UserLatitude, req.UserLongitude, currentTime, currentTime}

		uid, err := databases.AddUser(db, newUser)
		if err != nil {
			http.Error(w, "Failed to add userxs", http.StatusBadRequest)
			fmt.Println("[x] ", err)
			return
		}
		fmt.Println("User Id: ", uid)
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(uid)
	}
}

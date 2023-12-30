package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type RequestData struct {
	UserId             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
}

func TestAddProduct(t *testing.T) {

	t.Run("Valid User", func(t *testing.T) {
		url := "http://localhost:8000/products/add"

		data := RequestData{UserId: 1, ProductName: "Spice", ProductDescription: "Classic phones ever", ProductImages: []string{"https://drop.ndtv.com/TECH/product_database/images/2222016113841AM_635_htc_desire_530.jpeg"}}

		jsonData, err := json.Marshal(data)
		if err != nil {
			t.Error("Error Encoding JSON")
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Error("Unable to send request")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Error("Not able to insert product")
		}
	})

	t.Run("Invalid User", func(t *testing.T) {
		url := "http://localhost:8000/products/add"

		// Invalid User ID
		data := RequestData{UserId: -1, ProductName: "Some mobile", ProductDescription: "Classic phones ever", ProductImages: []string{"https://drop.ndtv.com/TECH/product_database/images/2222016113841AM_635_htc_desire_530.jpeg"}}

		jsonData, err := json.Marshal(data)
		if err != nil {
			t.Error("Error Encoding JSON")
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Error("Unable to send request")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusUnauthorized {
			t.Error("Not able to insert product")
		}
	})

}

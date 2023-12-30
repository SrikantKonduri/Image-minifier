// image/analysis_test.go

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func AddProduct() {
	url := "http://localhost:8000/products/add"

	data := RequestData{UserId: 1, ProductName: "Spice", ProductDescription: "Classic phones ever", ProductImages: []string{"https://drop.ndtv.com/TECH/product_database/images/2222016113841AM_635_htc_desire_530.jpeg"}}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error Encoding JSON")
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Unable to send request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Not able to insert product")
		return
	}
}

func BenchmarkAddProducts(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddProduct()
	}
}

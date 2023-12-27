package handlers

import "net/http"

func HandleAddProduct(http.ResponseWriter, *http.Request) {

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Product server Home</h1>"))
}

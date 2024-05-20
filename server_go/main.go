package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"` 
}

var products = []Product{
	{ID: 1, Name: "Product 1", Price: 100, Amount: 10},
	{ID: 2, Name: "Product 2", Price: 200, Amount: 20},
	{ID: 3, Name: "Product 3", Price: 300, Amount: 30},
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(products)
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var paymentData struct {
		ProductID int `json:"productID"`
		Quantity  int `json:"quantity"`
	}
	err := json.NewDecoder(r.Body).Decode(&paymentData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var foundProduct *Product
	for i := range products {
		if products[i].ID == paymentData.ProductID {
			foundProduct = &products[i]
			break
		}
	}
	if foundProduct == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	foundProduct.Amount -= paymentData.Quantity
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment successful!"))
}


func main() {
	http.HandleFunc("/products", getProductsHandler)
	http.HandleFunc("/payment", paymentHandler)

	log.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

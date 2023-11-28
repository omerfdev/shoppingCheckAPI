package main

import (
	"encoding/json"
	"net/http"
)

// Product struct represents a product with its ID and availability status.
type Product struct {
	ID     string `json:"id"`
	Exists bool   `json:"exists"`
	Stock  int    `json:"stock"`
}

// ProductDatabase simulates a database of products.
var ProductDatabase = map[string]int{
	"product123": 10,
	"product456": 5,
	// Diğer ürünleri ekleyebilirsiniz.
}

// ProductHandler handles requests for product information.
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// URL'den "id" parametresini alın.
	id := r.URL.Query().Get("id")

	// ID'ye sahip ürünü kontrol edin.
	stock, found := ProductDatabase[id]
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// JSON yanıtını oluşturun.
	product := Product{ID: id, Exists: true, Stock: stock}
	response, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// JSON yanıtını gönderin.
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	// "/product" endpoint'i için ProductHandler'ı tanımlayın.
	http.HandleFunc("/product", ProductHandler)

	// 8080 portunda bir HTTP sunucusu başlatın.
	http.ListenAndServe(":8080", nil)
}

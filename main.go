package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Product struct {
	ID int  			`json:"id"`
	Title string		`json:"title"`
	Description string	`json:"description"`
	Price float64		`json:"price"`

}
var products []Product

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am Sakib")
}
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if(r.Method != http.MethodGet) {
		http.Error(w, "This is a GET request route", http.StatusBadRequest)
		return
	}

	encode := json.NewEncoder(w)
	err := encode.Encode(products)

	if(err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProductsHandler)
	fmt.Println("Server started at localhost:5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}

func init () {
	product1 := Product{
		ID: 1,
		Title: "Mango",
		Description: "1st Fruit",
		Price: 22,
	} 
	product2 := Product{
		ID: 2,
		Title: "Banana",
		Description: "2nd Fruit",
		Price: 12,
	} 
	products = append(products, product1, product2)
}
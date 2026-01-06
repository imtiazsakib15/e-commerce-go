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
	handleCors(w)

	if(r.Method != http.MethodGet) {
		http.Error(w, "This is a GET request route", http.StatusBadRequest)
		return
	}
	sendData(w, products, http.StatusOK)
}
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightRequest(w, r)

	if(r.Method != http.MethodPost) {
		http.Error(w, "This is a POST request route", http.StatusBadRequest)
		return
	}

	var newProduct Product
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newProduct)
	if(err != nil) {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	sendData(w, newProduct, http.StatusCreated)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-control-allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodPost) {
		http.Error(w, "This is a POST request route", http.StatusBadRequest)
		return
	}
}
func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encode := json.NewEncoder(w)
	encode.Encode(data)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", handler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProductsHandler)
	mux.HandleFunc("/create-product", createProductHandler)
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
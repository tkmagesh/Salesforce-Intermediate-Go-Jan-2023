package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products []Product

func init() {
	products = []Product{
		{101, "Pen", 10},
		{102, "Pencil", 5},
		{103, "Marker", 50},
	}
}

func logger(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		log.Printf("%s - %s\n", r.Method, r.URL)
		handle(w, r, params)
	}
}

func profile(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()
		defer func() {
			elapsed := time.Since(start)
			fmt.Printf("%s - time taken %v: ", r.URL, elapsed)
		}()
		handle(w, r, params)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hi from go web server! [using httprouter]"))
}

func getProductsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getProductHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("id")
	w.Write([]byte(fmt.Sprintf("Product [id = %s] will be returned", productId)))
}

func postProductsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newProduct Product
	if err := decoder.Decode(&newProduct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newProduct.Id = len(products) + 101
	products = append(products, newProduct)
	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusCreated)
	if err := encoder.Encode(newProduct); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	/*
		router.GET("/products", profile(logger(getProductsHandler)))
		router.GET("/products/:id", profile(logger(getProductHandler)))
		router.POST("/products", profile(logger(postProductsHandler)))

	*/

	router.GET("/products", getProductsHandler)
	router.GET("/products/:id", getProductHandler)
	router.POST("/products", postProductsHandler)
	http.ListenAndServe(":8080", router)
}

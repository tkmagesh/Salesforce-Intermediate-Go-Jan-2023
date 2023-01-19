package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

/*
type Product struct {
	Id   int
	Name string
	Cost float32
}
*/

var products = []Product{
	{Id: 101, Name: "Pen", Cost: 10},
	{Id: 102, Name: "Pencil", Cost: 20},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		/*
			connect to the db
			get the data
			serialize into JSON
			send the response
		*/
		fmt.Fprint(w, "All customer data will be served")
	case "POST":
		/*
			deserialize the data from the request
			connect to the db
			process the data
			serialize into JSON
			send the response
		*/
		fmt.Fprint(w, "The new customer will be added")
	case "PUT":
		fmt.Fprint(w, "The given customer will be updated")
	case "DELETE":
		fmt.Fprint(w, "The given customer will be deleted")
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(products); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var newProduct Product
		if err := decoder.Decode(&newProduct); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newProduct.Id = len(products) + 101
		products = append(products, newProduct)
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(newProduct); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL)
		handler(w, r)
	}
}

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("Time taken :", elapsed)
	}
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func main() {
	/*
		http.HandleFunc("/", profile(log(indexHandler)))
		http.HandleFunc("/customers", profile(log(customersHandler)))
		http.HandleFunc("/products", profile(log(productsHandler)))
	*/
	http.HandleFunc("/", chain(indexHandler, log, profile))
	http.HandleFunc("/customers", chain(customersHandler, log, profile))
	http.HandleFunc("/products", chain(productsHandler, log, profile))
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
)

//Custom Multiplexer implementation
type App struct {
	handlers map[string]http.HandlerFunc
}

func (app *App) Register(pattern string, handler http.HandlerFunc) {
	app.handlers[pattern] = handler
}

func (app App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.Method, r.URL)
	if handler, exists := app.handlers[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func NewApp() *App {
	return &App{
		handlers: make(map[string]http.HandlerFunc),
	}
}

//Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Products requests are processed")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Customers requests are processed")
}

func main() {
	app := NewApp()
	app.Register("/", indexHandler)
	app.Register("/products", productsHandler)
	app.Register("/customers", customersHandler)
	http.ListenAndServe(":8080", app)
}

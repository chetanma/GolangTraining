package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RoutingDemo() {
	//store := NewProductsDataStore()

	r := mux.NewRouter()
	productsRouter := r.PathPrefix("/products").Subrouter()
	productsRouter.HandleFunc("", LoggerMiddleware(GetProductsHandler)).Methods("GET")
	productsRouter.HandleFunc("/{id}", LoggerMiddleware(GetProductHandler)).Methods("GET")
	//Just as a demo that any func/method with signature func(http.ResponseWriter, http.Request) can be used as handler for http
	//	productsRouter.HandleFunc("/{id}", LoggerMiddleware(store.ProductsHttpHandler)).Methods("GET")

	hotelsRouter := r.PathPrefix("/hotels").Subrouter()
	hotelsRouter.HandleFunc("", LoggerMiddleware(GetHotelsHandler)).Methods("GET")
	hotelsRouter.HandleFunc("/{id}", LoggerMiddleware(GetHotelHandler)).Methods("GET")

	r.HandleFunc("/info", LoggerMiddleware(InfoHandler))
	r.HandleFunc("/", LoggerMiddleware(InfoHandler))
	http.ListenAndServe(":8081", r)
}

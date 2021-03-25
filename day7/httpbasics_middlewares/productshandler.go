package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, r.URL.Path, " : ", vars["id"])
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, r.URL.Path, " : ", vars["id"])
}

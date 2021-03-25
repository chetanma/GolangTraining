package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHotelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, r.URL.Path, " : ", vars["id"])
}

func GetHotelsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, r.URL.Path, " : ", vars["id"])
}

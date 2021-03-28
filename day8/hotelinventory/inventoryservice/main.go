package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Hotel struct {
	Name string
}

func CreateHotelHandler(store HotelStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dataBytes, _ := ioutil.ReadAll(r.Body)
		//close the body deferred

		var hotels []Hotel
		_ = json.Unmarshal(dataBytes, &hotels)
		_ = store.AddHotels(hotels)

	}
}

func FetchHotelsHandler(store HotelStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := store.GetHotels()

		dataBytes, _ := json.Marshal(data)

		fmt.Fprint(w, string(dataBytes))

	}
}

func main() {
	var port = flag.Int("p", 8081, "please specify the port to run")
	var brand = flag.String("b", "Source-A", "Please specify brand")
	flag.Parse()

	fmt.Println("Starting server as ", *brand, " on port ", *port)
	store := HotelStoreJson{}

	r := mux.NewRouter()

	r.HandleFunc("/hotels", CreateHotelHandler(&store)).Methods("POST")
	r.HandleFunc("/hotels", FetchHotelsHandler(&store)).Methods("GET")

	// http.HandleFunc("/hotels", CreateHotelHandler(&store))
	err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
	fmt.Println(err)
}

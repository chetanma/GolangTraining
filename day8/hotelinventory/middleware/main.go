package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Hotel struct {
	Name string
}

type HotelOutPut struct {
	data []Hotel
	err  error
}

func FetchDataFromSourceUrl(url string) []Hotel {
	client := &http.Client{}
	payload := strings.NewReader(``)
	req, err := http.NewRequest("GET", url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	hotels := make([]Hotel, 0, 100)
	json.Unmarshal(body, &hotels)
	return hotels
}

func FetchHotelsHandler(sources []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		allhotel := make([]Hotel, 0, 200)
		// this is serial code fetch data from url sources one by one
		// We need to make it concurrent using concurrent go routine workers returning chan []Hotel
		for _, s := range sources {
			url := s + "/hotels"
			hotelList := FetchDataFromSourceUrl(url)
			allhotel = append(allhotel, hotelList...)
		}

		allHotelsData, _ := json.Marshal(allhotel)

		fmt.Fprint(w, string(allHotelsData))
	}
}
func main() {
	var port = flag.Int("p", 8081, "please specify the port to run")
	var source = flag.String("s", "", "Please specify sourceurl")

	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/hotels", FetchHotelsHandler([]string{*source})).Methods("GET")

	// http.HandleFunc("/hotels", CreateHotelHandler(&store))
	err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
	fmt.Println(err)
}

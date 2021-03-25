package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoggerMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling: ", r.URL)
		ctx := r.Context()
		userType := ctx.Value("authuser")
		if userType != nil {
			user, ok := userType.(string)
			if ok == true {
				fmt.Println("Request is authenticated, user:", user)
			}
		}

		h(w, r)
	}
}

func BasicAuthMiddleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !AuthCheck(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Not Authorized")
			return
		}
		context := context.WithValue(r.Context(), "authuser", user)
		h(w, r.WithContext(context))
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is go http server for demo of handlers and middleware!")
}

func HotelsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hotels data")
}

func CreateHotelHandler(store HotelStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := store.GetHotels()
		dataBytes, _ := json.Marshal(data)
		fmt.Fprint(w, string(dataBytes))
		//check the HTTP method
		//GET or POST Or DELETE or PUT
		//if GET then
		//----//if ID is provided in url param then return that specific hotel. Handle error in case of invalid or unknown id
		//---//If Id is not provided then it means return all the hotels
		//---only user1 should be able to fetch all the data. If user 2 return that you are allowed to get all the hotels
		//if POST
		//---//then accept the hotel object. Append to store. Add a method to interface to AddHotel(Hotel). Handle error for bad input
		//implement delete
		//---only user1 should be able to delete
		//implement put for modify
	}
}

func main() {
	HttpServerAndMiddlewareDemo()
}

func HttpServerAndMiddlewareDemo() {
	hStore := NewHotelsDataStore()

	http.HandleFunc("/info", LoggerMiddleware(InfoHandler))
	//http.HandleFunc("/hotels", LoggerMiddleware(HotelsHandler))
	http.HandleFunc("/hotels", BasicAuthMiddleWare(LoggerMiddleware(CreateHotelHandler(&hStore))))
	http.ListenAndServe(":8081", nil)
}

func AuthCheck(user, pass string) bool {
	if user == "user1" && pass == "pass1" {
		return true
	}

	if user == "user2" && pass == "pass2" {
		return true
	}

	return false
}

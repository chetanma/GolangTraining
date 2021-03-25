package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Id   int
}

type AppServer struct{}

func (as *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("URL: ", r.URL)
	fmt.Println("URL query params: ", r.URL.Query())
	fmt.Fprint(w, "Hello from go!")
}

func SimpleHttpServerDemo() {

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is info url of sample http server")
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("Method: ", r.Method)
		fmt.Println("URL: ", r.URL)
		fmt.Println("URL query params: ", r.URL.Query())
		fmt.Println(r.Form)
		fmt.Fprint(w, "Form accepted")
	})

	http.HandleFunc("/body", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method: ", r.Method)
		fmt.Println("URL: ", r.URL)
		fmt.Println("URL query params: ", r.URL.Query())
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err == nil {
			fmt.Println(string(body))
		} else {
			fmt.Println(err)
		}
		fmt.Fprint(w, "accepted")
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method: ", r.Method)
		fmt.Println("URL: ", r.URL)
		fmt.Println("URL query params: ", r.URL.Query())

		fmt.Println(r.Header["Content-Type"])

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err == nil {
				var p Person
				err := json.Unmarshal(body, &p)
				if err == nil {
					fmt.Println(p)
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
			fmt.Fprint(w, "accepted")
		} else { // assume http get
			person := Person{
				Name: "Person1",
				Id:   1002,
			}
			data, err := json.Marshal(person)
			if err == nil {
				fmt.Fprint(w, string(data))
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "error")
			}
		}

	})

	//http.ListenAndServe(":8081", &AppServer)

	http.ListenAndServe(":8081", nil)

}

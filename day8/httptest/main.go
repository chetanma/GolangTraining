package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("p", 8085, "specify port number")

	flag.Parse()

	ComputeSquareRoot(10)
	fmt.Println(" Starting server on port ", *port)

	http.HandleFunc("/sq", SquareRootHandler(ComputeSquareRoot))

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

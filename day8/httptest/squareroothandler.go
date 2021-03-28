package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func ComputeSquareRoot(input float64) float64 {
	return math.Sqrt(input)
}

type SqCompute func(float64) float64

func SquareRootHandler(sqComputer SqCompute) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("input")

		in, err := strconv.ParseFloat(input, 64)
		if err != nil {
			http.Error(w, "bad input", http.StatusBadRequest)
			return
		}

		out := sqComputer(in)

		fmt.Fprint(w, out)
	}
}

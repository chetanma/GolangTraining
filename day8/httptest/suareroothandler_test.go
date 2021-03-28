package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type test struct {
	name  string
	input float64
	out   float64
}

func TestSqareroot(t *testing.T) {

	tests := []test{
		{name: "test for 64", input: 64.0, out: 8.0},
		{name: "test for 9", input: 9.0, out: 3.0},
		{name: "test for 16", input: 16.0, out: 4.0},
		{name: "test for 81", input: 81.0, out: 9.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := ComputeSquareRoot(tt.input)
			if out != tt.out {
				t.Fatal("Expected ", tt.out, " but got ", out)
			}
		})
	}

}

type testSq struct {
	name         string
	input        string
	out          string
	expectedCode int
}

func TestSqarerootHandler(t *testing.T) {

	tests := []testSq{
		{name: "test for 64", input: "64.0", out: "8", expectedCode: http.StatusOK},
		{name: "test for 9", input: "9.0", out: "3", expectedCode: http.StatusOK},
		{name: "test for 16", input: "16.0", out: "4", expectedCode: http.StatusOK},
		{name: "test for 81", input: "81.0", out: "9", expectedCode: http.StatusOK},
		{name: "test for 81", input: "gar.0", out: "bad input", expectedCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqHandler := SquareRootHandler(ComputeSquareRoot)

			r := httptest.NewRequest("GET", "http://localhost/sq?input="+tt.input, nil)
			w := httptest.NewRecorder()

			//invoke the handler
			sqHandler(w, r)

			resp := w.Result()
			_, _ = ioutil.ReadAll(resp.Body)

			if resp.StatusCode != tt.expectedCode {
				t.Fatal("Expected ", tt.expectedCode, " got ", resp.StatusCode)
			}

			// if string(body) != tt.out {
			// 	t.Error("Expected ", tt.out, " got ", string(body))
			// }

		})
	}

}

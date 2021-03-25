package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductStore interface {
	GetProduct(id int) *Product
	GetProducts() []Product
}

type ProductDataStore struct {
	Data []Product
}

func NewProductsDataStore() ProductDataStore {
	data := make([]Product, 0, 100)
	data = append(data, Product{
		Name:  "P1",
		Id:    1001,
		Price: 100.0,
	})

	data = append(data, Product{
		Name:  "P2",
		Id:    1002,
		Price: 150.0,
	})

	data = append(data, Product{
		Name:  "P3",
		Id:    1003,
		Price: 300.0,
	})

	data = append(data, Product{
		Name:  "P4",
		Id:    1004,
		Price: 200.0,
	})

	data = append(data, Product{
		Name:  "P5",
		Id:    1005,
		Price: 100.0,
	})

	data = append(data, Product{
		Name:  "P6",
		Id:    1006,
		Price: 250.0,
	})

	return ProductDataStore{
		Data: data,
	}
}

func (hd *ProductDataStore) GetProduct(id int) *Product {
	for _, h := range hd.Data {
		if h.Id == id {
			return &h
		}
	}

	return nil
}

func (hd *ProductDataStore) GetProducts() []Product {
	return hd.Data
}

//Not a way as it breaks single responsibility principle
//Decoupling is broken, unit testing and mocking would be impacted
func (hd *ProductDataStore) ProductsHttpHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		//TODO return error into w
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		//TODO return error into w
		return
	}

	product := hd.GetProduct(id)
	if product == nil {
		//TODO return error into w
		return
	}

	productData, err := json.Marshal(product)
	if err != nil {
		//TODO return error into w
		return
	}

	fmt.Fprint(w, string(productData))

}

package applib

import (
	"applib/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Version() float32 {
	return 2.1
}

type CustomerFilter func(*models.Customer) bool

type CustomerLevelFilter func(models.CustomerLevel) bool

type CustomerRepo struct {
	Customers []models.Customer
}

func NewRepo() CustomerRepo {
	return CustomerRepo{
		Customers: make([]models.Customer, 0, 100),
	}
}

func LoadFromFile(file string) (*CustomerRepo, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var customers []models.Customer

	json.Unmarshal(data, &customers)

	return &CustomerRepo{
		Customers: customers,
	}, nil

}

func (repo *CustomerRepo) Add(c models.Customer) {
	repo.Customers = append(repo.Customers, c)
}

func (repo *CustomerRepo) StoreTofile(file string, f CustomerFilter) error {

	customersDataFiltered := make([]models.Customer, 0, 100)

	for _, c := range repo.Customers {
		if f(&c) {
			customersDataFiltered = append(customersDataFiltered, c)
		}
	}

	data, err := json.Marshal(customersDataFiltered)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//fmt.Println(string(data))

	ioerr := ioutil.WriteFile(file, data, 0644)
	if ioerr != nil {
		fmt.Println(ioerr)
		return ioerr
	}

	return nil
}

package main

import (
	"applib"
	"fmt"

	"applib/models"

	"rsc.io/quote"
)

func main() {
	repo, err := applib.LoadFromFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	cust, _ := models.NewCustomer("Cust333", 1456, 100, models.Gold)
	repo.Add(*cust)

	fmt.Println(repo)

}

func CustomerRepoStoreToFile() {
	repo := applib.NewRepo()

	cust, _ := models.NewCustomer("Cust1", 1500, 100, models.Gold)
	repo.Add(*cust)
	cust, _ = models.NewCustomer("Cust2", 1501, 150, models.Platinum)
	repo.Add(*cust)
	cust, _ = models.NewCustomer("Cust3", 1502, 50, models.Silver)
	repo.Add(*cust)

	repo.StoreTofile("data.json")
}

func DynamicSlice() {
	customer := make([]models.Customer, 0, 10)

	cust, _ := models.NewCustomer("Cust1", 1500, 100, models.Silver)
	customer = append(customer, *cust)

	cust, _ = models.NewCustomer("Cust2", 1600, 100, models.Silver)
	customer = append(customer, *cust)

	for _, c := range customer {
		//c is a copy of customer
		fmt.Println(c)
	}

	//How to avoid copy
	for i := range customer {
		fmt.Println(customer[i])
	}
}

func SlicingArray() {
	a := []int{1, 4, 66, 77, 88, 33}

	fmt.Println(a)
	s1 := a[1:3]
	fmt.Println(s1)

	s2 := a[3:len(a)]
	fmt.Println(s2)
}

func CustomerInit() {

	cust, err := models.NewCustomer("Cust1", 1500, 100, models.Silver)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cust)

	cust, err = models.NewCustomer("Cust1", 500, 100, models.Gold)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cust)

}

func Hello() string {
	return quote.Hello()
}

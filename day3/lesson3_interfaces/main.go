package main

import "fmt"

func main() {
	
	repo:=NewCustomerRepoDummy()
	cust, _ := NewCustomer("Cust1", 1500, 100, Gold)
	repo.AddCustomer(*cust)
	cust, _ = NewCustomer("Cust2", 1501, 150, Gold)
	repo.AddCustomer(*cust)
	cust, _ = NewCustomer("Cust3", 1502, 200, Gold)
	repo.AddCustomer(*cust)
	cust, _ = NewCustomer("Cust4", 1503, 250, Gold)
	repo.AddCustomer(*cust)
	cust, _ = NewCustomer("Cust5", 1504, 300, Gold)
	repo.AddCustomer(*cust)

	fmt.Println(repo)
	fmt.Println("customer found->",GetCustomerRepoOperation(&repo,1503))


}
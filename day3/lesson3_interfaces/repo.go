package main

// type Customer struct{

// }

type CustomerRepo interface {
	GetCustomer(id int) Customer
	//GetCustomers(id ...int) []Customer
}

//in another file
type CustomerRepositoryJsonFile struct {
	file string
}


//in some other dummy file like test
//in memory
type CustomerRepoDummy struct {
}

func GetCustomerRepoOperation(c CustomerRepo, id int) Customer{
	return c.GetCustomer(id)
}

// func GetCustomersRepoOperation( c CustomerRepo,id ..int) []Customer{

// }
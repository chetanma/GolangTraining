package main

type customerRepoDummy struct{
	Customers  []Customer
}

func NewCustomerRepoDummy() customerRepoDummy{
	return customerRepoDummy{
		Customers: make([]Customer, 0, 100),
	}
}

func (cust *customerRepoDummy) AddCustomer(c Customer) {
	cust.Customers= append(cust.Customers, c)
}

func (cust *customerRepoDummy) GetCustomer(id int) Customer{
	var t Customer
	for _,c := range cust.Customers{
		if c.ID== id{
			t=c
			break
		}
	}
	return t
}
package main

import (
	"applib"
	"fmt"

	"applib/models"

	"rsc.io/quote"
)

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func MakerAdder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	//Do practise on the following

	CustomerRepoStoreToFileWithFilters()

	AccumulatorWithClosure()
	ClosureDemo()
	//Read this: https://play.golang.org/p/2akYiBfjXA

}

func ClosureDemo() {

	fiveAdder := MakerAdder(5)
	twelveAdder := MakerAdder(12)

	out := fiveAdder(2)
	fmt.Println(out)

	out = twelveAdder(10)
	fmt.Println(out)

}

func AccumulatorWithClosure() {
	acc := adder()
	acc2 := adder()

	res := acc(5)
	res = acc(5)

	res2 := acc2(5)

	fmt.Println(res)
	fmt.Println(res2)
}

func CustomerRepoStoreToFileWithFilters() {
	repo := applib.NewRepo()
	cust, _ := models.NewCustomer("Cust1", 1500, 100, models.Gold)
	repo.Add(*cust)
	cust, _ = models.NewCustomer("Cust2", 1501, 150, models.Platinum)
	repo.Add(*cust)
	cust, _ = models.NewCustomer("Cust3", 1502, 50, models.Silver)
	repo.Add(*cust)
	cust, _ = models.NewCustomer("Cust4", 1503, 120, models.Gold)
	repo.Add(*cust)
	cust, _ = models.NewCustomer("Cust5", 1504, 150, models.Silver)
	repo.Add(*cust)

	// premiumFilter := func(c *models.Customer) bool {
	// 	if c.Level == models.Gold || c.Level == models.Platinum {
	// 		return true
	// 	}
	// 	return false
	// }

	onlySilverFilter := func(c *models.Customer) bool {
		if c.Level == models.Silver {
			return true
		}
		return false
	}

	repo.StoreTofile("data.json", onlySilverFilter)

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

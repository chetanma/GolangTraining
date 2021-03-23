package main

import "fmt"

func VariablesAndPointers() {
	var name string = "string1"
	name2 := "string2"

	fmt.Println(name)
	fmt.Println(name2)

	nameptr := &name

	var nameptr2 *string = &name2

	fmt.Printf("%p %v\n", nameptr, *nameptr)
	fmt.Printf("%p %v\n", nameptr2, *nameptr2)

	int1 := 10
	int1ptr := &int1

	fmt.Printf("%p %v\n", int1ptr, *int1ptr)
}

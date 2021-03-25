package main

import "fmt"

//should have only one method
type Squarable interface {
	Square(int) int
}

type SquareFn func(in int) int

//implement Squarable interface
func (sqF SquareFn) Square(in int) int {
	return sqF(in)
}

func DoSquare(si Squarable) {
	fmt.Println(si.Square(10))
}

func DoSquareWithFn(si SquareFn) {
	fmt.Println(si(10))
}
func FunctionalInfAdapterDemo() {
	sq := func(in int) int {
		return in * in
	}

	DoSquareWithFn(sq)

	// var sqinf Squarable = func(a int) int {
	// 	return a * a
	// }
	var sqi Squarable = SquareFn(sq)
	fmt.Println(sqi.Square(10))

	DoSquare(SquareFn(sq))
}

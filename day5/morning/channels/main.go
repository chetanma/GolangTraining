package main

import "fmt"

func main() {

	//incomple psuedo code of pipeline

	inCh := Generator(1, 2, 3, 4, 5, 8, 3)

	out1 := Doubler(inch)
	out2 := Doubler(inch)

	out := merge(out1, out2)

	for v := range out {
		fmt.Println(v)
	}

}

func Generator(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range vals {
			ch <- v
		}
	}()
	return ch
}

func Doubler(inch <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range inch {
			ch <- v * v
		}
	}()
	return ch
}

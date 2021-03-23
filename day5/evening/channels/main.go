package main

import (
	"time"
)

type Customer struct{}

func main() {

}

func Generate() <-chan int {
	ch := make(chan int)

	go func() {
		for {
			time.Sleep(time.Millisecond * 600)
			ch <- 10
		}
	}()
	return ch
}

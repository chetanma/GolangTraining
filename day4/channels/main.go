package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//SimpleChannelWorkers()
	SimpleChannelWorkers2()
}

func SimpleChannelDemo() {
	//Channels are unbuffered and buffered
	//Send and receive operations on a channel are blocking if there is nor space i.e. default
	//The sender (generator) should close the channel after use
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 10
		ch <- 11
		ch <- 12
		ch <- 13
	}()

	for out := range ch {
		fmt.Println(out)
	}

	fmt.Println("Done reading from the channel")
}

func SimpleChannelWorkers() {
	//Channels are unbuffered and buffered
	//Send and receive operations on a channel are blocking if there is nor space send blocks i.e. default and receive blocks waiting on data.
	//The sender (generator) should close the channel after use
	ch := make(chan int)

	//create a sender routine
	go func() {
		defer close(ch)
		time.Sleep(time.Second * 2)
		ch <- 10
		ch <- 11
		ch <- 12
		ch <- 13
		ch <- 14
		ch <- 15
		ch <- 16
		ch <- 17
		time.Sleep(time.Second * 2)
	}()

	var wg sync.WaitGroup

	//create many worker routines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for out := range ch {
				fmt.Println("worker number  ", n, "received : ", out)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Done reading from the channel")
}

func SimpleChannelWorkers2() {

	ch := make(chan int)

	//create a sender routine
	go func() {
		defer close(ch)
		time.Sleep(time.Second * 2)
		ch <- 10
		ch <- 11
		ch <- 12
		ch <- 13
		ch <- 14
		ch <- 15
		ch <- 16
		ch <- 17
		time.Sleep(time.Second * 2)
	}()

	w1 := worker(0, ch)
	w2 := worker(1, ch)

	<-w1
	<-w2

	//out, ok := <-w1
	//out, ok = <-w2

	fmt.Println("Done reading from the channel")
}

func worker(n int, inCh <-chan int) <-chan int {
	sigCh := make(chan int)
	go func(n int) {
		defer close(sigCh)
		for out := range inCh {
			fmt.Println("worker number  ", n, "received : ", out)
		}
	}(n)
	return sigCh
}

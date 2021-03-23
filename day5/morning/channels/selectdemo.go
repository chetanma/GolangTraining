package main

import (
	"fmt"
	"time"
)

func SelectDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		//fast producer/sender
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		defer close(ch2)
		for i := 100; i < 106; i++ {
			ch2 <- i
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	//Below code does not allow to consume fast producer(sender) in a fast way
	// for {
	// 	out1, ok1 := <-ch1
	// 	if ok1 {
	// 		fmt.Println(out1)
	// 	}

	// 	out2, ok2 := <-ch1
	// 	if ok2 {
	// 		fmt.Println(out2)
	// 	}
	// }

	//Below code will allow
	for ch1 != nil || ch2 != nil {
		select {
		case out, ok := <-ch1:
			if ok {
				fmt.Println(out)
			} else {
				ch1 = nil
			}
		case out, ok := <-ch2:
			if ok {
				fmt.Println(out)
			} else {
				ch2 = nil
			}
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

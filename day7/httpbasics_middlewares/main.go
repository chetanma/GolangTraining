package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//HttpServerAndMiddlewareDemo()

	RoutingDemo()
}

func DemoForClosureProblems() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(workerNum int) {
			defer wg.Done()
			for {
				//Problem in closure:
				//fmt.Println("Worker num : ", i)
				//Below not a problem
				fmt.Println("Worker num : ", workerNum)
				time.Sleep(time.Millisecond * 800)
			}
		}(i)
	}
	wg.Wait()
}

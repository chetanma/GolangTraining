package main

import (
	"fmt"
	"sync"
	"time"
)

func InvokeWorkersSync() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers are done!")
}

func worker(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	defer func(n int) {
		fmt.Println(" defering executing ", n)
	}(num)

	defer func(n int) {
		fmt.Println(" defering executing 2 for ", n)
	}(num)

	fmt.Println("Perform work ", num)
	time.Sleep(time.Millisecond * 100)
}

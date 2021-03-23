package main

import (
	"fmt"
	"time"
)

func Invoking_workers() {
	go worker_routine(1)
	go worker_routine(2)

	go func(indentifier int) {
		fmt.Println("In Another func in go routine with id ", indentifier)
	}(1001)

	time.Sleep(time.Millisecond * 100)
}

func worker_routine(num int) {
	fmt.Println("Worker ", num, " invoked")
}

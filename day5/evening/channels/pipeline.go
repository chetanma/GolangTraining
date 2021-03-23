package main

import (
	"fmt"
	"sync"
)

func SimpleWorkerPipeline() {
	inCh := Generator(1, 2, 3, 4, 5, 8, 3)

	out1 := Doubler(inCh)
	out2 := Doubler(inCh)
	out3 := Doubler(inCh)
	out4 := Doubler(inCh)

	out := merge(out1, out2, out3, out4)

	for v := range out {
		fmt.Println(v)
	}
}

func merge(chans ...<-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					ch <- v
				}
			}(c)
		}

		wg.Wait()

	}()

	return ch
}

func merge_s(a, b <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if ok {
					ch <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					ch <- v
				} else {
					b = nil
				}
			}
		}
	}()

	return ch

}

func Generator(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
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
			ch <- v * 2
		}
	}()
	return ch
}

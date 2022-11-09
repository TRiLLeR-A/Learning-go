package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	s := a + b
	fmt.Printf("Sum of %.2f and %.2f is %.2f\n", a, b, s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i < 5; i++ {
		go sum(10.3*float64(i), 20.1*float64(i), &wg)
	}

	wg.Wait()
}

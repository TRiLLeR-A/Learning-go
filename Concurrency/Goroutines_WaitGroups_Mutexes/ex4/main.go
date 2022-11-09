package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100; i < 150; i++ {
		go func(f int, wg *sync.WaitGroup) {
			x := math.Sqrt(float64(f))
			fmt.Printf("Square root of %.2v is %.4f\n", f, x)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

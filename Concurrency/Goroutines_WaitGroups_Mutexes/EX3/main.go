package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a float64, w *sync.WaitGroup) {
		fmt.Println(math.Sqrt(a))
		w.Done()
	}(2.2, &wg)
	wg.Wait()
}

/*
package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
		wg.Done()
	}(16.1, &wg)

	wg.Wait()
}
*/

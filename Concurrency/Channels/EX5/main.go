package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 1; i <= 100; i++ {
		go func(x int) {
			ch <- x * x
		}(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
	}
}

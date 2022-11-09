package main

import (
	"fmt"
)

func main() {
	// c := make(<-chan int)  -- It is an error
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}

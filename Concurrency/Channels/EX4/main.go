package main

import (
	"fmt"
)

func power(a int, c chan int) {
	c <- a * a
}

func main() {
	ch := make(chan int)

	for i := 1; i < 51; i++ {
		go power(i, ch)
		fmt.Println(<-ch)
	}
}

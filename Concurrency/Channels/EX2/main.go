package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(n string) {
		ch <- n
	}("Gophers!")

	value := <-ch
	fmt.Println("Value received from channel:", value)
}

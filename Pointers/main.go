package main

import "fmt"

// Ex1
func ex1() {
	x := 10.10
	fmt.Println(&x)
	ptr := &x
	fmt.Printf("%T", ptr)
	fmt.Printf("%T, %v", &ptr, *ptr)
	fmt.Println("\nSeccess1")
}

// Ex2
func ex2() {
	x, y := 10, 2
	ptrx, ptry := &x, &y
	z := *ptrx / *ptry
	fmt.Println(z, "\nSeccess2")
}

// Ex3
func ex3(a *float64, b *float64) {
	*a, *b = *b, *a
}

func main() {
	ex1()
	ex2()
	x, y := 5.5, 8.8
	ex3(&x, &y)
	fmt.Printf("%v, %v", x, y)
	fmt.Println("\nSeccess3")
}

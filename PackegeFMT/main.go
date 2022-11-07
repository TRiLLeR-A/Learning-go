package main

import "fmt"

// Coding Exercise #1
func ex1() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	fmt.Printf("%d,%f,%s,%v\n", x, y, z, score)
	fmt.Printf("z is %q\n", z)
	fmt.Printf("z is %q\n", z)

	fmt.Println("SUCCESS1")
}

// Coding Exercise #2
func ex2() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)
	fmt.Println("SUCCESS2")
}

// Coding Exercise #3
func ex3() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	fmt.Printf("y type: %T, score type: %T\n", y, score)
	fmt.Println("SUCCESS3")
	_, _ = x, z
}

func main() {
	ex1()
	ex2()
	ex3()
}

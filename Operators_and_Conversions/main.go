package main

import (
	"fmt"
	"strconv"
)

// Exercise 1
func ex1() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
	fmt.Println("SUCCESS1")
}

// Exercise 2
func ex2() {
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5
	b := 6.2 * float64(a)
	fmt.Println(b)
	fmt.Println("SUCCESS2")
}

func main() {
	ex1()
	ex2()
}

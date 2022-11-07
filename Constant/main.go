package main

import (
	"fmt"
	"math"
)

// Exercise 1
func ex1() {
	const daysWeek = 7
	const lightSpeed = 299792458
	const pi = 3.14159
	fmt.Println("Success1")

}

// Exercise 2
func ex2() {
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)
	fmt.Println("Success2")
}

// Exercise 3
func ex3() {
	const secPerDay = 60 * 60 * 24
	const daysYear = 365

	fmt.Printf("There are %d seconds in a year.\n", secPerDay*daysYear)
	fmt.Println("Success3")
}

// Exercise 4
func ex4() {
	const x int = 10

	// declaring a constant of type slice int ([]int)
	var m = []int{1: 3, 4: 5, 6: 8}
	_ = m
	fmt.Println("Success4")
}

// Exercise 5
func ex5() {
	const a = 7
	const b float64 = 5.6
	const c = a * b

	x := 8
	const xc int = 8
	_ = x

	var noIPv6 = math.Pow(2, 128)
	_ = noIPv6
	fmt.Println("Success5")
}

// Exercise 6
func ex6() {
	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Println(Jun, Jul, Aug, "Success6")
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

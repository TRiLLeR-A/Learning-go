package main

import "fmt"

func ex1() {
	// EXERCISE1
	//1
	var a int
	var b float64
	var c bool
	var d string
	// _, _, _, _ = a, b, c, d
	//2
	x := 20
	y := 15.5
	z := "Gopher!"
	// _,_,_ = x,y,z
	//3
	fmt.Println(a, b, c, d, x, y, z)
}

func ex2() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher!"
	_, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error
	fmt.Print("SUCCESS2")
}

func ex3() {
	var a float64 = 7.1
	x, y := true, 3.7
	a, x = 5.5, false
	_, _, _ = a, x, y
	fmt.Println("SUCCESS3")
}

func ex4() {
	version := "3.1"

	name := "Golang"
	fmt.Println(name, "SUCCES4")
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

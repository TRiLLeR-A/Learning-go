package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Exersice1
func cube(a float64) float64 {
	return math.Pow(a, 3)
}

// Exersice2
func f1(a uint) (int, int) {
	f := 1
	s := 0
	for i := int(a); i > 0; i-- {
		f *= i
		s += i
	}
	return f, s
}

// Exersice3
func myfunc(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("Cannot conver %q to in.", a)
		os.Exit(1)
	}
	i2, _ := strconv.Atoi(a + a)
	i3, _ := strconv.Atoi(a + a + a)
	return i + i2 + i3
}

// Exersice4
func sum(a ...int) int {
	s := 0
	for _, j := range a {
		s += j
	}
	return s
}

// Exersice5
func sum1(a ...int) (s int) {
	for _, j := range a {
		s += j
	}
	return
}

// Exersice6
func searchItem(a []string, b string) bool {
	for _, j := range a {
		if b == j {
			return true
		}
	}
	return false
}

// Exersice7
func searchItem1(a []string, b string) bool {
	for _, v := range a {
		if strings.EqualFold(v, b) {
			return true
		}
	}
	return false
}

// Exersice8
func print(msg string) {
	fmt.Println(msg)
}

// Exersice9
func ex9(a int) {
	fmt.Println(a)
}

func main() {
	fmt.Println(cube(3))
	fmt.Println("Seccess1")

	f, s := f1(4)
	fmt.Println("f:", f, "s:", s)
	f, s = f1(10)
	fmt.Println("f:", f, "s:", s)
	fmt.Println("Seccess2")

	fmt.Println(myfunc("9"), "\nSeccess3")

	fmt.Println(sum(14, 15, 16, 7), "\nSeccess4")

	fmt.Println(sum1(1, 2, 3, 4, 5), "\nSeccess5")

	animals := []string{"lion", "tiger", "bear"}

	fmt.Println(searchItem(animals, "bear"), "\nSeccess6")

	fmt.Println(searchItem1(animals, "LION"), "\nSeccess6")

	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
	fmt.Println("\nSeccess8")

	a := ex9
	fmt.Printf("%T \n", a)
	a(23)
	fmt.Println("\nSeccess9")
}

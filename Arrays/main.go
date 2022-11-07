package main

import "fmt"

// Exercise 1
func ex1() {
	var cities [2]string
	fmt.Printf("%#v \n", cities)

	grades := [3]float64{4.5, 9.7}
	fmt.Printf("%#v \n", grades)

	taskDone := [...]bool{}
	fmt.Printf("%#v \n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%v, ", cities[i])
	}
	fmt.Println()
	for i, j := range grades {
		fmt.Printf("%v:%v, ", i, j)
	}
	fmt.Println()
	fmt.Println("SUCCESS1")
}

// Exercise 2
func ex2() {
	nums := [...]int{30, -1, -6, 90, -6}
	var counts int
	for _, j := range nums {
		if j > 0 {
			counts++
		}
	}
	fmt.Println(counts, "\n SUCCESS2")
}

// Exercise 3
func ex3() {
	myArray := [3]float64{1.2, 5.6}
	myArray[2] = 6
	a := 10
	myArray[0] = float64(a)
	myArray[1] = 10.10
	fmt.Println(myArray)
	fmt.Println("SUCCESS3")
}

func main() {
	ex1()
	ex2()
	ex3()
}

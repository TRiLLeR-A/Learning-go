package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Excersice1
func ex1() {
	cities := []string{"A", "B", "C"}
	for i, j := range cities {
		fmt.Printf("%v:%v ", i, j)
	}
	fmt.Println()
	fmt.Println("Success1")
}

// Excersice2
func ex2() {
	mySlice := []float64{1.2, 5.6}

	// mySlice[2] = 6 //	// ERROR -> index out of range [2] with length 2

	a := float64(10)
	mySlice[0] = a

	// mySlice[3] = 10.10 	// ERROR -> index out of range [3] with length 2

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
	fmt.Println()
	fmt.Println("Success2")
}

// Excersice3
func ex3() {
	var nums = []float64{1.1, 2.1, 3.1}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)
	var n = []float64{3.3, 4.4}
	nums = append(nums, n...)
	fmt.Println(nums)
	fmt.Println()
	fmt.Println("Success3")
}

// Excersice4
func ex4() {
	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)
	fmt.Println("Success4")
}

// Excersice5
func ex5() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, j := range nums[2 : len(nums)-2] {
		fmt.Printf("%d, ", j)
		sum += j
	}
	fmt.Println()
	fmt.Printf("sum = %d ", sum)
	fmt.Println()
	fmt.Println("Success5")
}

// Excersice6
func ex6() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)
	yourFriends[0] = "Dan"
	fmt.Println(friends, yourFriends)
	fmt.Println("Success6")
}

// Excersice7
func ex7() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	fr2 := []string{}
	fr2 = append(fr2, friends...)
	fmt.Println()
	fr2[0] = "Antuan"
	fmt.Println(friends, fr2)
	fmt.Println("Success7")
}

// Excersice8
func ex8() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(newYears, years[:3]...)
	newYears = append(newYears, years[len(years)-3:]...)
	// Second_way
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Printf("%v \n%v\n", years, newYears)
	fmt.Println()
	fmt.Println("Success8")
}
func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
}

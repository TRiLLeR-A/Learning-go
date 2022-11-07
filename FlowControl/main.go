package main

import "fmt"

// EXERCISE1
func ex1() {
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		} else {
			fmt.Printf("%v, ", i)
		}
	}
	fmt.Println()
	fmt.Println("Success1")
}

// EXERCISE2
func ex2() {
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%v, ", i)
		}
	}
	fmt.Println()
	fmt.Println("Success2")
}

// EXERCISE3
func ex3() {
	for i, j := 1, 0; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%v, ", i)
		j++
		if j > 2 {
			break
		}
	}
	fmt.Println()
	fmt.Println("Success3")
}

// EXERCISE4
func ex4() {
	for i := 1; i <= 500; i++ {
		if i%7 != 0 && i%5 != 0 {
			continue
		}
		fmt.Printf("%v, ", i)

	}
	fmt.Println()
	fmt.Println("Success4")
}

// EXERCISE5
func ex5() {
	birthYear, currentYear := 1998, 2022

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
	fmt.Println("Success5")
}

// EXERCISE6
func ex6() {
	age := 23
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
	fmt.Printf("Success6")
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

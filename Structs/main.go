package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

// Ex1
func ex1() {
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}
	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
	fmt.Println("Seccess1")
}

// Ex2
func ex2() {
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	for i, j := range me.colors {
		fmt.Printf("%v: %v\n", i, j)
	}
	fmt.Println("Seccess2")
}

// Ex3
func ex3() {
	type grades struct {
		grade  int
		course string
	}
	type person struct {
		name   string
		age    int
		colors []string
		gr     grades
	}
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}, gr: grades{1, "A"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}, gr: grades{1, "A"}}
	me.gr.grade = 98
	me.gr.course = "Golang"
	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
	fmt.Println("Seccess3")
}

func main() {
	ex1()
	ex2()
	ex3()
}

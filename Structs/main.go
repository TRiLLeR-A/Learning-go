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
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}
	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
	fmt.Println("Seccess2")
}

// Ex3
func ex3() {
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}
	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
	fmt.Println("Seccess3")
}

// Ex4
func ex4() {
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}
	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
	fmt.Println("Seccess4")
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

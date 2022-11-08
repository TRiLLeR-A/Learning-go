package main

import "fmt"

// ////////////////////////////////////////////////////////////////////////////////////////
// Ex1
type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func ex1() {
	var vehicle vehicle = car{licenceNo: "123", brand: "Mers"}
	fmt.Println(vehicle.License(), vehicle.Name(), "Seccess1")
}

// /////////////////////////////////////////////////////////////////////////////
// Ex2 - Type Assertion
func ex2() {
	var emty interface{}

	emty = 1
	fmt.Println(emty)

	emty = 1.2
	fmt.Println(emty)

	emty = []int{1, 2, 4}
	fmt.Println(emty)

	emty = append(emty.([]int), 3)
	fmt.Println(emty)

	fmt.Println("Seccess2")
}

// /////////////////////////////////////////////////////////////////////////////
// Ex3
type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func ex3() {
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
	fmt.Println("Seccess3")
}

// /////////////////////////////////////////////////////////////////////////////
func main() {
	ex1()
	ex2()
	ex3()
}

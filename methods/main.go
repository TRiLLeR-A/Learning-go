package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

// Ex1
func ex1() {
	var m money = 2.4
	m.print()
	fmt.Println("Seccess1")
}

type money1 float64

func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

// //////////////////////////////////////////////////////////////////////////
// Ex2
func ex2() {
	var m money = 2.4
	a := m.printStr()
	fmt.Println(a, "Seccess2")
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * 0.9
}

// //////////////////////////////////////////////////////////////////////////
// Ex3
func ex3() {
	bestBook := book{title: "The Trial", price: 10}
	vat := bestBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)
	fmt.Println("Seccess3")
}

// /////////////////////////////////////////////////////////////////////////////
type book3 struct {
	title string
	price float64
}

// method for book type
func (b *book3) changePrice(p float64) {
	b.price = p
}

// Ex4
func ex4() {
	bestBook := book3{title: "The Trial by Kafka", price: 9.9}
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
	fmt.Println("Seccess4")
}

// //////////////////////////////////////////////////////////////////////////
func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcdefghijkl"
	fmt.Println(s1[2:5]) //
	fmt.Println("\n" + strings.Repeat("#", 20))
	fmt.Println((s1[0:2]))
}

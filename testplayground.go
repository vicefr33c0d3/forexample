package main

import (
	"fmt"
)
var y int
var z = "Sample test for learning"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	z := 0
	fmt.Println(z)
}

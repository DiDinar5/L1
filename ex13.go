package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	// Меняем местами a и b с использованием XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a = %d, b = %d\n", a, b)
}

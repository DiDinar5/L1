package main

import (
	"fmt"
)

func main() {
	variables := []interface{}{42, "Hello", true, make(chan int)}

	for _, v := range variables {
		switch v.(type) {
		case int:
			fmt.Println(v, "is an int.")
		case string:
			fmt.Println(v, "is a string.")
		case bool:
			fmt.Println(v, "is a bool.")
		case chan int:
			fmt.Println(v, "is a channel of int.")
		default:
			fmt.Println(v, "is of a different type.")
		}
	}
}

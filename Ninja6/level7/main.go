package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Its an INT: ", i.(int))
	case string:
		fmt.Println("Its an String: ", i.(string))
	default:
		fmt.Println("It is not clear")
	}
}

func main() {

	findType("Hello")
	findType(64)
	findType(34.2)
}

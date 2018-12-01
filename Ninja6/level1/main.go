package main

import "fmt"

func foo() int {
	return 42
}
func bar() (int, string) {
	return 41, "Hello World"
}

func main() {

	fmt.Println(foo())
	fmt.Println(bar())
}

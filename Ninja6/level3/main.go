package main

import (
	"fmt"
)

func foo() {
	fmt.Print("is the first one")
}

func main() {
	defer foo()
	fmt.Print("will not be the last one ")
}

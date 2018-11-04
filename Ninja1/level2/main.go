package main

import (
	"fmt"
)

func main() {
	var a int = 43
	var b string = "I hob di liab"
	var c bool = false

	fmt.Println(a, b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

}

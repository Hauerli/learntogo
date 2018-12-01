package main

import "fmt"

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
func bar(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func main() {
	x := []int{22, 23, 24}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

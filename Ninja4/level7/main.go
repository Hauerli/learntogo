package main

import (
	"fmt"
)

func main() {

	xxs := [][]string{
		[]string{"James", "Bond"},
		[]string{"Money", "Penny"},
	}
	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Println(i, "\t", v)
	}
}

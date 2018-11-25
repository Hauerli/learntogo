package main

import (
	"fmt"
)

func main() {
	x := make([]string, 5, 5)
	x = []string{"Alabama", "Alaska", "Arizona", "Arkansa", "California"}
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

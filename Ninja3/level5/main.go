package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Println("When", i, "is divided by 4 the modulo is ", i%4)
	}

}

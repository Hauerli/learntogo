package main

import (
	"fmt"
)

func main() {

	emp := struct {
		first, last string
		age         int
	}{
		"Maximilian", "Hauer", 30,
	}

	fmt.Println(emp)

}

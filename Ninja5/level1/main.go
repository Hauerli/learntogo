package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	favIce []string
}

func main() {
	p1 := person{
		first:  "Max",
		last:   "Hauer",
		favIce: []string{"Schoko", "Vanilla", "Stracciatella"},
	}
	p2 := person{
		first:  "Stephanie",
		last:   "Hauer",
		favIce: []string{"Haselnuss", "Erdbeer", "Cookie"},
	}

	fmt.Println(p1.first, p1.last, p1.favIce)
	fmt.Println(p2)
}

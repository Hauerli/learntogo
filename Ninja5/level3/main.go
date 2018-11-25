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

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for _, v1 := range m {
		fmt.Println(v1.first, " ", v1.last)
		for _, v2 := range v1.favIce {
			fmt.Println(v2)
		}
		fmt.Print("\n")
	}

}

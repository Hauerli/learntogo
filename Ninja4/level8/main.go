package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"bond_james":      []string{"Shaken", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}
	for k, v := range m {
		fmt.Println("This is the Record for ", k)
		for i, v2 := range v {
			fmt.Println("The Index is", i, "\t this is the Value", v2)
		}
	}
}

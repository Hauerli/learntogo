package main

import "fmt"

func main() {

	x := "James"
	switch {
	case x == "James":
		fmt.Println("Hi my Name is James")
	case x == "Penny":
		fmt.Println("Hi my Name is Moneypenny")
	}

}

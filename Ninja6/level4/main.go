package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi my name is ", p.first, " ", p.last, "and ", p.age, " years old")
}

func main() {
	p1 := person{
		first: "Maximilian",
		last:  "Hauer",
		age:   30}
	p1.speak()
}

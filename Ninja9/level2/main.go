package main

import "fmt"

type person struct {
	name string
}

type human interface {
	myname() string
}

func (p *person) myname() string {
	return p.name
}

func tellName(h human) {
	fmt.Println("Hi my name is ", h.myname())
}

func main() {

	p1 := person{name: "Hugo"}
	p2 := person{name: "Francine"}

	tellName(&p1)
	tellName(&p2)

}

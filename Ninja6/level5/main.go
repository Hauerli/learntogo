package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(sh shape) {
	fmt.Println("The shape is ", sh.area())
}

func main() {
	circ := circle{radius: 24.22}
	squ := square{length: 34.33}
	info(circ)
	info(squ)
}

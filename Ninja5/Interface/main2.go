package main

import (
	"fmt"
)

type blln struct {
	monthrate int
	duration  int
	finalrate int
}
type loan struct {
	monthrate int
	duration  int
}
type leas struct {
	monthrate int
	duration  int
	endcosts  int
}
type CostCalculator interface {
	calculateCost() int
	contrDetails()
}

func contrDetails(c CostCalculator) {
	switch c.(type) {
	case loan:
		fmt.Println("The Contracdetails for this Loan Contract are:", "\nMonthly rate: ", c.(loan).monthrate, "\nDuration: ", l.(loan).duration)
	case blln:
		fmt.Println("The Contracdetails for this Balloon Contract are:", "\nMonthly rate: ", c.(blln).monthrate, "\nFinal rate: ", c.(blln).finalrate, "\nDuration: ", c.(blln).duration)
	}
}

func (b blln) calculateCost() int {
	return (b.duration * b.monthrate) + b.finalrate
}
func (l loan) calculateCost() int {
	return l.duration * l.monthrate
}

func calcSum(costs []CostCalculator) {

	for _, v := range costs {
		v.contrDetails()
		sum := 0
		sum += v.calculateCost()
		fmt.Println("The Contractcost are: ", sum)
		fmt.Println()
	}

}

func main() {

	contr1 := blln{134, 48, 3000}
	contr2 := loan{250, 36}
	contracts := []CostCalculator{contr1, contr2}
	calcSum(contracts)
}

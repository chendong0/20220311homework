package main

import "fmt"

type Employee struct {
	name string
	salary int
	currency string
}
/*
displaySalary() method has Employee as the receiver type
*/

func (e *Employee) displaySalary()  {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)

}

func main() {
	emp1 := Employee{
		name : "DD",
		salary: 5000,
		currency: "$",
	}
	emp1.displaySalary()
}

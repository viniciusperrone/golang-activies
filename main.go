package main

import "fmt"

func main() {
	salary := 850.00
	var salaryPlusBonus float64

	salaryPlusBonus = salary

	if salary < 1000 {
		salaryPlusBonus = (salaryPlusBonus + 100)
	}

	fmt.Println("Salary: ", salaryPlusBonus)
}

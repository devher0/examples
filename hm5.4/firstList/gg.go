package main

import "fmt"

func main() {

	fmt.Println("Введите сумму дохода")
	var income float64
	fmt.Scan(&income)

	var tax float64
	if income > 10000 {
		tax = 10000 * 0.13
		if income > 50000 {
			tax += (50000-10000)*0.2 + (income-50000)*0.3
		} else {
			tax += (income - 10000) * 0.2
		}
	} else {
		tax = income * 0.13
	}
	fmt.Println(tax)
}

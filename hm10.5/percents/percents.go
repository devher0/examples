package main

import (
	"fmt"
	"math"
)

func main() {
	var deposit float64
	var percentage float64
	var horizont int
	var bankProfit float64

	fmt.Print("Пожалуйста, введите сумму депозита: ")
	fmt.Scan(&deposit)

	fmt.Print("Пожалуйста, введите сумму процентную ставку: ")
	fmt.Scan(&percentage)

	fmt.Print("Пожалуйста, введите количество лет: ")
	fmt.Scan(&horizont)

	months := horizont * 12
	percentage /= 100

	for i := 0; i < months; i++ {
		deposit += deposit * percentage
		_, precision := math.Modf(deposit * 100)
		precision /= 100
		deposit = math.Floor(deposit*100) / 100
		bankProfit += precision
	}

	fmt.Printf("%v, %v", deposit, bankProfit)
}

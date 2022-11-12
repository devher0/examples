package main

import (
	"fmt"
	"math"
)

func main() {
	var digits uint64
	var x float64

	fmt.Print("Пожалуйста, введите x: ")
	fmt.Scan(&x)

	fmt.Print("Пожалуйста, введите точность (знаков после запятой): ")
	fmt.Scan(&digits)

	epsilon := 1 / math.Pow(10, float64(digits))

	fmt.Printf("exp = %v\n", exp(x, 0, epsilon))
}

func exp(x float64, n uint64, precision float64) float64 {
	current := math.Pow(x, float64(n)) / float64(fact(n))
	if current < precision {
		return current
	}
	return current + exp(x, n+1, precision)
}

func fact(num uint64) uint64 {
	if num <= 1 {
		return 1
	}
	return num * fact(num-1)
}

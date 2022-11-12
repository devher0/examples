package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Калькулятор квадратных уравнений")
	fmt.Println("Введите коэффиценты квадратного уровнения")

	fmt.Println("Введите коэффицент - a ,где а не равно нулю !")
	var a float64
	fmt.Scan(&a)

	fmt.Println("Введите коэффицент - b")
	var b float64
	fmt.Scan(&b)

	fmt.Println("Введите коэффицент - c")
	var c float64
	fmt.Scan(&c)

	D := math.Pow(b, 2) - 4*a*c

	if D > 0 {
		fmt.Println("Так как дискриминант больше нуля то, квадратное уравнение имеет два действительных корня")

		xOne := (-b - math.Sqrt(D)) / (2 * a)
		xTwo := (-b + math.Sqrt(D)) / (2 * a)

		fmt.Println("Первый корень равен", xOne)
		fmt.Println("Второй корень равен", xTwo)
	} else if D == 0 {
		fmt.Println("Так как дискриминант равен нулю то, квадратное уравнение имеет один действительный корень")

		xOne := (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Корень равен", xOne)
	} else {
		fmt.Println("Корней нет в данном уровнении!")
	}

}

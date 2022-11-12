package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Минимальный тип данных")
	fmt.Println("=====================")

	num1 := 0
	num2 := 0
	var sum int32

	fmt.Println("Введите первое число")
	fmt.Println("Число должно вмещаться в тип int16!")
	fmt.Scan(&num1)

	fmt.Println("Введите второе число")
	fmt.Println("Число должно вмещаться в тип int16!")
	fmt.Scan(&num2)

	sum = int32(num1) * int32(num2)

	fmt.Println(sum)

	switch {
	case sum <= math.MaxUint8 && sum >= 0:
		fmt.Println("результат умножения этих чисел можно сохранить в минимальный тип данных uint8")
	case sum <= math.MaxInt8 && sum >= math.MinInt8:
		fmt.Println("результат умножения этих чисел можно сохранить в минимальный тип данных int8")
	case sum <= math.MaxUint16 && sum >= 0:
		fmt.Println("результат умножения этих чисел можно сохранить в минимальный тип данных uint16")
	case sum <= math.MaxInt16 && sum >= math.MinInt16:
		fmt.Println("результат умножения этих чисел можно сохранить в минимальный тип данных int16")
	case sum <= math.MaxInt32 && sum >= 0:
		fmt.Println("результат умножения этих чисел можно сохранить в минимальный тип данных uint32")
	default:
		fmt.Println("результат умножения этих чисел можно сохранить в минимальный тип данных int32")
	}

}

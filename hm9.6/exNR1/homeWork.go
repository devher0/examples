package main

import (
	"fmt"
	"math"
)

func main() {

	sumUint8 := 0
	sumUint16 := 0

	for cicle := 0; cicle <= math.MaxUint32; cicle++ {

		if cicle%math.MaxUint8 == 0 {
			sumUint8++
		}

		if cicle%math.MaxUint16 == 0 {
			sumUint16++
		}

	}

	fmt.Println("переполнений чисел типа uint8 в диапазоне от 0 до uint32 равна:", sumUint8)
	fmt.Println("переполнений чисел типа uint16 в диапазоне от 0 до uint32 равна:", sumUint16)

}

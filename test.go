package main

import (
	"fmt"
	"math"
)

func main() {

	var limUint8 int8 = 0
	sumUint8 := 0
	limUint16 := 0
	sumUint16 := 0

	for cicle := 0;  cicle <= math.MaxUint32; cicle++ {

		limUint8 = limUint8 + 1

		if limUint8 = math.MaxUint8 {
			sumUint8 = sumUint8 + 1
			limUint8 == 0 
		}
	}

	fmt.Println("переполнений чисел типа uint8 в диапазоне от 0 до uint32 равна:", sumUint8)
	fmt.Println("переполнений чисел типа uint16 в диапазоне от 0 до uint32 равна:", sumUint16)

}

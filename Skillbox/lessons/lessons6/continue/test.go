package main

import (
	"fmt"
)

func main() {

	number := 0
	sum := 0

	for number < 40 {

		number = number + 1

		if number%4 != 0 {
			continue
		}

		sum = sum + number

		fmt.Println(sum)
	}

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ёлочка")

	fmt.Println("Введите высоту ёлочки")
	var height int
	fmt.Scan(&height)

	for i := 1; i <= height; i++ {

		for j := 0; j < height*2-1; j++ {
			if j+i+1 <= height || j-i+1 >= height {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}

		fmt.Printf("\n")

	}

}

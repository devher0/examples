package main

import (
	"fmt"
)

func main() {

	fmt.Println("шахматная доска")

	fmt.Println("Введите ширину доски")
	var width int
	fmt.Scan(&width)

	fmt.Println("Введите высоту доски")
	var height int
	fmt.Scan(&height)

	for i := 1; i <= height; i++ {

		if i%2 != 0 {

			for j := 1; j <= width; j++ {
				if i%2 != 0 && j%2 != 0 {
					fmt.Printf("*")
				} else if j%2 == 0 {
					fmt.Printf(" ")
				}
			}

			fmt.Printf(" \n")

		} else if i%2 == 0 {

			for j := 1; j <= width; j++ {
				if i%2 == 0 && j%2 == 0 {
					fmt.Printf("*")
				} else if j%2 != 0 {
					fmt.Printf(" ")
				}

			}

			fmt.Printf(" \n")
		}
	}
}

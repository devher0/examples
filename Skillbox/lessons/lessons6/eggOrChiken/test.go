package main

import (
	"fmt"
)

func main() {

	fmt.Println("Что появилось раньше ?")
	fmt.Println("Куривца или яйцо ?")

	for {

		answer := ""
		fmt.Scan(&answer)

		if answer == "Надоело" {
			break
		}

		fmt.Println("Не правильно, Яйцо =)")
	}

	fmt.Println("Ха-ха, я победил ! =P")
}

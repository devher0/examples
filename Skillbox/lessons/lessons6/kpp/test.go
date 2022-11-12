package main

import (
	"fmt"
)

func main() {

	i := 0

	kpp := -10
	stepRight := 10
	stepLeft := -10

	for i < 100 {
		i = i + 1
		if kpp < 0 {
			for {
				kpp = kpp + 1
				fmt.Println("Шаг", kpp)
				if kpp == 0 {
					fmt.Println("Часовой на КПП")
				} else if kpp == stepRight {
					break
				}
			}
		} else if kpp > 0 {
			for {
				kpp = kpp - 1
				fmt.Println("Шаг", kpp)
				if kpp == 0 {
					fmt.Println("Часовой на КПП")
				} else if kpp == stepLeft {
					break
				}
			}
		}
	}

}

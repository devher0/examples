package main

import (
	"fmt"
)

func main() {

	var i uint16 = 10
	for ; i != (1<<16 - 2); i-- {
		if i == 0 {
			fmt.Println("Колла закончилась ... ")
			fmt.Println("Но мы начнем с начала ")
			continue
		}
		fmt.Printf("осталось бутылок: %d \n", i)
	}
}

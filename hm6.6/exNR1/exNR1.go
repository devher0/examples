package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите число")
	var num int
	fmt.Scan(&num)

	for i := 0; i <= num; i = i + 1 {

		fmt.Println(i)
	}

}

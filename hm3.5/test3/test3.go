package main

import (
	"fmt"
)

func main() {

	fmt.Println("Обмен местами значений переменных")

	a := 42
	b := 153

	fmt.Println("a :", a)
	fmt.Println("b :", b)

	a, b = 153, 42

	fmt.Println("a :", a)
	fmt.Println("b :", b)

}

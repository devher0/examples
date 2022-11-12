package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет данная программ выщитывает рост васшего бамбука")
	fmt.Println("Введите высоту саженца")
	var startHeight int
	fmt.Scan(&startHeight)
	fmt.Println("Введите скорость роста саженца")
	var growth int
	fmt.Scan(&growth)
	fmt.Println("Введите скорость поедания вредителей")
	var losses int
	fmt.Scan(&losses)
	fmt.Println("Введите сколько дней бамбук будет рости")
	var hours24 int
	fmt.Scan(&hours24)

	finalHeight := startHeight + (growth-losses)*hours24

	fmt.Println("Ваш бамбук после", hours24, "дней, выростит на", finalHeight, "сантиметров")

}

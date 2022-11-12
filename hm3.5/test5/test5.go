package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет данная программ выщитывает созревание бамбука")
	fmt.Println("Введите высоту саженца")
	var startHeight int
	fmt.Scan(&startHeight)
	fmt.Println("Введите скорость роста саженца")
	var growth int
	fmt.Scan(&growth)
	fmt.Println("Введите скорость поедания вредителей")
	var losses int
	fmt.Scan(&losses)
	fmt.Println("Введите введите целевую высоту взрослого бамбука")
	var finalGrowth int
	fmt.Scan(&finalGrowth)

	term := (finalGrowth - startHeight) / (growth - losses)

	fmt.Println("Бамбук достигнет нужного роста, в", finalGrowth, "см, через", term, "дней")

}

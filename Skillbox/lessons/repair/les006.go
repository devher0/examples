package main

import (
	"fmt"
)

func main() {

	totalPrice := 4000000
	entrance := 10
	totalApartments := entrance * 40
	endPrice := totalPrice / totalApartments

	fmt.Println(totalPrice, "(сумма ремонта)", "/", totalApartments, "(общее количество квартир)", "=", endPrice, "(сумма ремонта на одну квартиру)")

}

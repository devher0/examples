package main

import (
	"fmt"
)

func main() {

	startHeight := 100.0
	growth := 50.0
	losses := 20.0
	need4Sell := 300.0

	fmt.Println("Задача - 1")
	finalGrowth := growth*2 + growth/2
	finalLosses := losses * 2
	finalHeight := startHeight + finalGrowth - finalLosses
	fmt.Println("Высота бамбука на третий день равна", finalHeight)

	fmt.Println("Задача - 2")
	hours24 := growth - losses
	andGrowth := (need4Sell - startHeight) / hours24

	fmt.Printf("%.2f", andGrowth)
	fmt.Println(" дней бамбук должен расти чтоб его можно было продать.")

	/*
		300 = 100start + [50growth - 20losses } 1day
		200/30



	*/

	/*
		curentHeight := startHeight
		curentHeight += growth
		curentHeight -= losses
		curentHeight += growth
		curentHeight -= losses
		curentHeight += growth / 2
	*/

	/*
		можно сократить данное решение

		finalGrowth = growth * 2 + growth / 2
		finalLosses = losses * 2
		finalHeight = startHeight + finalGrowth - finalLosses
	*/
}

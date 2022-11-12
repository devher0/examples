package main

import (
	"fmt"
)

func main() {

	mirrorTicket := 0

	for ticket := 100000; ticket <= 999999; ticket++ {

		var reversedDigit int
		searchNumber := ticket

		firstPart := searchNumber / 1000
		firstDigit := firstPart * 1000
		secondPart := searchNumber - firstDigit

		for firstPart > 0 {
			lastDigit := firstPart % 10
			reversedDigit = reversedDigit*10 + lastDigit

			firstPart = firstPart / 10

		}

		if secondPart == reversedDigit {
			mirrorTicket = mirrorTicket + 1
		}
	}
	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров 	от 100000 до 999999:", mirrorTicket)
}

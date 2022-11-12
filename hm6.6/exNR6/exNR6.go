package main

import (
	"fmt"
)

func main() {

	fmt.Println("Это исмулятор лифта")
	// fmt.Println("Укажите на каком этаже находятся люди")
	// fmt.Println("-------------------")

	// var floor int
	// fmt.Scan(&flour)

	totalFloors := 24
	floor := 1

	pasStatOne := 10
	pasStatTwo := 7
	pasStatThree := 4

	quantityPasOne := 1
	quantityPasTwo := 1
	quantityPasThree := 1

	fullCap := 2
	capOfLift := 0

	floorOne := 0

	for i := 0; i < 4; i++ {

		if floor == 1 {
			fmt.Println("Этаж №", floor)

			for {

				floor = floor + 1
				fmt.Println("Этаж №", floor)

				if floor != totalFloors {
					continue
				} else if floor == totalFloors {
					break
				}

			}
		} else if floor == totalFloors {

			for {

				floor = floor - 1
				fmt.Println("Этаж №", floor)

				if floor != 1 && capOfLift == fullCap {
					continue
				} else if floor == 1 {
					floorOne = floorOne + capOfLift
					fmt.Println("Вышло", capOfLift, "чел.")
					capOfLift = 0
					fmt.Println("На этаже", floorOne, "чел.")
					break
				} else if floor == pasStatOne && quantityPasOne > 0 {
					capOfLift = capOfLift + quantityPasOne
					fmt.Println("Остановка, подборали, + ", quantityPasOne, "чел.")
					quantityPasOne = 0
					continue
				} else if floor == pasStatTwo && quantityPasTwo > 0 {
					capOfLift = capOfLift + quantityPasTwo
					fmt.Println("Остановка, подборали, + ", quantityPasTwo, "чел.")
					quantityPasTwo = 0
					continue
				} else if floor == pasStatThree && quantityPasThree > 0 {
					capOfLift = capOfLift + quantityPasThree
					fmt.Println("Остановка, подборали, + ", quantityPasThree, "чел.")
					quantityPasThree = 0
					continue
				}

			}
		}

	}

}

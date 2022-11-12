package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет, давай поиграем!")
	fmt.Println("Загадай целое число от 1 до 10 (включительно)")
	fmt.Println("А я попробую его отгадать !")
	fmt.Println("Если я отгадаю введи 'yes', а если число будет больше или меньше")
	fmt.Println("То введиите 'more' или 'less'")
	fmt.Println("Если число не верное, введите 'no")
	fmt.Println("Ну поехали !")
	fmt.Println("Ваше чило 10 ?")

	var try string
	fmt.Scan(&try)

	answer := try == "no"

	if answer == true {
		fmt.Println("ваше число больше или меньше ?")

		var try string
		fmt.Scan(&try)

		answer := try == "less"

		if answer == true {
			secondAnswer := 10 / 2
			fmt.Println("Ваше число", secondAnswer, "?")

			var try string
			fmt.Scan(&try)

			answer := try == "no"

			if answer == true {
				fmt.Println("ваше число больше или меньше ?")

				var try string
				fmt.Scan(&try)

				answer := try == "less"

				if answer == true { // 1234
					thirdAnswer := secondAnswer / 2
					fmt.Println("Ваше число", thirdAnswer, "?")

					var try string
					fmt.Scan(&try)

					answer := try == "no"

					if answer == true {
						fmt.Println("ваше число больше или меньше ?")

						var try string
						fmt.Scan(&try)

						answer := try == "less"

						if answer == true {
							fmt.Println("Ваше число", 1, "!")
						} else if answer == false { // 34
							fmt.Println("Ваше число", 3, "?")

							var try string
							fmt.Scan(&try)

							answer := try == "no"

							if answer == true {
								fmt.Println("Ваше число", 4, "!")
							} else if answer == false {
								fmt.Println("Да, я знал что 3 !")
							}
						}
					} else if answer == false {
						fmt.Println("Да я знал что 2 !")
					}
				} else if answer == false { //6789
					thirdAnswer := secondAnswer + 2 //7
					fmt.Println("Ваше число", thirdAnswer, "?")

					var try string
					fmt.Scan(&try)

					answer := try == "no"

					if answer == true {
						fmt.Println("ваше число больше или меньше ?")

						var try string
						fmt.Scan(&try)

						answer := try == "less"

						if answer == true {
							fmt.Println("Ваше число", 6, "!")
						} else if answer == false { // 89
							fmt.Println("Ваше число", 9, "?")

							var try string
							fmt.Scan(&try)

							answer := try == "no"

							if answer == true {
								fmt.Println("Ваше чило", 8, "!")
							} else if answer == false {
								fmt.Println("Да, я знал что 9 !")
							}
						}
					} else if answer == false {
						fmt.Println("Да, я знал что 7 !")
					}
				}

			} else if answer == false {
				fmt.Println("Да, я знал что 5 !")
			}
		} else if answer == false {
			fmt.Println("Не возможно !")
			fmt.Println("Число должно быть не больше 10 !")
			fmt.Println("Перезапустите программу !")
		}

	} else if answer == false {
		fmt.Println("Да я знал что 10 !")
	}
}

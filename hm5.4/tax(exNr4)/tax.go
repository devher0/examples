package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите сумму оплаты")
	var payment int
	fmt.Scan(&payment)

	fmt.Println("Введите номеналы 3х монет")
	var coinOne, coinTwo, coinThree int
	fmt.Scan(&coinOne)
	fmt.Scan(&coinTwo)
	fmt.Scan(&coinThree)

	if coinOne <= coinTwo && coinOne <= coinThree {

		if payment-coinOne == 0 {
			fmt.Println("Достаточно одной монеты для оплаты")
			fmt.Println("Сдача не требуется")
		} else if payment-(coinTwo+coinThree) == 0 {
			fmt.Println("Достаточно двух монет для оплаты")
			fmt.Println("Сдача не требуется")
		} else if payment-coinOne < 0 {
			fmt.Println("Достаточно одной монеты для оплаты")
			rest := payment - coinOne
			rest = -rest
			fmt.Println("требуется сдача :", rest)
		} else if payment-(coinTwo+coinThree) < 0 {
			fmt.Println("Достаточно двух монет для оплаты")
			rest := payment - (coinTwo + coinThree)
			rest = -rest
			fmt.Println("требуется сдача :", rest)
		} else if payment-coinOne > 0 && payment-(coinTwo+coinThree) > 0 {
			if payment-(coinOne+coinTwo) == 0 || payment-(coinOne+coinThree) == 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				fmt.Println("Сдача не требуется")
			} else if payment-(coinOne+coinTwo) < 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				rest := payment - (coinOne + coinTwo)
				rest = -rest
				fmt.Println("требуется сдача :", rest)
			} else if payment-(coinOne+coinThree) < 0 {
				rest := payment - (coinOne + coinThree)
				rest = -rest
				fmt.Println("требуется сдача :", rest)
			} else if payment-(coinOne+coinTwo) > 0 {
				if payment-(coinOne+coinTwo+coinThree) == 0 {
					fmt.Println("Достаточно трёх монет для оплаты")
					fmt.Println("Сдача не требуется")
				} else if payment-(coinOne+coinTwo+coinThree) < 0 {
					fmt.Println("Достаточно трёх монет для оплаты")
					rest := payment - (coinOne + coinTwo + coinThree)
					rest = -rest
					fmt.Println("требуется сдача :", rest)
				} else {
					fmt.Println("Оплата не возможга, недостаточно средств")
					difference := payment - (coinOne + coinTwo + coinThree)
					fmt.Println("Вам не хватает:", difference)
				}
			}
		}
	} else if coinTwo <= coinThree {

		if payment-coinTwo == 0 {
			fmt.Println("Достаточно одной монеты для оплаты")
			fmt.Println("Сдача не требуется")
		} else if payment-(coinOne+coinThree) == 0 {
			fmt.Println("Достаточно двух монет для оплаты")
			fmt.Println("Сдача не требуется")
		} else if payment-coinTwo < 0 {
			fmt.Println("Достаточно одной монеты для оплаты")
			rest := payment - coinTwo
			rest = -rest
			fmt.Println("требуется сдача :", rest)
		} else if payment-(coinOne+coinThree) < 0 {
			fmt.Println("Достаточно двух монет для оплаты")
			rest := payment - (coinOne + coinThree)
			rest = -rest
			fmt.Println("требуется сдача :", rest)
		} else if payment-coinTwo > 0 && payment-(coinOne+coinThree) > 0 {
			if payment-(coinOne+coinTwo) == 0 || payment-(coinTwo+coinThree) == 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				fmt.Println("Сдача не требуется")
			} else if payment-(coinOne+coinTwo) < 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				rest := payment - (coinOne + coinTwo)
				rest = -rest
				fmt.Println("требуется сдача :", rest)
			} else if payment-(coinTwo+coinThree) < 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				rest := payment - (coinTwo + coinThree)
				rest = -rest
				fmt.Println("требуется сдача :", rest)
			} else if payment-(coinOne+coinTwo) > 0 {
				if payment-(coinOne+coinTwo+coinThree) == 0 {
					fmt.Println("Достаточно трёх монет для оплаты")
					fmt.Println("Сдача не требуется")
				} else if payment-(coinOne+coinTwo+coinThree) < 0 {
					fmt.Println("Достаточно трёх монет для оплаты")
					rest := payment - (coinOne + coinTwo + coinThree)
					rest = -rest
					fmt.Println("требуется сдача :", rest)
				} else {
					fmt.Println("Оплата не возможга, недостаточно средств")
					difference := payment - (coinOne + coinTwo + coinThree)
					fmt.Println("Вам не хватает:", difference)
				}
			}
		}

	} else {
		if payment-coinThree == 0 {
			fmt.Println("Достаточно одной монеты для оплаты")
			fmt.Println("Сдача не требуется")
		} else if payment-(coinTwo+coinOne) == 0 {
			fmt.Println("Достаточно двух монет для оплаты")
			fmt.Println("Сдача не требуется")
		} else if payment-coinThree < 0 {
			fmt.Println("Достаточно одной монеты для оплаты")
			rest := payment - coinThree
			rest = -rest
			fmt.Println("требуется сдача :", rest)
		} else if payment-(coinTwo+coinOne) < 0 {
			fmt.Println("Достаточно двух монет для оплаты")
			rest := payment - (coinTwo + coinOne)
			rest = -rest
			fmt.Println("требуется сдача :", rest)
		} else if payment-coinOne > 0 && payment-(coinTwo+coinThree) > 0 {
			if payment-(coinThree+coinTwo) == 0 || payment-(coinOne+coinThree) == 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				fmt.Println("Сдача не требуется")
			} else if payment-(coinThree+coinTwo) < 0 {
				fmt.Println("Достаточно двух монет для оплаты")
				rest := payment - (coinThree + coinTwo)
				rest = -rest
				fmt.Println("требуется сдача :", rest)
			} else if payment-(coinOne+coinThree) < 0 {
				rest := payment - (coinOne + coinThree)
				rest = -rest
				fmt.Println("требуется сдача :", rest)
			} else if payment-(coinThree+coinTwo) > 0 {
				if payment-(coinOne+coinTwo+coinThree) == 0 {
					fmt.Println("Достаточно трёх монет для оплаты")
					fmt.Println("Сдача не требуется")
				} else if payment-(coinOne+coinTwo+coinThree) < 0 {
					fmt.Println("Достаточно трёх монет для оплаты")
					rest := payment - (coinOne + coinTwo + coinThree)
					rest = -rest
					fmt.Println("требуется сдача :", rest)
				} else {
					fmt.Println("Оплата не возможга, недостаточно средств")
					difference := payment - (coinOne + coinTwo + coinThree)
					fmt.Println("Вам не хватает:", difference)
				}
			}
		}
	}
}

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа высчитывает если хватает барбберов")
	fmt.Println("для обслуживания всех мужчин города")

	fmt.Println("Сколько мужчин проживает в городе?")
	var mens int
	fmt.Scan(&mens)

	fmt.Println("Сколько барберов работают в барбершопах?")
	var barber int
	fmt.Scan(&barber)

	days := 30
	shift := 8

	notservesed := mens - (barber*shift)*days
	barbersNeeded := notservesed / (shift * days)

	if (barber*shift)*days >= mens || barbersNeeded == 0 {
		fmt.Println("Барберов достатачно для того чтоб обслужить всех мужчин за один месяц")
	} else if (barber*shift)*days < mens {
		fmt.Println("Вы не успеите обслужить всех мужчин в данном городе.")
		fmt.Println("Вам понадобится еще", barbersNeeded, ",барб.")
	}

}

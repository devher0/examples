package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет! Это плотформа текстовых квестов.")
	fmt.Println("Для начала введите имя вашего Рейнджера.")
	var Ranger string
	fmt.Scan(&Ranger)

	fmt.Println("Отлично! Теперь укажите с какой планеты ваши союзники.")
	var Planet string
	fmt.Scan(&Planet)

	fmt.Println("Хороший выбор! Укажите в какой звёздной системе находится данная планета.")
	var Star string
	fmt.Scan(&Star)

	fmt.Println("Замечательно! Теперь укажите за какую сумму кредивтов вы готовы помочь.")
	var Money int
	fmt.Scan(&Money)

	fmt.Println("Начинаем!")
	fmt.Println("=======================")
	fmt.Print("Как вам, ", Ranger, ", известно, мы раса мирная, поэтому на наших военных кораблях используются наёмники с других\n")
	fmt.Print("планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты\n")
	fmt.Print(Planet, " системы ", Star, ".\n")
	fmt.Print("", "\n")
	fmt.Print("Но случилась неприятность: в связи с большими потерями в последнее время престиж\n")
	fmt.Print("профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар\n")
	fmt.Print("планеты ", Planet, " впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с\n")
	fmt.Print("закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!\n")
	fmt.Print("", "\n")
	fmt.Print(Ranger, ", вы должны прибыть на планету ", Planet, " системы ", Star, " и помочь выполнить план\n")
	fmt.Print("призыва. Мы готовы выплатить вам премию в ", Money, " кредитов за эту маленькую услугу.")

}

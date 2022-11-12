package main

import (
	"fmt"
)

func main() {
	fmt.Println("Баллы ЕГЭ")
	fmt.Println("Введите результаты первого экзамена")
	var firstExam int
	fmt.Scan(&firstExam)
	fmt.Println("Введите результаты второго экзамена")
	var secondExam int
	fmt.Scan(&secondExam)
	fmt.Println("Введите результаты третьего экзамена")
	var thridExam int
	fmt.Scan(&thridExam)

	fmt.Println("Сумма проходных баллов: 275")

	result := firstExam + secondExam + thridExam
	fmt.Println("Количество набранных баллов:", result)

	if result >= 275 {
		fmt.Println("Вы зачислены!")
	} else if result < 275 && result > 0 {
		fmt.Println("Вы не поступили")
	} else {
		fmt.Println("Ошибка!")
		fmt.Println("Вы ввели не корректные значения!")
		fmt.Println("Перезапустите программу")
	}
}

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа высчитывает среднюю зарплату отдела")
	fmt.Println("А также высчитывает разницу между самой высокой и низкой зарпатой")

	fmt.Println("Введите зарплату первого сотрудника")
	var firstEmployer int
	fmt.Scan(&firstEmployer)

	fmt.Println("Введите зарплату второго сотрудника")
	var secondEmployer int
	fmt.Scan(&secondEmployer)

	fmt.Println("Введите зарплату третьего сотрудника")
	var thirdEployer int
	fmt.Scan(&thirdEployer)

	midSalary := (firstEmployer + secondEmployer + thirdEployer) / 3
	fmt.Println("Средняя зарплата отдела :", midSalary)

	if firstEmployer < secondEmployer && secondEmployer < thirdEployer {
		difference := thirdEployer - firstEmployer
		fmt.Println("разница между самой высокой и низкой зарплатой :", difference)
	} else if firstEmployer < thirdEployer && thirdEployer < secondEmployer {
		difference := secondEmployer - firstEmployer
		fmt.Println("разница между самой высокой и низкой зарплатой :", difference)
	} else if thirdEployer < secondEmployer && secondEmployer < firstEmployer {
		difference := firstEmployer - thirdEployer
		fmt.Println("разница между самой высокой и низкой зарплатой :", difference)
	} else if secondEmployer < thirdEployer && thirdEployer < firstEmployer {
		difference := firstEmployer - secondEmployer
		fmt.Println("разница между самой высокой и низкой зарплатой :", difference)
	} else if secondEmployer < firstEmployer && firstEmployer < thirdEployer {
		difference := thirdEployer - secondEmployer
		fmt.Println("разница между самой высокой и низкой зарплатой :", difference)
	} else if thirdEployer < firstEmployer && firstEmployer < secondEmployer {
		difference := secondEmployer - thirdEployer
		fmt.Println("разница между самой высокой и низкой зарплатой :", difference)
	} else {
		fmt.Println("Ошибка !")
		fmt.Println("Вы ввели неверные значения, перезапустите программу.")
	}

}

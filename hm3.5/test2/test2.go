package main

import (
	"fmt"
)

func main() {

	totalPassengers := 0
	totalMoney := 0
	var passengerGetOff int
	var passengerEntered int

	fmt.Println("Программа симуляции работы маршрутного такси.")
	fmt.Println("Прибываем на остановку «Улица Программистов».")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("Сколько пассажиров вышло на остановке ?")
	fmt.Scan(&passengerGetOff)
	totalPassengers -= passengerGetOff

	fmt.Println("Сколько пассажиров вошло на остановке ?")
	fmt.Scan(&passengerEntered)
	totalPassengers += passengerEntered
	totalMoney += passengerEntered * 20

	fmt.Println("Пассажиров вышло на остановке :", passengerGetOff)
	fmt.Println("Пассажиров вошло на остановке :", passengerEntered)
	fmt.Println("Отправляемся с остановки «Улица Программистов».")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку «Проспект Алгоритмов»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("Сколько пассажиров вышло на остановке ?")
	fmt.Scan(&passengerGetOff)
	totalPassengers -= passengerGetOff

	fmt.Println("Сколько пассажиров вошло на остановке ?")
	fmt.Scan(&passengerEntered)
	totalPassengers += passengerEntered
	totalMoney += passengerEntered * 20

	fmt.Println("Отправляемся с остановки «Проспект Алгоритмов»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку «Улица Зелинского»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("Сколько пассажиров вышло на остановке ?")
	fmt.Scan(&passengerGetOff)
	totalPassengers -= passengerGetOff

	fmt.Println("Сколько пассажиров вошло на остановке ?")
	fmt.Scan(&passengerEntered)
	totalPassengers += passengerEntered
	totalMoney += passengerEntered * 20

	fmt.Println("Отправляемся с остановки «Улица Зелинского»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку «Улица Мира»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("Сколько пассажиров вышло на остановке ?")
	fmt.Scan(&passengerGetOff)
	totalPassengers -= passengerGetOff

	fmt.Println("Сколько пассажиров вошло на остановке ?")
	fmt.Scan(&passengerEntered)
	totalPassengers += passengerEntered
	totalMoney += passengerEntered * 20

	fmt.Println("Отправляемся с остановки «Улица Мира»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку «Улица Менделеева»")
	fmt.Println("В салоне пассажиров:", totalPassengers)

	fmt.Println("Сколько пассажиров вышло на остановке ?")
	fmt.Scan(&passengerGetOff)
	totalPassengers -= passengerGetOff

	fmt.Println("Сколько пассажиров вошло на остановке ?")
	fmt.Scan(&passengerEntered)
	totalPassengers += passengerEntered
	totalMoney += passengerEntered * 20

	salary := totalMoney / 4
	fuelCost := totalMoney / 5
	taxes := totalMoney / 5
	carReparation := totalMoney / 5
	income := totalMoney - (salary + fuelCost + taxes + carReparation)

	fmt.Println("В салоне пассажиров:", totalPassengers)
	fmt.Println("Всего заработано:", totalMoney, "руб")
	fmt.Println("Зарплата водитлея:", salary, "руб")
	fmt.Println("Расходы на топливо:", fuelCost, "руб")
	fmt.Println("Налоги:", taxes, "руб")
	fmt.Println("Расходы на ремонт машины:", carReparation, "руб")
	fmt.Println("Итого доход:", income, "руб")

}

/*
Зарплата водителя: 25 рублей.
Расходы на топливо: 20 рублей.
Налоги: 20 рублей.
Расходы на ремонт машины: 20 рублей.
Итого доход: 15 рублей.
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет , для входа вам нужно зарегистрирыватся")
	fmt.Println("Введите имя")
	var name string
	fmt.Scan(&name)

	fmt.Println("Введите пароль")
	var password string
	fmt.Scan(&password)

	fmt.Println("Введите ваш возраст")
	var age int
	fmt.Scan(&age)

	fmt.Println("Поздравляю", name, "теперь вы зарегистрированы!")
	fmt.Println("Ваш пароль:", password)
	fmt.Println("Ваш возраст:", age)

}

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Здравствуй, войдите в ваш аккаунт")
	fmt.Println("Введите имя")
	var name string
	fmt.Scan(&name)

	fmt.Println("Введите пароль")
	var password string
	fmt.Scan(&password)

	fmt.Println("=============")
	fmt.Println(name, "вы успешно зашли!")
}

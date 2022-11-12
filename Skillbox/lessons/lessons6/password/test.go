package main

import (
	"fmt"
)

func main() {

	for {

		fmt.Println("Введите пароль")

		password := ""
		fmt.Scan(&password)

		if password == "sikkim222555" {
			fmt.Println("Вы успешно вошли!")
			break
		} else {
			fmt.Println("Пароль не верный, повторите опытку")
		}
	}

}

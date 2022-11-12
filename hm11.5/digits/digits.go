package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var stringWithDigits string
	var onlyDigits []int

	fmt.Print("Пожалуйста, введите строку для проверки: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringWithDigits = scanner.Text()

	s := strings.Split(stringWithDigits, " ")

	for i := 0; i < len(s); i++ {
		if digit, err := strconv.Atoi(s[i]); err == nil {
			onlyDigits = append(onlyDigits, digit)
		}
	}
	fmt.Println(onlyDigits)
}

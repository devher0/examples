package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var capitalLettersCount = 0
	var sentence string

	fmt.Print("Пожалуйста, введите исходную строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence = scanner.Text()

	f := func(c rune) bool {
		return c == ',' || c == '.' || c == ' '
	}
	s := strings.FieldsFunc(sentence, f)

	for i := 0; i < len(s); i++ {
		var word = s[i]
		if unicode.IsUpper(rune(word[0])) {
			capitalLettersCount++
		}
	}

	fmt.Printf("В тексте %v слов с заглавной буквы", capitalLettersCount)
}

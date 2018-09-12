package main

import "fmt"

// Разворачивает строку
func ReverseString(inputString string) string {
	runes := []rune(inputString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Меняем местами символы
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	inputStr := "абв пg"
	outputStr := ReverseString(inputStr)
	fmt.Println("Введеная строка: ", inputStr)
	fmt.Println("Перевернутая строка: ", outputStr)
}

package main

import (
	"fmt"
)

// ReverseString переворачивает строку, поддерживая символы Unicode.
func ReverseString(s string) string {
	// Преобразуем строку в слайс rune для корректной работы с Unicode символами.
	runes := []rune(s)

	// Переворачиваем слайс rune.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем обратно в строку и возвращаем результат.
	return string(runes)
}

func main() {
	originalString := "главрыба — абырвалг"
	reversedString := ReverseString(originalString)

	fmt.Printf("Оригинальная строка: %s\nПеревернутая строка: %s\n", originalString, reversedString)
}

package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке.
func ReverseWords(s string) string {
	// Разбиение строки на слова.
	words := strings.Fields(s)

	// Переворачивание слайса слов.
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Сборка слов обратно в строку.
	return strings.Join(words, " ")
}

func main() {
	originalString := "snow dog sun"
	reversedWordsString := ReverseWords(originalString)

	fmt.Printf("Оригинальная строка: %s\nСлова в обратном порядке: %s\n", originalString, reversedWordsString)
}

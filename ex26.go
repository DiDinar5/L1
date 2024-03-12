package main

import (
	"fmt"
	"strings"
)

// IsUnique проверяет, состоит ли строка из уникальных символов, не учитывая регистр.
func IsUnique(str string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	str = strings.ToLower(str)
	// Создаем карту для отслеживания встреченных символов
	seen := make(map[rune]bool)

	// Проверяем каждый символ в строке
	for _, char := range str {
		if _, exists := seen[char]; exists {
			// Если символ уже встречался, строка не уникальна
			return false
		}
		// Отмечаем символ как встреченный
		seen[char] = true
	}

	// Все символы уникальны
	return true
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", " aabcd"}

	for _, str := range testStrings {
		fmt.Printf("Строка \"%s\" состоит из уникальных символов: %t\n", str, IsUnique(str))
	}
}

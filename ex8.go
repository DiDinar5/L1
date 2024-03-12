package main

import (
	"fmt"
)

// SetBit устанавливает i-й бит числа num в значение bit (1 или 0)
// и возвращает результат. i считается с 0 и начинается с младшего бита.
func SetBit(num int64, i int, bit int) int64 {
	if bit == 1 {
		// Установка i-го бита в 1
		num |= (1 << i)
	} else {
		// Установка i-го бита в 0
		num &= ^(1 << i)
	}
	return num
}

func main() {
	var num int64 = 0 // Начальное число
	var i int = 5     // Позиция бита, который хотим изменить
	var bit int = 1   // Значение бита, которое хотим установить

	// Установка i-го бита в bit
	num = SetBit(num, i, bit)
	fmt.Printf("Результат после установки %d-го бита в %d: %d\n", i, bit, num)

	// Пример установки того же бита обратно в 0
	bit = 0
	num = SetBit(num, i, bit)
	fmt.Printf("Результат после установки %d-го бита в %d: %d\n", i, bit, num)
}

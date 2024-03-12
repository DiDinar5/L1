package main

import (
	"fmt"
)

func main() {
	// Инициализация каналов
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Массив чисел для обработки
	numbers := []int{1, 2, 3, 4, 5}

	// Горутина для чтения из первого канала, обработки чисел и записи во второй канал
	go func() {
		for x := range inputChan {
			outputChan <- x * 2 // Умножаем x на 2 и отправляем в outputChan
		}
		close(outputChan) // Закрываем outputChan после обработки всех чисел
	}()

	// Горутина для вывода результатов из второго канала в stdout
	go func() {
		for result := range outputChan {
			fmt.Println(result)
		}
	}()

	// Запись чисел из массива в первый канал
	for _, num := range numbers {
		inputChan <- num
	}
	close(inputChan) // Важно закрыть inputChan после записи всех чисел для завершения горутины
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	// Количество чисел для обработки
	n := len(numbers)
	// Создание канала для передачи результатов вычислений
	results := make(chan string, n)

	var wg sync.WaitGroup // Используем WaitGroup для ожидания завершения всех горутин

	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func(number int) {
			defer wg.Done() // По завершению горутины уменьшаем счетчик WaitGroup
			square := number * number
			// Отправляем результат в канал
			results <- fmt.Sprintf("Квадрат числа %d равен %d", number, square)
		}(number) // Передаем number как аргумент функции, чтобы избежать захвата переменной из цикла
	}

	wg.Wait()      // Ожидаем завершения всех горутин
	close(results) // Закрываем канал после завершения всех горутин

	// Выводим результаты из канала
	for result := range results {
		fmt.Println(result)
	}
}

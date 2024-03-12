package main

import (
	"fmt"
	"sync"
)

func square(number int, ch chan int) {
	ch <- number * number
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	sum := 0

	for _, number := range numbers {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			square(number, ch)
		}(number)
	}

	// Закрытие канала в отдельной горутине после завершения всех вычислений
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Сбор результатов из канала и вычисление суммы
	for result := range ch {
		sum += result
	}

	fmt.Printf("Сумма квадратов чисел равна: %d\n", sum)
}

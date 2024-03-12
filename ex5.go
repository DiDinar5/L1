package main

import (
	"fmt"
	"time"
)

func main() {
	var N int
	fmt.Print("Введите длительность работы программы в секундах: ")
	fmt.Scanf("%d", &N)

	dataChan := make(chan int) // Создаем канал для передачи целых чисел
	// Таймер для завершения программы по истечении N секунд
	timer := time.NewTimer(time.Duration(N) * time.Second)

	go func() {
		for i := 0; ; i++ {
			dataChan <- i
			time.Sleep(1 * time.Second) // Имитация работы
		}
	}()

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Время вышло")
				close(dataChan)
				return
			case value := <-dataChan:
				fmt.Println("Получено значение:", value)
			}
		}
	}()

	// Даем время для завершения горутин
	<-timer.C
}

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // Создание канала для передачи сигнала остановки

	go func() {
		for {
			select {
			case <-done: // Ожидание сигнала остановки
				fmt.Println("Горутина остановлена через done")
				return
			default:
				// Имитация некоторой работы
				fmt.Println("Работа в горутине (done)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second) // Имитация работы
	done <- true                // Отправка сигнала остановки
	time.Sleep(1 * time.Second) // Даем время для корректного завершения горутины
}

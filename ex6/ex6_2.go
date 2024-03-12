package main

import (
	"fmt"
	"time"
)

func main() {
	stopChan := make(chan struct{}) // Создание канала без передачи данных

	go func() {
		for {
			select {
			case <-stopChan: // Ожидание сигнала остановки
				fmt.Println("Горутина остановлена через stopChan")
				return
			default:
				// Имитация некоторой работы
				fmt.Println("Работа в горутине (stopChan)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second) // Имитация работы
	close(stopChan)             // Отправка сигнала остановки путем закрытия канала
	time.Sleep(1 * time.Second) // Даем время для корректного завершения горутины
}

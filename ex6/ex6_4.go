package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second) // Создание таймера на 3 секунды

	go func() {
		for {
			select {
			case <-timer.C: // Ожидание истечения таймера
				fmt.Println("Горутина остановлена по таймеру")
				return
			default:
				// Имитация некоторой работы
				fmt.Println("Работа в горутине (таймер)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	<-timer.C // Ожидаем истечение таймера в главной горутине
	// Таймер истекает автоматически, дополнительные действия для остановки горутины не требуются.
	time.Sleep(1 * time.Second) // Даем время для корректного завершения горутины
}

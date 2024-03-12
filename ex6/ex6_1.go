package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создание контекста с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())

	// Запуск горутины с передачей ей контекста
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Проверка на сигнал отмены из контекста
				fmt.Println("Горутина остановлена")
				return // Выход из горутины
			default:
				// Регулярное выполнение работы, если сигнала отмены не поступало
				fmt.Println("Горутина выполняет работу")
				time.Sleep(500 * time.Millisecond) // Имитация загруженной работы
			}
		}
	}(ctx)

	// Имитация работы в главной горутине (например, выполнение других задач)
	time.Sleep(2 * time.Second)

	// После некоторого времени отправляем сигнал остановки горутине
	cancel()

	// Даем время для корректного завершения работы горутины
	time.Sleep(1 * time.Second)

	fmt.Println("Главная горутина завершена")
}

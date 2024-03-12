package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// worker представляет собой воркера, который читает данные из канала
func worker(id int, dataChan <-chan string, doneChan <-chan bool) {
	for {
		select {
		case data, ok := <-dataChan: // Чтение данных из канала
			if !ok { // Если канал закрыт, завершаем работу воркера
				fmt.Printf("Воркер %d завершил работу\n", id)
				return
			}
			fmt.Printf("Воркер %d: %s\n", id, data)
		case <-doneChan: // Если пришёл сигнал на завершение, выходим из цикла
			fmt.Printf("Воркер %d получил сигнал о завершении\n", id)
			return
		}
	}
}

func main() {
	var N int
	fmt.Println("Введите количество воркеров:")
	fmt.Scan(&N) // Считываем количество воркеров

	dataChan := make(chan string) // Создаём канал для передачи данных воркерам
	doneChan := make(chan bool)   // Канал для сигнала о завершении работы воркеров

	// Запускаем воркеров
	for i := 0; i < N; i++ {
		go worker(i, dataChan, doneChan)
	}

	go func() {
		// Бесконечный цикл записи данных в канал
		for {
			select {
			case <-doneChan: // Если пришёл сигнал на завершение, выходим из цикла
				return
			default:
				dataChan <- fmt.Sprintf("Данные %v", time.Now())
				time.Sleep(time.Second) // Имитируем задержку
			}
		}
	}()

	// Обработка сигнала Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan // Ожидаем сигнал

	close(doneChan) // Посылаем сигнал воркерам о завершении работы
	close(dataChan) // Закрываем канал с данными
	fmt.Println("Программа завершена")
}

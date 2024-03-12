package main

import (
	"fmt"
	"time"
)

// MySleep приостанавливает выполнение программы на заданное количество секунд.
func MySleep(seconds int) {
	// Создаем канал, который получит сигнал после указанного времени.
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начинаем спать")
	MySleep(2) // "Спим" 2 секунды
	fmt.Println("Проснулись")
}

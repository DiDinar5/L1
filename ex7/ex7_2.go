package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	// Пример конкурентной записи в sync.Map
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(i, i*i) // Конкурентная запись
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Итерация по значениям sync.Map
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // Продолжить итерацию
	})
}

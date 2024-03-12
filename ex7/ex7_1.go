package main

import (
	"fmt"
	"sync"
)

// SafeMap структура, обеспечивающая потокобезопасный доступ к карте с помощью мьютекса
type SafeMap struct {
	m   map[int]int // Карта для хранения данных
	mux sync.Mutex  // Мьютекс для синхронизации доступа
}

// Set безопасно устанавливает значение в карте
func (sm *SafeMap) Set(key int, value int) {
	sm.mux.Lock()     // Блокировка мьютекса перед записью в карту
	sm.m[key] = value // Установка значения
	sm.mux.Unlock()   // Разблокировка мьютекса
}

// Get безопасно получает значение из карты
func (sm *SafeMap) Get(key int) (int, bool) {
	sm.mux.Lock()              // Блокировка мьютекса перед чтением из карты
	value, exists := sm.m[key] // Получение значения
	sm.mux.Unlock()            // Разблокировка мьютекса
	return value, exists
}

func main() {
	sm := SafeMap{m: make(map[int]int)}
	var wg sync.WaitGroup

	// Пример конкурентной записи в карту
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(i, i*i)
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Проверка значений в карте
	for i := 0; i < 100; i++ {
		if value, exists := sm.Get(i); exists {
			fmt.Printf("Key: %d, Value: %d\n", i, value)
		}
	}
}

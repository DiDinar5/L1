package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика с поддержкой конкурентного доступа.
type Counter struct {
	val int64
	mux sync.Mutex // Мьютекс для синхронизации доступа к val
}

// Increment увеличивает значение счетчика на 1.
// Метод защищен мьютексом, что делает его безопасным для использования в конкурентной среде.
func (c *Counter) Increment() {
	c.mux.Lock()   // Блокировка доступа к val
	c.val++        // Инкрементирование val
	c.mux.Unlock() // Разблокировка доступа
}

// Value возвращает текущее значение счетчика.
// Получение значения также защищено мьютексом для обеспечения консистентности.
func (c *Counter) Value() int64 {
	c.mux.Lock()         // Блокировка доступа к val
	defer c.mux.Unlock() // Разблокировка доступа будет выполнена автоматически после возврата значения
	return c.val
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{} // Создание счетчика

	// Запуск 1000 горутин, каждая из которых увеличивает счетчик на 1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}

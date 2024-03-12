package main

import "fmt"

// RemoveAtIndex удаляет элемент из слайса по указанному индексу,
// сохраняя порядок оставшихся элементов.
func RemoveAtIndex(a []int, index int) []int {
	if index < 0 || index >= len(a) {
		fmt.Println("Индекс вне диапазона слайса")
		return a
	}
	// Склеиваем часть слайса до индекса и после, исключая элемент под индексом.
	return append(a[:index], a[index+1:]...)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", a)

	index := 2 // Индекс элемента, который хотим удалить
	a = RemoveAtIndex(a, index)
	fmt.Println("Слайс после удаления элемента:", a)
}

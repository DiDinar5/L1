package main

import (
	"fmt"
)

// BinarySearch ищет элемент target в отсортированном срезе nums и возвращает его индекс.
// Если элемент не найден, возвращает -1.
func BinarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2 // Вычисляем средний индекс, избегая переполнения

		if nums[mid] == target {
			return mid // Элемент найден
		} else if nums[mid] < target {
			lo = mid + 1 // Искомый элемент находится в правой половине
		} else {
			hi = mid - 1 // Искомый элемент находится в левой половине
		}
	}

	return -1 // Элемент не найден
}

func main() {
	nums := []int{1, 2, 4, 5, 6, 8, 9} // Отсортированный срез
	target := 5                        // Элемент, который мы ищем

	index := BinarySearch(nums, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции: %d\n", target, index)
	} else {
		fmt.Println("Элемент не найден.")
	}
}

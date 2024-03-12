package main

import (
	"fmt"
)

// Intersect возвращает пересечение двух множеств
func Intersect(set1, set2 []int) []int {
	// Создание карты для отслеживания элементов первого множества
	set1Map := make(map[int]bool)
	for _, item := range set1 {
		set1Map[item] = true
	}

	// Поиск пересечений, проверяя наличие элементов второго множества в карте первого
	intersection := []int{}
	for _, item := range set2 {
		if _, exists := set1Map[item]; exists {
			intersection = append(intersection, item)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	intersection := Intersect(set1, set2)
	fmt.Println("Пересечение множеств:", intersection)
}

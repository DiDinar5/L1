package main

import (
	"fmt"
	"math"
)

// Point структура, представляющая точку в двумерном пространстве
type Point struct {
	x, y float64
}

// NewPoint конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo метод для вычисления расстояния между текущей точкой и другой точкой
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем и выводим расстояние между ними
	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

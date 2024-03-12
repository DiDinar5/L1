package main

import "fmt"

// Определение структуры Human
type Human struct {
	Name string
	Age  int
}

// метод для структуры Human
func (h Human) Greet() string {
	return fmt.Sprintf("Привет, меня зовут %s", h.Name)
}

// стуктура, которая "наследует" HUman
type Action struct {
	Human  //встраивание структуры
	Action string
}

// метод для структуры Action
func (a Action) Perform() string {
	return fmt.Sprintf("%s совершает действие: %s", a.Name, a.Action)
}

func main() {
	//определение полей Action
	action := Action{
		Human:  Human{Name: "Алекс", Age: 30},
		Action: "бегает",
	}
	//вызов функций
	fmt.Println(action.Greet())
	fmt.Println(action.Perform())
}

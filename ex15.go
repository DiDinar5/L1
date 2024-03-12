package main

func createHugeString(size int) string {
	// Допустим, что здесь какая-то логика создания большой строки
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // Создаем большую строку размером 1024 символа
	// Создаем новую строку на основе подстроки, чтобы избежать удержания всей большой строки в памяти
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}

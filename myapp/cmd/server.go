package main // Имя текущего пакета

// Импорты других пакетов
import "fmt"

// Неявная инициализация пакета
func init() {
	fmt.Println("Hello from init!")
}

// Функция main как точка входа
func main() {
	foo()
}

func foo() {
	fmt.Println("Foo!")
}

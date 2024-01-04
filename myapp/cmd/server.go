package main // Имя текущего пакета

// Импорты других пакетов
import (
	"fmt"

	"github.com/devomot/jenitter/most"
	_ "github.com/devomot/jenitter/most"
)

// Неявная инициализация пакета
func init() {
	fmt.Println("Hello from init!")
}

// Функция main как точка входа
func main() {
	foo()
}

func foo() {
	fmt.Println("Zigel")
	most.Mcad()
}

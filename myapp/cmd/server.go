package test // Имя текущего пакета

// Импорты других пакетов
import (
	"fmt"
	"time"
)

func findUserBySNILS() (string, bool) {
	return "Moscow", false
}

func dsc() int {
	return 13 + 20
}

func Example2(age int, myBool bool, str string) {
	var myVar = "123"

	switch age {
	case 25, 45:
		fmt.Println(30)
		fmt.Println(25)
	case dsc():
		fmt.Println(100)
		fallthrough
	case 13 + 20:
		fmt.Println(200)
		fallthrough
	default:
		fmt.Printf("Age = %d\n", age)
	}

	t := time.Now()
	switch {
	case t.Second() > 15:
		fmt.Printf("> 15 %v\n", t)
		fallthrough
	case t.UnixMicro() > 10000:
		fmt.Printf("> 10000 %v\n", t)
		fallthrough
	case !myBool:
		fmt.Printf("= myBool %v\n", t)
	default:
		fmt.Printf(">45 %v\n", t)
	}

	switch str {
	case "aaa", "bbb":
		fmt.Println("Not")
	case "ddd":
		fmt.Println("Yes")
	default:
		fmt.Printf("Str = %s\n", str)
	}
	fmt.Println("myVar", myVar)
}

// Функция main как точка входа
func main() {
	Example2(33, true, "bbb")
	fmt.Println("Foo-1")
}

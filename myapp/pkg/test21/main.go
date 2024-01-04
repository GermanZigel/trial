package main

import "fmt"

func main() {
	s := "hello world"            // в двойных кавычках, на одной строке
	s2 := "hello \n world \u9333" // c непечатными символами
	// если нужно включить в строку кавычки или переносы строки
	// - используем обратные кавычки
	s3 := `<div>hello</div>
"cruel"
'world'
`
	d1 := fmt.Sprintf("%s", "It's me")
	d2 := fmt.Sprintf(`%s`, "It's me")
	fmt.Println(s, s2, s3)
	fmt.Println(d1, d2)
}

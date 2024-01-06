package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		} else if i > 3 {
			fmt.Printf("II = %d\n", i)
			break
		}
		fmt.Printf("i = %d\n", i)
	}
}

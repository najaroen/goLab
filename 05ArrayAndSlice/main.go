package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		fmt.Println("Hello world->", numbers[i])
	}
	for _, number := range numbers {
		fmt.Println("index->", number)

	}
}

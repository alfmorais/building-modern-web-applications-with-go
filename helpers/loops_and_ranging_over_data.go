package main

import (
	"log"
)

func main() {
	// for i := 0; i < 100000; i++ {
	// 	log.Println(i)
	// }

	// arr := [5]string{"a", "b", "c", "d", "e"}
	// for index, value := range arr {
	// 	log.Println(index, value)
	// }

	var firstLine = "Once upon a time"

	for index, char := range firstLine {
		log.Println(index, ":", char, ":", string(char))
	}
}

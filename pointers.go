package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to:", myString)
	changeUsinngPointer(&myString)
	log.Println("after func call myString is set to:", myString)
}

func changeUsinngPointer(s *string) {
	log.Println("`s` is set to:", s)
	newValue := "Blue"
	*s = newValue
}

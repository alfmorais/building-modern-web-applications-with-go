package main

import "fmt"

func main() {
	var whatToSay string = "Hello, world."
	var number int
	fmt.Println(whatToSay)
	fmt.Println("Number =", number)

	number = 42
	fmt.Println("Number =", number)

	whatWasSaid, whatWasSaidSeond := saySomething()
	fmt.Println(whatWasSaid, whatWasSaidSeond)
}

func saySomething() (string, string) {
	return "something", "else"
}

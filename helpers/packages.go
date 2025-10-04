package main

import (
	"log"

	"github.com/alfmorais/myprogram/helpers"
)

func main() {
	log.Println("Hello, World!")
	var myVar helpers.SomeType = helpers.SomeType{
		TypeName:   "Example",
		TypeNumber: 1,
	}
	log.Println("myVar is set to:", myVar)
}

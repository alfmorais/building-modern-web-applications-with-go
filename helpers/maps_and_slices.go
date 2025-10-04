package main

import (
	"log"
	"sort"
)

// type User struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

func main() {
	// What make functions does?
	// myMap := make(map[string]string)

	// myMap["dog"] = "John"
	// myMap["cat"] = "Alice"
	// myMap["dog"] = "Kyra"

	// log.Println((myMap["dog"]))
	// log.Println((myMap["cat"]))

	// myMap2 := make(map[string]int)
	// myMap2["one"] = 1
	// myMap2["two"] = 2
	// myMap2["three"] = 3

	// log.Println(myMap2)
	// log.Println(myMap2["two"])
	// myMap := make(map[string]User)

	// me := User{
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Age:       30,
	// }

	// myMap["me"] = me
	// log.Println((myMap["me"].FirstName))

	// var myNewMap float32
	// myNewMap = 3.14
	// log.Println(myNewMap)
	var mySlice []string
	mySlice = append(mySlice, "dog")
	mySlice = append(mySlice, "cat")
	mySlice = append(mySlice, "fish")
	mySlice = append(mySlice, "bird")
	log.Println(mySlice)
	log.Println(mySlice[2])
	log.Println(len(mySlice))
	log.Println(cap(mySlice))
	log.Println(mySlice[len(mySlice)-1])
	log.Println(mySlice[1:3])
	log.Println(mySlice[0:2])
	log.Println(mySlice[2:])
	log.Println(mySlice[:3])
	log.Println(mySlice[:])
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	})
}

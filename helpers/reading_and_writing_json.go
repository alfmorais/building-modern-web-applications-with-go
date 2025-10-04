package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	superHeroesJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "Black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "Dark Brown",
			"has_dog": false
		}
	]
	`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(superHeroesJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}
	log.Printf("Unmarshalled Data: %+v\n", unmarshalled)
	log.Printf("Unmarshalled Data: %v\n", unmarshalled)

	var mySlice []Person
	var m1 Person
	m1.FirstName = "Diana"
	m1.LastName = "Prince"
	m1.HairColor = "Black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	m2 := Person{
		FirstName: "Barry",
		LastName:  "Allen",
		HairColor: "Blond",
		HasDog:    true,
	}

	mySlice = append(mySlice, m2)

	marshalled, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}
	log.Println("Marshalled Data:", string(marshalled))
	fmt.Printf("Marshalled Data: %s\n", marshalled)
}

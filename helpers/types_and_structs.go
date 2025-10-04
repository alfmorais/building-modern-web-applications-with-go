package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		PhoneNumber: "123-456-7890",
		BirthDate:   time.Date(1993, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	log.Println(user)
	log.Println(user.FirstName)
	log.Println(user.LastName)
	log.Println(user.Age)
	log.Println(user.PhoneNumber)
	log.Println(user.BirthDate)
}

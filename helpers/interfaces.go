package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	var dog Dog = Dog{Name: "Buddy", Breed: "Golden Retriever"}
	var gorilla Gorilla = Gorilla{Name: "George", Color: "Black", NumberOfTeeth: 32}

	PrintInfo(&dog)
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says:", a.Says())
	log.Println("This animal has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ooh ooh aah aah!"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

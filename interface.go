package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	Bibi := Person{Name: "Bibi"}

	SayHello(Bibi)

	var bibi Person
	bibi.Name = "bibi"

	SayHello(bibi)

	cat := Animal{
		Name: "Pus",
	}

	SayHello(cat)
}

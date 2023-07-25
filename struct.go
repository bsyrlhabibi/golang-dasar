package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayGodBye() {
	fmt.Println("See you from", a.Name)
}

func main() {
	var bibi Customer
	bibi.Name = "Bibi"
	bibi.Address = "Semarang"
	bibi.Age = 22

	bibi.sayHi("Joko")
	bibi.sayGodBye()

	fmt.Println(bibi.Name)
	fmt.Println(bibi.Address)
	fmt.Println(bibi.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Tegal",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 20}
	fmt.Println(budi)
}

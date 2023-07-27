package main

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bibi := Man{"Bibi"}
	bibi.Married()
	println(bibi.Name)
}

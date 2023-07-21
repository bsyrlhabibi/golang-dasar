package main

import "fmt"

func main() {

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	//init statement & post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for slice
	slice := []string{"Basyarul", "Habibi", "Bibi"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range tanpa index
	for _, value := range slice {
		fmt.Println(value)
	}
	//for range dengan index
	// names := []string{"Basyarul", "Habibi", "Bibi"}
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	// for range map

	person := make(map[string]string)
	person["name"] = "Bibi"
	person["title"] = "Student"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

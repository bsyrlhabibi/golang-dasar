package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	// Your main code here
	total := sumAll(10, 90, 40, 50)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}

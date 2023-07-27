package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	result, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}

package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var NoKtpBibi NoKtp = "1234567890001"
	var marriedStatus Married = false

	fmt.Println(NoKtpBibi)
	fmt.Println(marriedStatus)
}

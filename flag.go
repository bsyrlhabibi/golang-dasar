package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "locahost", "Put your host")
	var user *string = flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)

}

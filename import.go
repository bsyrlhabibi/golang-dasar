package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.SayHello("Habibi")
	//helper.sayGoodBye("Habibi") //error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) //error
}

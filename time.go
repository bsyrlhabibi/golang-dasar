package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Local())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// utc := time.Date(2023, time.September, 10, 23, 0, 0, 0, time.UTC)
	// fmt.Println(utc)

	utc := time.Date(2023, 10, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	// time.RFC3339 //time format

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2023-10-01")
	fmt.Println(parse)
}

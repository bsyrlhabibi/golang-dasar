package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian bool = ujian >= 80
	var lulusAbsensi bool = absensi > 80
	fmt.Println(lulusUjian)   //cek nilai boolean
	fmt.Println(lulusAbsensi) //cek nilai boolean

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	//for expert
	fmt.Println(ujian >= 80 && absensi >= 80)
}

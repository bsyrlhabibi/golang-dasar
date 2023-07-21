package main

import (
	"errors"
	"fmt"
)

func main() {
	// Deklarasi variabel dataBatas dengan slice dan inisialisasi
	var dataBatas = make([]int, 0, 5)

	// Batas maksimum untuk loop
	var batas_maksimum = 10 // Anda bisa mengubahnya menjadi lebih dari 7

	// Loop untuk menambahkan data ke dalam slice hingga mencapai batas maksimum
	for i := 1; i <= 15; i++ { // Ubah i <= 11 untuk melampaui batas maksimum
		if i <= batas_maksimum {
			// Melakukan append data ke slice
			dataBatas = append(dataBatas, i)
			fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(dataBatas), cap(dataBatas))
		} else {
			err := errors.New("Error: i melebihi batas maksimum!")
			fmt.Println(err)
			break // Keluar dari loop jika i melebihi batas maksimum
		}
	}

	// Mencetak batas maksimum sebagai hasilnya
	fmt.Printf("Batas maksimum: %d\n", batas_maksimum)

	// Mengakses data di dalam slice
	fmt.Println(dataBatas)
}

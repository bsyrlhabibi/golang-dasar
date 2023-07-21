package main

import "fmt"

func main() {
	// Deklarasi variabel data dengan slice dan inisialisasi
	var data = make([]int, 2, 5)

	// Deklarasi variabel i dan inisialisasi dengan nilai 1
	var i = 1

	// Loop untuk menambahkan data ke dalam slice
	for i <= 10 {
		// Melakukan append data ke slice
		data = append(data, i)
		fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(data), cap(data))

		// Increment nilai i
		i++
	}

	// Mengakses data di dalam slice
	fmt.Println(data)

	//BATAS MAKSIMUM
	// Deklarasi dan inisialisasi slice dengan nilainya sebagai slice kosong
	var dataBatas = make([]int, 0, 5)

	// Batas maksimum untuk loop
	var batas_maksimum = 7

	// Loop untuk menambahkan data ke dalam slice hingga mencapai batas maksimum
	for i := 1; i >= batas_maksimum; i++ {
		// Melakukan append data ke slice
		data = append(dataBatas, i)
		fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(data), cap(dataBatas))
	}

	// Mengakses data di dalam slice
	fmt.Println(dataBatas)
}

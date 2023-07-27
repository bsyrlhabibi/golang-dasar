package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	address1 := Address{"Semarang", "Jawa Tengah", "Indonesia"}

	//address2 := Address1 // Address1 tidak berubah

	// Address2 pointer ke Address1 data berubah "&"
	address2 := &address1
	address2.City = "Bandung"

	//Address 1 tetap Address2 tetap tidak berubah "&"
	//address2 = &Address{"Tegal", "Jawa Tengah", "Indonesia"}

	//Semua yang mengacu ke Address akan berubah "*"
	*address2 = Address{"Tegal", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	//Membuat Address baru
	var address3 *Address = new(Address)
	address3.City = "Jakarta"
	fmt.Println(address3)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// defer pulihkan_saya()

	// membuat variabel nama untuk menampung input scanln
	var nama string
	fmt.Print("Masukan Nama :")
	fmt.Scanln(&nama)

	// membuat kondisi jika terjadi error pada input
	if valid, err := validasi(nama); valid {
		fmt.Println("halo", nama)
	} else {
		panic(err.Error())
	}

}

// membuat function validasi ketika input kosong
func validasi(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Tidak Bisa Kosong")
	}
	return true, nil
}

/* func pulihkan_saya() {
	if r := recover(); r != nil {
		fmt.Println("Akhirnya saya ditampilkan")
	} else {
		fmt.Println("Ini Lancar Saja")
	}
} */

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type callback func(hasil []string)

// func pencarianNama(nama string, jumlah int, callback callback) {
// 	var hasilPencarian []string
// 	//
// 	for _, n := range name { //Underscore (_) sebagai blank identifier, variabel digunakan untuk mengabaikan nilai yang dikembalikan oleh sebuah fungsi atau variabel yang tidak diperlukan dalam suatu konteks.
// 		if strings.Contains(n, nama) {
// 			hasilPencarian = append(hasilPencarian, n)
// 			if len(hasilPencarian) == jumlah {
// 				break
// 			}
// 		}
// 	}
// 	callback(hasilPencarian)
// }

// func main() {
// 	pencarianNama("An", 3, func(hasil []string) { //hasil pencarian (kondisi, jumlah array, hasil(array))
// 		fmt.Println("Output:", hasil)
// 	})
// }

// var name = []string{"Abigail", "Alexandra", "Alison", "Amanda", "Bella", "Carol", "Caroline", "Carolyn", "Deirde", "Diana", "Elizabeth", "Ella", "Faith", "Olivia", "Penelope"}

package main

import (
	"fmt"
	"strings"
)

type callback func(hasil []string)

func pencarianNama(nama string, jumlah int, callback callback) {
	var hasilPencarian []string
	//
	for _, n := range name {
		// Mengubah kedua string menjadi huruf kecil agar pencarian tidak case-sensitive
		if strings.Contains(strings.ToLower(n), strings.ToLower(nama)) {
			hasilPencarian = append(hasilPencarian, n)
			if len(hasilPencarian) == jumlah {
				break
			}
		}
	}
	callback(hasilPencarian)
}

func main() {
	pencarianNama("AN", 3, func(hasil []string) {
		fmt.Println("Output:", hasil)
	})
}

var name = []string{"Abigail", "Alexandra", "Alison", "Amanda", "Bella", "Carol", "Caroline", "Carolyn", "Deirde", "Diana", "Elizabeth", "Ella", "Faith", "Olivia", "Penelope"}

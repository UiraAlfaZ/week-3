package main

import (
	"fmt"
)

func SeleksiNilai(nilaiAwal, nilaiAkhir int, dataArray []int) interface{} { //interface digunakan karena data array pada beberapa seleksiNilai tidak sama
	//jumlah data dalam dataArray harus dari 5, kondisi no 3 didapat
	if len(dataArray) <= 5 {
		return "Jumlah angka dalam dataArray harus lebih dari 5"
	}

	//nilaiAwal < nilaiAkhir, kondisi no 2 didapat
	if nilaiAwal >= nilaiAkhir {
		return "Nilai akhir harus lebih besar dari nilai awal"
	}

	// menggunakan range untuk mencari jumlah data array, lalu nilai yang muncul harus lebih besar dari nilaiAwal dan lebih kecil dari nilaiAkhir
	var hasilPencarian []int
	for _, nilai := range dataArray {
		if nilai >= nilaiAwal && nilai <= nilaiAkhir {
			hasilPencarian = append(hasilPencarian, nilai)
		}
	}

	// kondisi no 4 didapat
	if len(hasilPencarian) > 0 {
		return hasilPencarian
	} else {
		return "Nilai tidak ditemukan"
	}
}

func main() {
	// fungsi SeleksiNilai (nilaiAwal, nilaiAkhir, dataArray[]int{})
	fmt.Println(SeleksiNilai(5, 20, []int{2, 25, 4, 14, 17, 30, 8})) //nilaiAwal harus lebih dari 5, nilaiAkhir lebih dari 20, data array yang ditemukan 14, 17, 8.
	fmt.Println(SeleksiNilai(15, 3, []int{2, 25, 4, 14, 17, 30, 8}))
	fmt.Println(SeleksiNilai(5, 17, []int{2, 25, 4}))
	fmt.Println(SeleksiNilai(5, 17, []int{2, 25, 4, 1, 30, 18})) //tidak ada nilai diantar 5 dan 17, sehingga outputnya "Nilai tidak ditemukan"
}

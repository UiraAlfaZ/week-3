package main

import (
	"fmt"
	"math"
)

func main() {
	//fungsi untuk mencetak ke layar
	fmt.Println("Hello, World!") //output "Hello, World!"

	//fungsi untuk menghitung panjang string
	str := "Hello, World!"
	length := len(str)
	fmt.Println("panjang string:", length) //output "panjang string: 13"

	//fungsi untuk menghitung akar kuadrat dari suatu angka
	squareRoot := math.Sqrt(25)
	fmt.Println("akar kuadrat dari 25:", squareRoot) //output "akar kuadrat dari 25: 5"

	//Fungsi untuk mengonversi nilai float64 menjadi int
	floatValue := 3.14
	intValue := int(floatValue)
	fmt.Println("konversi float ke int:", intValue) // output "konversi float ke int: 3"

	//fungsi untuk mengonversi nilai int menjadi float64
	intValue2 := 42
	floatValue2 := float64(intValue2)
	fmt.Println("konversi int ke float:", floatValue2) // output "konversi int ke float: 42"

	//fungsi untuk menghasilkan nilai maksimum dari dua angka
	maxValue := math.Max(10, 20)
	fmt.Println("nilai maksimum:", maxValue) //output "nilai maksimum: 20"

	//fungsi untuk menghasilkan nilai minimum dari dua angka
	minValue := math.Min(30, 40)
	fmt.Println("nilai minimum:", minValue) //output "nilai minimum: 30"

	//fungsi untuk membuat map baru
	myMap := make(map[string]int)
	myMap["a"] = 1
	myMap["b"] = 2
	fmt.Println("membuat map:", myMap) // output "membuat map: map[a:1 b:2]"

	//fungsi untuk menghasilkan nilai absolut dari suatu angka
	absValue := math.Abs(-50)
	fmt.Println("nilai absolut:", absValue) // output "nilai absolut: 50"

	//fungsi untuk menghasilkan nilai pangkat dari suatu angka
	powValue := math.Pow(2, 3)
	fmt.Println("2^3:", powValue) //output "2^3: 8"
}

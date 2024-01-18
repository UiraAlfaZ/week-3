package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func denganReturn(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

func tanpaReturn(rect Rectangle) {
	area := rect.Width * rect.Height
	fmt.Printf("Luas dari Rectangle dengan Width %.2f dan Height %.2f adalah %.2f\n", rect.Width, rect.Height, area)
}

func main() {
	tanpaPengembalian := Rectangle{Width: 4.2, Height: 7.2}

	Return := denganReturn(tanpaPengembalian)
	fmt.Printf("Luas (dengan pengembalian): %.2f\n", Return)

	tanpaReturn(tanpaPengembalian)
}

//jawaban sebelumnya
// type Rectangle struct {
//   Width  float64
//   Height float64
// }

// //tanpa return

// //perintah mengkalikan pada nilai variabel width * height
// func luas(rect Rectangle) float64 {
//   return rect.Width * rect.Height
// }

// func main() {
//   //mengisi nilai width dan height untuk dikalikan
//   persegi := Rectangle{Width: 4.2, Height: 7.2}
//   area := luas(persegi)

//   result := fmt.Sprint(persegi.Width) + "*" + fmt.Sprint(persegi.Height) + "=" + fmt.Sprint(area)

//   fmt.Println(result)
// }

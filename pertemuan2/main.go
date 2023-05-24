package main

import "fmt"

func main() {
	var alas int = 30 // normal decl
	tinggi := 60      //short decl
	var luas int

	luas = (alas * tinggi) / 2

	fmt.Println(luas)

	//kondisi
	var umur int = 17

	if umur == 17 {
		fmt.Println("Happy Sweet Seventeen")
	} else if umur < 17 && umur > 7 {
		fmt.Println("Anda belum berumur 17 tahun")
	} else {
		fmt.Println("ada bukan umur 17")
	}

	//perulangan
	for i := 1; i < 5; i++ {
		fmt.Println("Hello Prakerja")
	}
	for a := 1; a <= 5; a++ {
		for b := 1; b <= a; b++ {
			print("*")
		}
		fmt.Println("*")
	}
}

// NOTE
// go bersifat static type; js bersifat dinamic
/* var name string = "alterra"
var isGenap bool = true
var phi float32 = 3.14

oprator:
*) aritmatika : + - * / % ++ --
*) kondisi : == != < > <= >=
*) logika : && || !
*) penugasan : = += -= *= /= %=
*) Bitwise : & | ^ << >> &^
*/

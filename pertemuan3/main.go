package main

import "fmt"

func main() {
	// Array (statis)
	// var umurAnak [6]int = [6]int{1, 3, 2}
	umurAnak := [6]int{1, 3, 2}
	fmt.Println(umurAnak) //[1 3 2 0 0 0]

	// Slice (dinamis)
	// var dataSlice []int
	// fmt.Println(dataSlice) //[]
	// var dataSlice []int
	// dataSlice = append(dataSlice, 10)
	// dataSlice = append(dataSlice, 20, 30)
	// fmt.Println(dataSlice) //[10 20 30]
	var dataSlice []int
	dataSlice = append(dataSlice, 10)
	dataSlice = append(dataSlice, 20, 30)
	dataSlice = append(dataSlice[0:1], dataSlice[2:3]...) //titik 3: variadic
	// dataSlice = append(dataSlice[:1], dataSlice[2:]...)
	fmt.Println(dataSlice)

	// Map(key, vallue)
	dataBulan := map[string]int{"maret": 03, "agustus": 5}
	dataBulan["januari"] = 10
	omset := dataBulan["januari"]
	fmt.Println(dataBulan, omset)
}

//NOTE
/*
Struktur Data
1. Array : menyimpan banyak nilai, tipe data sama, index, ukuran tidak dapat diubah
2. Slice : Referensi ke Bagian dari Data, Dinamis, Terhubung dengan Array, Panjang dan Kapasitas, Pemotongan (Slicing), Mutable, Alokasi Memori dan Dapat Digunakan sebagai Parameter dan Nilai Kembali
	 Membuat slice kosong
	var angka []int
	fmt.Println(angka)      // Output: []

	 Membuat slice dengan menggunakan literal slice
	huruf := []string{"a", "b", "c", "d"}
	fmt.Println(huruf)      // Output: [a b c d]

	 Mengakses elemen slice menggunakan indeks
	fmt.Println(huruf[2])   // Output: c

	 Mengubah nilai elemen slice
	huruf[1] = "x"
	fmt.Println(huruf)      // Output: [a x c d]

	 Mengiris (slicing) slice
	iris := huruf[1:3]
	fmt.Println(iris)       // Output: [x c]

	 Menambahkan elemen baru ke slice
	huruf = append(huruf, "e")
	fmt.Println(huruf)      // Output: [a x c d e]

	 Menggabungkan dua slice
	huruf2 := []string{"f", "g"}
	huruf = append(huruf, huruf2...)
	fmt.Println(huruf)      // Output: [a x c d e f g]
3. Map
*/

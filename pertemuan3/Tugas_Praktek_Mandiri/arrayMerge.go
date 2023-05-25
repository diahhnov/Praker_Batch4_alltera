package main

import "fmt"

func arrayMerge(arrayA, arrayB []string) []string {
	// membuat map untuk menympan elemen yang sudah ada
	mergeMap := make(map[string]bool)

	// array hasil gabungan
	var mergeArray []string

	// menambahka elemen dari arrayA ke array gabungan
	for _, key := range arrayA {
		mergeMap[key] = true
		mergeArray = append(mergeArray, key)
	}

	// menambahkan elemen dari arrayB ke array gabungan jika belum ada di mergeMap
	for _, key := range arrayB {
		if !mergeMap[key] {
			mergeMap[key] = true
			mergeArray = append(mergeArray, key)
		}
	}
	return mergeArray
}

func main() {
	fmt.Println(arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))          // [king devil jin akuma eddie steve geese]
	fmt.Println(arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))                         //[sergei jin steve bryan]
	fmt.Println(arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"})) // [alisa yoshimitsu devil jin law]
	fmt.Println(arrayMerge([]string{}, []string{"devil jin", "sergei"}))                                          // [devil jin sergei]
	fmt.Println(arrayMerge([]string{"hwoarang"}, []string{}))                                                     // [hwoarang]
	fmt.Println(arrayMerge([]string{}, []string{}))                                                               // []
}

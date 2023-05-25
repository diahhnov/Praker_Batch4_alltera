package main

import "fmt"

func munculSekali(angka string) []int {
	count := make(map[int]int)

	//menghitung kemunculan setiap angka string
	for _, key := range angka {
		count[int(key-'0')]++
	}

	//mencari angka yang hanya muncul sekali
	var result []int
	for key, deret := range count {
		if deret == 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}

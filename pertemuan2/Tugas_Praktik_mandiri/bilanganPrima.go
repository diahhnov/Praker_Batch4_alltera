package main

import (
	"fmt"
	"math"
)

func numPrimer(num int) bool {
	if num <= 1 {
		return false
	}
	// Mengecek faktor bilangan hingga akar kuadrat dari num
	cekNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= cekNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	countPrimer := 0
	fmt.Println("Bilangan prima:")

	for i := 1; i <= 100; i++ {
		if numPrimer(i) {
			fmt.Println(i)
			countPrimer++
		}
	}
	fmt.Println("Jumlah bilangan prima:", countPrimer)
}

package main

import "fmt"

func Mapping(slice []string) map[string]int {
	count := make(map[string]int)

	for _, key := range slice {
		count[key]++
	}
	return count
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[asd:2, qwe:3, adi: 1]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      //map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         //map[]
}

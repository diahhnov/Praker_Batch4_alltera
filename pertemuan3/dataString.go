package main

import (
	"fmt"
)

func main() {
	// var name string = "Alterra Academy"
	// fmt.Println(strings.Contains(name, "terras"))
	// fmt.Println(strings.ReplaceAll(name, "a", "i"))
	// fmt.Println(name[0:8], "Prakerja Batch 4", name[9:])

	var nameArray = [3]string{"Alterra", "Prakerja", "Academy"}

	for i := 0; i < len(nameArray); i++ {
		fmt.Println(nameArray[i])
	}

	var nameMap = map[string]bool{"Alterra": true, "Prakerja": true, "Academy": true}
	for _, key := range nameMap {
		fmt.Println(key)
	}
}

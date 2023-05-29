// package main

// import "fmt"

// func main() {
// 	//pointer
// 	var umur int = 20
// 	var name string = "alterra"
// 	// var alamat string = "jalan bintaro"
// 	var alamatUmur *int = &umur
// 	fmt.Println(&umur, &name, &alamatUmur)// 0xc00001a098 0xc000050270 0xc00000a028
// 	fmt.Println(alamatUmur, &alamatUmur)// 0xc00001a098 0xc00000a028
// }

// package main

// import "fmt"

// type People struct {
// 	name    string
// 	age     int
// 	address string
// 	health  int
// }

//metode
// func (people People) Jalan() int {
// 	return people.age * 2
// }

// func main() {
// var people People // people := People{}
// people.name = "Diah"
// people.age = 10
// people.address = "jaktim"
// people.Jalan() // memanggil metode
// people.health = Jalan(human) // memangil func
// fmt.Println(people.health)

// ubahData(people)
// fmt.Println(people.name, people.age, people.address) // memanggil
// }

// func ubahData(people People) {
// 	people.name = ""
// 	people.age = 1
// 	people.address = ""
// }

/*
 */

// Interface : jembatan antara request dengan server
package main

import "fmt"

type People struct {
	name    string
	age     int
	address string
	health  int
}

type HumanInterface interface {
	Lari()
	Jalan()
}

func (people *People) Lari() {
	people.health = people.health * 2
}

//metode
func (people *People) Jalan() {
	people.health = people.health * 2
}

func main() {
	people := People{}
	people.name = "Diah"
	people.age = 10
	people.address = "jaktim"
	people.health = 17
	people.Jalan() // memanggil metode
	// people.health = Jalan(human) // memangil func
	fmt.Println(people.health)

	// ubahData(people)
	// fmt.Println(people.name, people.age, people.address) // memanggil

	var humanInterface HumanInterface
	humanInterface = &people
	humanInterface.Lari()
	humanInterface.Jalan()
	fmt.Println(humanInterface.Jalan)
	fmt.Println(humanInterface.Lari)
}

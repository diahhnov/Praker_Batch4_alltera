package main

import "fmt"

func main() {
	fmt.Println("ini adalah bilangan kelipatan 7:")
	for i := 1; i < 100; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}

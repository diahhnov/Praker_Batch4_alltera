package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) CalculateDistance() float64 {
	flueEfek := 1.5 // liter per mill

	//menghitung jarak perkiraan yang bisa ditempuh
	jarak := c.FuelIn * flueEfek

	return jarak
}

func main() {
	car := Car{
		Type:   "BMW",
		FuelIn: 20,
	}

	jarak := car.CalculateDistance()
	fmt.Printf("Perkiraan jarak yang bbisa ditempuh: %.2f km\n", jarak)
}

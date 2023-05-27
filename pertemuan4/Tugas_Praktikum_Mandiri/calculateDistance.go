package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) CalculateDistance() float64 {
	flueEfek := 1.5 // liter per mill

	//menghitung jarak perkiraan yang bisa ditempuh
	distance := c.FuelIn * flueEfek

	return distance
}

func main() {
	car := Car{
		Type:   "BMW",
		FuelIn: 20,
	}

	distance := car.CalculateDistance()
	fmt.Printf("Perkiraan jarak yang bbisa ditempuh: %.2f km\n", distance)
}

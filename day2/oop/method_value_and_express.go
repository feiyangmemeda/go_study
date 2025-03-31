package main

import "fmt"

func main() {
	car := Car{
		price: 100,
		brand: "BMW",
	}
	car.printInfo()
	// method value,no ()
	info := car.printInfo
	info()
	// method express
	info2 := (*Car).printInfo
	info2(&car)
}

type Car struct {
	price int
	brand string
}

func (c *Car) printInfo() {
	fmt.Println(*c)
}

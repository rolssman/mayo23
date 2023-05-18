package main

import "fmt"

// Es la forma de hacer clases
type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 1999
	fmt.Println(otherCar)
}

package main

import (
	"fmt"
)

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi1: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado

	fmt.Println("El area de un cuadrado se calcula con a*a")

	const basedecuadrado = 10
	areadecuadrado := basedecuadrado * basedecuadrado
	fmt.Println("el lado es: ", basedecuadrado)
	fmt.Println("el area es: ", areadecuadrado)
}

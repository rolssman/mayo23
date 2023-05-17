package main

import "fmt"

func main() {
	//Area base de un cuadrado
	const lado = 10
	areacuadrado := lado * lado
	fmt.Println("area de un cuadrado de 10 = ", areacuadrado)

	x := 10
	y := 50

	fmt.Println("Valores en las operaciones x = ", x)
	fmt.Println("Valores en las operaciones y = ", y)
	fmt.Println(" ")
	//suma
	result := x + y
	fmt.Println("Suma de ", x, "+", y, "=", result)

	//resta
	result = y - x
	fmt.Println("Resta de ", y, "-", x, "=", result)

	//multiplicación
	result = x * y
	fmt.Println("Multiplicación de ", x, "*", y, "=", result)

	//división
	result = x / y
	fmt.Println("División de ", x, "/", y, "=", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo de ", x, "%", y, "=", result)

	//incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	y--
	fmt.Println("Decremental:", y)
}

package main

import "fmt"

// Por estandar la segunda palabra de la función va con mayuscula
func normalFunction(message string) {
	fmt.Println("Hola mundo ", message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, c, b)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func areaCuadrado(a int) int {
	return a * a
}

func areaTriangulo(a, b int) float64 {
	valor := (b * a) / 2
	return float64(valor)
}

func main() {
	//
	normalFunction("1")
	tripleArgument(11, 12, " y ")
	value := returnValue(2)
	fmt.Println("valor: ", value)
	value1, value2 := dobleReturn(2)
	fmt.Println("Valor 1 y Valor 2:", value1, value2)
	area2 := areaCuadrado(10)
	fmt.Println("el area de un cuadrado de 10 es ", area2)
	area3 := areaTriangulo(3, 10)
	fmt.Println("El area de un triangulo de alto 3 y base 10 es: ", area3)
}

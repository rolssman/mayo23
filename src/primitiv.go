package main

import "fmt"

func main() {
	//Explicación de datos primitivos
	fmt.Println("int 8/16/32/64: indica el tipo de dato int con signo, y además el tamaño del dato")
	fmt.Println("uint8/16/32/64: indica el tipo de dato int pero sin signo, solo positivos, y además el tamaño máximo del dato")
	fmt.Println("int uint : toma el tamaño de bits en el que está basado el procesador, con signo - sin signo")
	fmt.Println("float32: 32 bits con signo")
	fmt.Println("float64: 64 bits con signo")
	fmt.Println("string: se deben declarar con comillas dobles")
	fmt.Println("bool: true o false")
	fmt.Println("complex64: números complejos, real e imaginario float32")
	fmt.Println("complex128: números complejos, real e imaginario float64")
}

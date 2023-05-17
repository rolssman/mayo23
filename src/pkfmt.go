package main

import "fmt"

func main() {
	//declaración de variables
	helloMessage := "Hello"
	worldMessage := "World"

	//println
	fmt.Println(helloMessage, " ", worldMessage)
	fmt.Println(helloMessage, " ", worldMessage)

	//printf
	nombre := "Rolas"
	Horas := 200
	//%s es para string y %d para entero
	fmt.Printf("%s tiene más de %d pesos\n", nombre, Horas)
	//con %v no tienes que tipo de dato irá ahí
	fmt.Printf("%v tiene más de %v pesos\n", nombre, Horas)

	//Sprintf
	mesagge := fmt.Sprintf("%s tiene menos de %d dineros, dice", nombre, Horas)
	fmt.Println(mesagge)

	//Tipo de dato
	fmt.Printf("HelloMessage: %T\n", helloMessage)
	fmt.Printf("Cursos: %T\n", nombre)
}

package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
		valor1++
	} else {
		fmt.Println("No es 1")
	}
	fmt.Println("")
	if valor1 != valor2 {
		fmt.Println("Son diferentes los 2 valores")
	} else {
		fmt.Println("Son iguales los 2 valores")
	}

	//whit and
	if valor1 == 2 && valor2 == 2 {
		fmt.Println("valor1 y valor 2 son 2, mira con and ", valor1, valor2)
	}

	//with or
	if valor1 == 2 || valor2 == 1 {
		fmt.Println("Es verdad el or ira: ", valor1, valor2)
	}

	//convertir texto a número
	value, err := strconv.Atoi("15")
	if err != nil {
		log.Fatal(err) //Para imprimir en consola un error
	}
	fmt.Println("Valor convertido: ", value)
}

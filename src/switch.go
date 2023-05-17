package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("No es par")
	}

	//sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor que 0")
	default:
		fmt.Println("No condición")
	}
}

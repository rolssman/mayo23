package main

import "fmt"

func main() {
	//Make: Sirve para generar diccionarios
	m := make(map[string]int)

	m["jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrat valor
	value, ok := m["jose"]
	fmt.Println(value, ok)
}

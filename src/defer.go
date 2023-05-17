package main

import "fmt"

func main() {
	//Defer va a ejecutar la última funcion antes de que muera el sistema (como para cerrar las conexiones por ejemplo)
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continuea y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//Continue
		if i == 2 {
			fmt.Println("El número es 2: ", i)
			continue
		}

		if i == 8 {
			fmt.Println("EL número es 8: ", i)
			break
		}
	}
}

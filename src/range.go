package main

import (
	"fmt"
	"strings"
)

func isPalindromo(palabra string) {
	text := strings.ToLower(palabra)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		//Recorremos el texto que resivimos y lo vamos agregando en una nueva variable pero en reversa para verificar si es palindromo
		//El for junto con el parseo a string van desarmando la palabra por caracter y va en reversa leyendo cada caracter y agregandolo a la nueva variable
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"hola", "qué", "hace"}

	for valor := range slice {
		fmt.Println(valor)
	}

	isPalindromo("Ama")
}

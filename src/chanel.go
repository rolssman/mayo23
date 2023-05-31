package main

import "fmt"

//para salida es en c chan del lado izquierdo c <-chan y de entrada sería del lado contrario c text<-
func say(text string, c chan<- string){
	c <- text
}

func main(){
	//Se escribe make (Canal con chan / tipo de dato que va a pasar por el canal / *Opcional, cuántos datos van a pasar por el canal -Si no hay dato se vuelve dinamico-)
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
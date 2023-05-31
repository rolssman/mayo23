package main

import ("fmt"
"sync"
"time"
)

func say(text string, wg *sync.WaitGroup){
//defer es lo último que se ejecuta
	defer wg.Done()

	fmt.Println(text)
}

func main(){
	//Acumula un conjunto de go routines
	var wg sync.WaitGroup

	
	fmt.Println("Hello")
	//Agrega 1 go rutine al grupo de las wait group
	wg.Add(1)
//Se le manda una impresión y un puntero del wait group
	go say("World", &wg)
//Para esperar a todas las rutinas del wait group
	wg.Wait()

	go func(text string){
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
package main

import "fmt"

func main() {
	//Array: Un arreglo en go es impermutable (que su dimención no cambia)
	var array [4]int
	array[0] = 1
	array[1] = 8
	//Len: Sirve para ver la longitud de un arreglo
	//Cap: Sirve para conocer la capacidad maxima de datos que tiene el arreglo
	fmt.Println(array, len(array), cap(array))

	//Slice: Un slice si es mutable
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	//El elemento 0
	fmt.Println(slice[0])
	//Hasta el elemento número 3... El número después de los dos puntos es exclusivo
	fmt.Println(slice[:3])
	//Del numero 2 al número 4... el 2 es inclusivo y el 4 es exclusivo
	fmt.Println(slice[2:4])
	//Del número 4 en adelante
	fmt.Println(slice[4:])

	//append: Agregar datos al slice
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append, agregando una nueva lista al slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

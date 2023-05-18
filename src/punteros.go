package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (pc pc) ping() {
	fmt.Println(pc.brand, "pong")
}

func (pc *pc) duplicatedRam() {
	pc.ram = pc.ram * 2
}

func main() {
	a := 50
	//el & en la variable b contiene la dirección de memoria
	b := &a

	//*b para acceder al valor de la dirección de memoria
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100

	fmt.Println(a)

	pc := pc{ram: 8, disk: 500, brand: "DELL"}
	fmt.Println(pc)

	pc.ping()

	fmt.Println(pc.ram)

	pc.duplicatedRam()
	fmt.Println(pc.ram)
	pc.duplicatedRam()
	fmt.Println(pc.ram)
}

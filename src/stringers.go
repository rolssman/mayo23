package main

import "fmt"

type pc struct{
	ram int
	disk int
	brand string
}

func (compu pc) String() string{
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de disco, y es una %s", compu.ram, compu.disk, compu.brand)
}

func main(){
	compu := pc{ram: 8, disk: 500, brand: "Lenovo"}
	fmt.Println(compu)

}
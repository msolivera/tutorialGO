package main

import "fmt"

var estado bool

func main() {
	estado = true
	//no  se usa parentesis salvo que tengas que agrupar varias condiciones
	if estado == true {
		fmt.Println(estado)

	} else {
		fmt.Println("es falso")
	}

	//puedo llegar a asignar valores dentro del if y despues preguntar por el mismo:
	/*if estado = false; estado == true {
		fmt.Println(estado)

	} else {
		fmt.Println("es falso")
	}*/
	switch numero:= 5, numero{
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("mayor a 5")
	}
}

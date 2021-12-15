package main

import "fmt"

//iteraciones en go
func main() {

	//i := 1
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}

	//numero := 0
	//for {
	//	fmt.Println("continuo")
	//	fmt.Println("ingrese el numero secreto")
	//	fmt.Scan(&numero)
	//como romper un ciclo for
	//	if numero == 100 {
	//		break
	//	}
	//}

	//var i = 0
	//for i < 10 {
	//	fmt.Printf("\n valor de i: %d", i)
	//	if i == 5 {
	//		fmt.Printf(" multiplicamos por 2 \n")
	//		i = i * 2
	//CONTINUE SIRVE PARA ROMPER EL FOR Y VOLVER AL INICIO
	//		continue
	//	}
	//	fmt.Printf("            paso por aqui \n")
	//	i++
	//}

	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("voy a rituna sumando 2 a i")

			//GOTO sire para yo romper el bucle y hacer que continue desde donde yo quiera
			goto RUTINA
		}
		fmt.Println("valor de i:", i)
		i++
	}

}

//quede en video 11

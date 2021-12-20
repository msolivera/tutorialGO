package main

import "fmt"

//creo variable de tipo funcion
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	//funcion anonima, es una funcion que no tiene nombre y que se puede modificar ren tiempo de ejecucion
	//fmt.Println("Sumo 5 + 7 =  ", Calculo(5, 7))

	//ahora< redefino calculo
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2

	}
	//fmt.Println("Resto 6 - 4 =  ", Calculo(6, 4))
	//operaciones()

	//closures
	//variable para iniciar el numero
	tablaDel := 2
	//variable que la convertimos en nuestro closure
	//le pasamos a convertir en una func pasandole la funcion de tabla(de algo)
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		//cuando yo ejecuto Tabla() estoy ejecutando el return de la funcion
		fmt.Println(MiTabla())
	}
}
func operaciones() {
	resultado := func() int {
		var a int = 10
		var b int = 50

		return a + b

	}
	fmt.Println(resultado())
}

//closures pueden acceder a variables creadas por fuera de la funcion
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	//cuando yo ejecuto Tabla, estoy ejecutando este parte (en la variable MiTabla)
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

package main

import "fmt"

func main() {

	fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println("-------------")
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(2, 2, 2))
	fmt.Println(Calculo(10, 10, 10, 10))

}

//esa es una funcion comun y corriente
func uno(numero int) int {

	return numero * 2
}

//esta es una funcion que devuelve dos valores
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

//puedo creae vvariables y nombrarlas y despues devolverlas
func tres(numero int) (z int) {
	z = numero * 2
	return
}

//ejemplo de funcion que recibe una cantidad variable de parametros
func Calculo(numero ...int) int {
	total := 0
	//range devuelve una lista de los que tenga a la derecha (al ser variable lo convierte en lista)
	//range siempre devuelve dos cosas 1) numero del elemento 2) el elemento
	//en go cuando algo devuelve dos valores y yo no quiero usarlo, se usa el _ para alojarlo (porque no lo puedo obviar)
	// _ no reserva memoria0
	for _, num := range numero {

		total = total + num
	}
	return total
}

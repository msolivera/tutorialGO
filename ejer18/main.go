package main

import "fmt"

/*son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben
y devuelven los mismos tipos de variables*/

var result int

func main() {

	//los middlewares son para encapsular una funcion o una ejecucion preexistente en otra
	fmt.Println("inicio")
	//llamo el middleware
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(25, 5)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//el middleware se usa para funciones que recben y devuelven el mismo tipo devariable
//en su definicion tengo que poner la funcion que recibe (su tipo) y el tipo de func que devuelve
func operacionesMidd(funcion func(int, int) int) func(int, int) int {

	return func(a, b int) int {
		fmt.Println("inicio de la operacion")
		return funcion(a, b)
	}

}

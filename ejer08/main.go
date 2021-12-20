package main

import "fmt"

//1er forma de crear un vector (generica)
//var tabla [10]int

//como crear una matriz
var matriz [5][7]int

func main() {

	//slice vectores dinamicos, vectores que puedoo ampliarlos etc en ejecicion
	vectorDinamico := []int{2, 5, 4}
	//2da forma de crear un vector
	tabla := [10]int{5, 6, 98, 1, 4, 0, 3, 65, 8, 99}
	//tabla[0] = 1
	//tabla[5] = 15
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	matriz[3][5] = 1
	//fmt.Println(tabla)
	fmt.Println(matriz)
	fmt.Println(vectorDinamico)
	variante2()
	variante3()
	variante4()

}

func variante2() {

	// crear un slice a partir de un vector mas grande
	vector := [5]int{1, 2, 3, 4, 5}
	porcion := vector[3:]
	fmt.Println(porcion)
}

func variante3() {

	//crear un slice con un comando especial y dotarlo de capacidad para adicionar elementos en algun momento de ejecuciions
	//priner posicion de que tipo es el slice
	//cantidad de lementos inicial
	//cantidad de lugares maximos que reserva en memoria
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d , capacidad %d", len(elementos), cap(elementos))
}

func variante4() {

	//como nums no tiene espacio, para agregar datos hay que hacer append, no se puede asignar derecho
	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		//se van agregando y redimencionando a la vez
		//lo ideal en realidad es usar append cuando supere el largo porque es mas costoso
		nums = append(nums, i)
	}
	fmt.Printf("\n largo %d, capacidad %d", len(nums), cap(nums))
}

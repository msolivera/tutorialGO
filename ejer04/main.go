package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {

	//asi se hace para agarrar info desde el teclado
	fmt.Println("ingrese numero 1")
	fmt.Scanln(&numero1)

	fmt.Println("ingrese numero 2")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	//creo el escanner que on la libreria bufio sirve para leer escribir archvos etc.
	//uso la libreria os para indicarle al scaner de donde va a recibir la entrada Stdin = el teclado
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {

		leyenda = scanner.Text()
		//fmt.Printf("%d", numero1+numero2)
	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

}

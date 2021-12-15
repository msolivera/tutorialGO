package main

import (
	"fmt"
	"strconv"
)

//como se declaran las variables?
//1) variables globales
var Numero int
var texto string
var flag bool = true

//2)SCOUP DE VARIABLES: cuando las variables empiezan con minuscula se puede acceder solo desde el paquete actual, cuando
//empiezan en mayuscula, la puede ver desde otro paquetes. lo mismo con las funciones

func main() {

	//variables locales
	//var numero2 int
	//uso := para setearle un nuevo valor a la variable
	//numero3 := 4
	//uso = cuando quier cambiar el valor de una variable ya existente
	//numero3 = 15
	//crear varias variables juntas
	//var numeroA, numeroB, numeroC int

	var moneda int64 = 500
	numero1, numero2, numero3, texto, flag := 1, 2, 3, "jelou", false
	//si yo ahora quiero convertir un dato en otro lo hago asi:
	numero1 = int(moneda) //encapsulo para poder parsear
	fmt.Println(numero1)
	//para convertir de numero a texto:
	//texto = fmt.Sprintf("%d", numero2) //%d es poque recibe un decimal sino tengo que decirle que vaa a recibir
	texto = strconv.Itoa(int(moneda)) //esta es la forma usando la libreria permite convertir varias cosas.
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(texto)
	//cuando la llamo aca, flag se toma como local
	fmt.Println(flag)
	mostrarStatus()

}
func mostrarStatus() {
	//cuando la llamo aca, flag se toma como global
	fmt.Println(flag)
}

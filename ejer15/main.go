package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go miNombreLento("Micaela")
	//go adelante de una funcion dice que la va a ejecutar de forma asincrona
	fmt.Println("Estoy aqui")
	var x string
	//esto va a pedir por pantalla un valor que va a ser asignado a x

	fmt.Scanln(&x)
	//recordar que si yo ejecuto el scan muentras esta funcionando la func asincona
	// al final de la ejecucion del scan voy a terminar el programa
	//nos falta aca la istruccion para que continue la funcion asinctrona

}
func miNombreLento(nombre string) {
	//esto me va aseparar cada letra por un ""
	//letras convierte nombre en un vector
	letras := strings.Split(nombre, "")

	for _, letras := range letras {
		//range toma el vector y por cada elemento hace una iteracion
		time.Sleep(1000 * time.Millisecond)
		//casa un segundo hago el sleep y muestro una letra
		fmt.Println(letras)
	}
}

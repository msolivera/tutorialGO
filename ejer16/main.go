package main

import (
	"fmt"
	"time"
)

func main() {

	//los canales son espacion de memoria de dialogo
	//en la cual hablan varias rutinas, cuando se aloja algo en ese espacio
	//la rutina que pide un valor a cambio atua en consecuencia.
	//es espacio es un tipo de dato

	//creo un canal
	//time.Duration es el tipo de dato del canal
	canal1 := make(chan time.Duration)

	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")

	//para escuchar el canal 1
	msg := <-canal1
	fmt.Println(msg)

}

func bucle(canal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000; i++ {

	}
	final := time.Now()
	//sub devuelve la duracion de iempo entre inicio y fin
	canal <- final.Sub(inicio)

}

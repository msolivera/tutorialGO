package main

import (
	"log"
)

func main() {
	/*archivo := "prueba.txt"

	//construccion de funcion open
	f, err := os.Open(archivo)
	//cuando pongo una instruccion
	//seguido por un defer, esa instuccion no se ejecuta
	//no secuencialmente, sino cuando salga de la funcion
	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}
	//el defer se ejecura siempre auque el archivo no existe
	//libero buffer de memoria.*/

	//panic lo usa go para forzar un error
	ejemploPanic()

}

func ejemploPanic() {
	//todo esto se usa para que el panic no me pare el programa y se escriba todo en un log
	//pongo defer para controlar el panic
	defer func() {
		//se crea funcion anonima para poder ar varias instrucciones
		reco := recover() //se fija si hubo un panic y lo trae
		if reco != nil {
			//fatalf fujciona como printf
			//%v trae el error que tiene el reco
			log.Fatalf("Error que genero el PANIC \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("SE ENCONTRO EL VALOR DE 1")
	}
}

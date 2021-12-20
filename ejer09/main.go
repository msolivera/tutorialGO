package main

import "fmt"

func main() {

	//como crear un mapa
	paises := make(map[string]string)

	paises["mexico"] = "D.F."
	paises["Argentina"] = "Bs.As"
	fmt.Println(paises)
	fmt.Println(paises["mexico"])

	//otra forma de crear mapa

	//clave tipo string pero valor tipo entero
	campeonato := map[string]int{
		"barcelona":   39,
		"real madrid": 20,
		"chivas":      10}

	//adicionar un item
	campeonato["river plate"] = 25
	campeonato["chivas"] = 55 //puedo modificar

	//como borrar un elemento
	delete(campeonato, "chivas")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {

		fmt.Printf("el equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	//como saber si un registro existe
	puntaje, existe := campeonato["nacional"]
	//puntaje, existe := campeonato["barcelona"]
	fmt.Printf("el puntaje capturado es %d, y el quipo existe %t \n", puntaje, existe)

}

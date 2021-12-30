package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	grabarArchivo()
	grabarArchivo2()
}

func leoArchivo() {
	//este abre lee y cierra el archivo, no me deja manipularlo mucho mas
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("hubo error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("hubo error")
	} else {
		//este metodo me da mas libertas porque puedo hacer cosas con las lineas
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Println("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func grabarArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("hubo error")
		return
	}
	fmt.Fprintln(archivo, "nueva linea")
	archivo.Close()
}

func grabarArchivo2() {
	filename := "./nuevoArchivo.txt"
	if Append(filename, "/n esta es una segunda linea") == false {
		fmt.Println("hubo error en segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("hubo error")
		return false
	}

	//_, se usa cuando una variabe o funcion devuelve un erro o variale
	//cuando no me interesa recibir esa info, pongo _,
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("hubo error")
		return false
	}
	return true
}

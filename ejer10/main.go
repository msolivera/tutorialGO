package main

import (
	"fmt"
	"time"

	us "C:\Program Files\Go\src\tutorialGO\ejer10\user"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "micaela olivera", time.Now(), true)
	fmt.Println(u.Usuario)
}

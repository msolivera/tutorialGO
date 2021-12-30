package main

import "fmt"

//creacion de interfaces
type serVivo interface {
	estaVivo() bool
}
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

//humanos
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

//no es necesario que el hombre implemente la interface en su definicion
//si tiene las mismas funciones que la interface go ya se da cuenta
func (h *hombre) respirar()      { h.respirando = true }
func (h *hombre) comer()         { h.comiendo = true }
func (h *hombre) pensar()        { h.pensando = true }
func (h *hombre) sexo() string   { return "hombre" }
func (h *hombre) estaVivo() bool { return h.vivo }

func (m *mujer) respirar()      { m.respirando = true }
func (m *mujer) comer()         { m.comiendo = true }
func (m *mujer) pensar()        { m.pensando = true }
func (m *mujer) sexo() string   { return "mujer" }
func (m *mujer) estaVivo() bool { return m.vivo }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {

	Pedro := new(hombre)
	HumanosRespirando(Pedro)
	Pedro.vivo = true

	Maria := new(mujer)
	HumanosRespirando(Maria)

	fmt.Printf("Estoy vivo = %t", estoyVivo(Pedro))
}

//creacion de interfaces CASO 2
/*
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

//humanos
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type mujer struct {
	hombre
}

//no es necesario que el hombre implemente la interface en su definicion
//si tiene las mismas funciones que la interface go ya se da cuenta
func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre{
		return "hombre"
	}else{
		return "mujer"
	}



func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func main() {

	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)
}
*/

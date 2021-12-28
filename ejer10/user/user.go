package user

import "time"

type Usuario struct {
	//definicion
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	//this *Usuario dice que va a impactar en la estructura de usuario, es para cuando se llama.
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}

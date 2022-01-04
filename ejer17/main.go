package main

import "net/http"

func main() {

	//cuando mi pagina vaya a / ejecuto home
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)

}

//recibimos los parametros del http cuando envio y recibo res
func home(w http.ResponseWriter, r *http.Request) {
	//cuando estoy en esta funcion muestro index
	http.ServeFile(w, r, "./index.html")
}

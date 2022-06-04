package main

import (
	"log"

	"github.com/JoseUgal/go-http-api/cmd/api/bootstrap"
)

func main() {
	// Gestionamos el desplegue de nuestra aplicación con un
	// método para encapsular la lógica en caso de querer
	// realizar un test futuro.
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
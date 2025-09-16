package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define una función que maneja las solicitudes a la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Escribe una respuesta para el cliente
		fmt.Fprintf(w, "¡Hola desde el backend de Go!")
	})

	// Imprime un mensaje en la consola para saber que el servidor está listo
	fmt.Println("Servidor Go escuchando en el puerto 8080...")

	// Inicia el servidor HTTP y lo mantiene en escucha en el puerto 8080
	// log.Fatal detendrá el programa si ocurre un error
	log.Fatal(http.ListenAndServe(":8080", nil))
}

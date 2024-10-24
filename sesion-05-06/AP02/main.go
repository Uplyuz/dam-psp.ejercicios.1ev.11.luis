package main

import (
	"fmt"
	"os"
	"strings"
)

// Función para contar letras 'a' y enviar el resultado a un canal
func contarLetrasA(nombreArchivo string, resultado chan int, errores chan error) {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		errores <- err // Enviar error por el canal de errores
		return
	}

	contenido := string(data)
	contador := strings.Count(contenido, "a") + strings.Count(contenido, "A")
	resultado <- contador // Enviar resultado al canal
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <nombre_archivo>")
		return
	}

	nombreArchivo := os.Args[1]

	// Crear canales para resultados y errores
	resultado := make(chan int)
	errores := make(chan error)

	go contarLetrasA(nombreArchivo, resultado, errores)

	// Manejar el resultado o el error usando select
	for {
		select {
		case res := <-resultado:
			fmt.Printf("La letra 'a' aparece %d veces en el archivo %s\n", res, nombreArchivo)
			return // Salir del bucle después de manejar el resultado

		case err := <-errores:
			fmt.Printf("Error al abrir el archivo: %s\n", err)
			return // Salir del bucle después de manejar el error
		}
	}
}

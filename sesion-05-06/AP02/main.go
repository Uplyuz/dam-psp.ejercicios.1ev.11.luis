package main

import (
	"fmt"
	"os"
	"strings"
)

func contarLetrasA(nombreArchivo string, resultado chan int, errores chan error) {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		errores <- err
		return
	}

	contenido := string(data)
	contador := strings.Count(contenido, "a") + strings.Count(contenido, "A")
	resultado <- contador
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <nombre_archivo>")
		return
	}

	nombreArchivo := os.Args[1]

	resultado := make(chan int)
	errores := make(chan error)

	go contarLetrasA(nombreArchivo, resultado, errores)

	for {
		select {
		case res := <-resultado:
			fmt.Printf("La letra 'a' aparece %d veces en el archivo %s\n", res, nombreArchivo)
			return
		case err := <-errores:
			fmt.Printf("Error al abrir el archivo: %s\n", err)
			return
		}
	}
}

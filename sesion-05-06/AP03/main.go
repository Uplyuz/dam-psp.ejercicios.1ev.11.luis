package main

import (
	"fmt"
	"os"
	"strings"
)

func contarLetrasA(nombreArchivo string, resultado chan string, errores chan error) {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		errores <- fmt.Errorf("error al abrir el archivo '%s': %w", nombreArchivo, err)
		return
	}

	contenido := string(data)
	contador := strings.Count(contenido, "a") + strings.Count(contenido, "A")
	resultado <- fmt.Sprintf("La letra 'a' aparece %d veces en el archivo '%s'", contador, nombreArchivo)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <nombre_archivo1> <nombre_archivo2> ...")
		return
	}

	resultado := make(chan string)
	errores := make(chan error)

	for _, nombreArchivo := range os.Args[1:] {
		go contarLetrasA(nombreArchivo, resultado, errores)
	}

	for i := 0; i < len(os.Args)-1; i++ {
		select {
		case res := <-resultado:
			fmt.Println(res)
		case err := <-errores:
			fmt.Println(err)
		}
	}
}

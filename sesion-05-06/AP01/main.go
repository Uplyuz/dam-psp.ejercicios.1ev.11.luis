package main

import (
	"fmt"
	"os"
	"strings"
)

func contarLetrasA(nombreArchivo string) error {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		return err
	}

	contenido := string(data)

	contador := strings.Count(contenido, "a") + strings.Count(contenido, "A")

	fmt.Println("La letra 'a' aparece", contador, " veces en el archivo", nombreArchivo)

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <nombre_archivo>")
		return
	}

	nombreArchivo := os.Args[1]

	if err := contarLetrasA(nombreArchivo); err != nil {
		fmt.Println("Error al abrir el archivo: \n", err)
	}
}

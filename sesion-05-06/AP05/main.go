package main

import (
	"fmt"
	"os"
	"strings"
)

// Función para contar letras 'letra' en el archivo
func contarLetras(nombreArchivo string, letra string, resultado chan string, errores chan error, cuenta chan int) {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		errores <- fmt.Errorf("error al abrir el archivo '%s': %w", nombreArchivo, err)
		cuenta <- 0 // Enviar 0 si hay un error
		return
	}

	contenido := string(data)
	contador := strings.Count(contenido, letra) // Contar solo la letra especificada
	resultado <- fmt.Sprintf("La letra '%s' aparece %d veces en el archivo '%s'", letra, contador, nombreArchivo)
	cuenta <- contador // Enviar la cuenta parcial
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <letra> <nombre_archivo1> <nombre_archivo2> ...")
		return
	}

	letra := os.Args[1] // La letra a contar
	if len(letra) != 1 {
		fmt.Println("Por favor, proporciona solo un carácter como letra a contar.")
		return
	}

	resultado := make(chan string)
	errores := make(chan error)
	cuenta := make(chan int)

	// Iterar sobre los argumentos (nombres de archivos)
	for _, nombreArchivo := range os.Args[2:] {
		go contarLetras(nombreArchivo, letra, resultado, errores, cuenta)
	}

	total := 0                      // Variable para el total de apariciones de la letra
	numArchivos := len(os.Args) - 2 // Número de archivos

	// Recibir resultados y errores
	for i := 0; i < numArchivos; i++ { // Espera el número de archivos
		select {
		case res := <-resultado:
			fmt.Println(res) // Imprimir el resultado
		case err := <-errores:
			fmt.Println(err) // Imprimir el error
		case parcial := <-cuenta:
			total += parcial // Sumar la cuenta parcial al total
		}
	}

	// Mostrar el total de apariciones de la letra
	fmt.Printf("Total de apariciones de la letra '%s' en todos los archivos: %d\n", letra, total)
}
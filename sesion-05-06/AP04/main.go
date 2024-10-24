package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func contarLetrasA(nombreArchivo string, wg *sync.WaitGroup, resultado chan string, errores chan error, cuenta chan int) {
	defer wg.Done()

	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		errores <- fmt.Errorf("error al abrir el archivo '%s': %w", nombreArchivo, err)
		cuenta <- 0
		return
	}

	contenido := string(data)
	contador := strings.Count(contenido, "a") + strings.Count(contenido, "A")
	resultado <- fmt.Sprintf("La letra 'a' aparece %d veces en el archivo '%s'", contador, nombreArchivo)
	cuenta <- contador
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <nombre_archivo1> <nombre_archivo2> ...")
		return
	}

	resultado := make(chan string)
	errores := make(chan error)
	cuenta := make(chan int)

	var wg sync.WaitGroup

	for _, nombreArchivo := range os.Args[1:] {
		wg.Add(1)
		go contarLetrasA(nombreArchivo, &wg, resultado, errores, cuenta)
	}

	go func() {
		wg.Wait()
		close(resultado)
		close(errores)
		close(cuenta)
	}()

	total := 0

	// Recibir resultados y errores
	for i := 0; i < len(os.Args)-1; i++ {
		select {
		case res, ok := <-resultado:
			if ok {
				fmt.Println(res)
			}
		case err, ok := <-errores:
			if ok {
				fmt.Println(err)
			}
		case parcial := <-cuenta:
			total += parcial
		}
	}

	fmt.Printf("Total de apariciones de la letra 'a' en todos los archivos: %d\n", total)
}

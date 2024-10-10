package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <texto> <archivo>")
		return
	}

	//variables que recogen los argumentos del comando
	texto := os.Args[1]
	archivo := os.Args[2]

	file, err := os.Create("EJ01.1.salida.txt")

	// Crear el comando "grep" con los argumentos proporcionados
	cmd := exec.Command("grep", texto, archivo)

	cmd.Stdout = file

	err = cmd.Run()

	if err != nil {
		fmt.Println("Error ejecutando el comando:", err)
		return
	}
	// Imprimir la salida del comando
	fmt.Println("Salida redirigida a EJ01.1.salida.txt")
}

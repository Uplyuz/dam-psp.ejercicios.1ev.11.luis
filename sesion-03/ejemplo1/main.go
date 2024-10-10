package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// Crear el comando "cat" que espera recibir datos por stdin y los imprime
	cmd := exec.Command("cat")

	// Obtener el pipe de entrada estándar (stdin) del comando
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creando el pipe stdin:", err)
		return
	}

	stdout, err := cmd.StdoutPipe()

	// Iniciar el comando
	if err := cmd.Start(); err != nil {
		fmt.Println("Error iniciando el comando:", err)
		return
	}

	// Escribir datos en el pipe stdin
	io.WriteString(stdin, "Hola desde Go!\n")
	stdin.Close() // Es importante cerrar stdin para indicar que ya no enviaremos más datos

	// Obtener la salida del comando (stdout)
	//output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error obteniendo la salida:", err)
		return
	}

	output, err := io.ReadAll(stdout)

	// Imprimir la salida del comando
	fmt.Println(string(output))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// variables de entorno
	envVars := os.Environ()

	// Iterar sobre cada variable de entorno
	for _, env := range envVars {
		fmt.Println(env)
	}
}

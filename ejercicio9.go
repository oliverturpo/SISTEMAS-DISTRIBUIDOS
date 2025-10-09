package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese su nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	saludo := saludar(nombre)
	fmt.Println(saludo)
}

func saludar(nombre string) string {
	return "Bienvenido a Sistemas Distribuidos " + nombre
}

package main

import "fmt"

func main() {
	var nombre string
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)
	saludo := saludar(nombre)
	fmt.Println(saludo)
}

func saludar(nombre string) string {
	return "Bienvenido a Sistemas Distribuidos " + nombre
}

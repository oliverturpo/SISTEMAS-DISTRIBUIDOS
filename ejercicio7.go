package main

import "fmt"

func main() {
	nombre := "Cliver"
	saludo := saludar(nombre)
	fmt.Println(saludo)

}

func saludar(nombre string) string {
	return "Bienvenido a Sistemas Distribuidos " + nombre

}

package main

import "fmt"

func main() {
	// Definición de variable corta
	nombre := "Cliver"

	var nombre2 string = "Turpo"

	var (
		peso   = 75.5
		altura = 1.70
		estado = true
	)

	var edad int = 22
	var temperatura float64 = 37.0
	var aprobado bool = true

	fmt.Println(nombre, nombre2, peso, altura, edad, estado)
	fmt.Println(edad, temperatura, aprobado)
}

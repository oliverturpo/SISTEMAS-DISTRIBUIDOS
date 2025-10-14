package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Este es el cliente de saludo Computadora B")

	//pedir nombre del usuario
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese su nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	//conectarse al servidor
	fmt.Println("Conectando con el servidor")
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return
	}

	fmt.Println("Se esta enviando el nombre al servidor")
	conn.Write([]byte(nombre))

	//recibir saludo del servidor
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	respuesta := string(buffer[:n])
	fmt.Println("Saludo del servidor: ", respuesta)
	conn.Close()
}

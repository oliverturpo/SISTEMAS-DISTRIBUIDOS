package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Este es el servidor de saludo Computadora C")
	fmt.Println("Escuchando las conexiones...")

	listener, _ := net.Listen("tcp", "localhost:9000")
	for {
		conn, _ := listener.Accept()
		fmt.Println("Cliente conectado")

		//leer el nombre del cliente

		buffer := make([]byte, 1024)
		n, _ := conn.Read(buffer)
		nombre := string(buffer[:n])
		fmt.Println("Recibi el nombre: ", nombre)

		//enviar saludo al cliente
		respuesta := "Hola " + nombre + " Bienvenido a sistema de Distribucion "
		conn.Write([]byte(respuesta))
		fmt.Println("Enviar Saludo")
		conn.Close()
	}
}

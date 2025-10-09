package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Numero: ", i)
	}

	numero := 2
	for i := 0; i <= 10; i++ {

		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}

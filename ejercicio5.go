package main

import "fmt"

func main() {
	contador := 0
	/*for contador < 10 {
		fmt.Println("Contador: ", contador)
		contador++
	}*/

	for {

		if contador >= 20 {
			break
		}
		fmt.Println("Contador: ", contador)
		contador++
	}

}

package main

import "fmt"

func main() {
	// exemplo e variável com o tipo explicito
	var variavel1 string = "Variável 1"

	// exemplo de variável com o tipo oculto
	variavel2 := "Variável 2"

	// definindo multiplas variáveis de uma vez
	var(
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	// definindo multiplas variáveis sem inferência de tipo
	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4, variavel5, variavel6)
}
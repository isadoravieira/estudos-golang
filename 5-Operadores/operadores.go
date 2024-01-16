package main

import "fmt"

func main() {

	// '=' Operador de atribuição no qual precisa indicar o tipo da variável
	var variavel1 string = "Erick Lindo"

	// ':=' Operador de atribuição no qual não precisa indicar o tipo da vaariável
	variavel2 := "Laila Linda"

	fmt.Println(variavel1, variavel2)

	// para negar o valor de uma variável
	variavel := true
	fmt.Print(!variavel)

	// OBS: não existe operador ternário em golang
}
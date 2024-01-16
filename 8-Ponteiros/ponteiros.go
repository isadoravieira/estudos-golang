package main

import "fmt"

// ponteiros são variáveis que salvam dentro delaas um endereço de memória de algo
// é uma referência de memória

func main(){

	var variavel int
	var ponteiro *int

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro)
	fmt.Println(&ponteiro) // verifica onde o ponteiro está armazenado
	fmt.Println(*ponteiro)//desreferênciação -> mostra o valor que está salvo dentro do endereço
}
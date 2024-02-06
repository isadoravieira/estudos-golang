package main

import "fmt"

// funções recursivas chamam a si mesmas.
// esse tipo de função deve ter uma condição de parada, para não gerar um estouro de pilha.
// funções recursivas não são recomendadas se é necessário fazer várias execução.

func fibonacci(posicao uint) uint{

	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main(){
	fmt.Println("\n--- Funções Variáticas ---")

	posicao := uint(83)

	fmt.Println("A posição", posicao, "contém o número", fibonacci(posicao),".")
}
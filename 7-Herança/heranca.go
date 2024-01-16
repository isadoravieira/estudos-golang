package main

import "fmt"

// não existe herança em golang, pois golang não é uma linguagem orientada a objetos
// mas é possível dar um "jeitinho" -> é uma pseudo herança

type pessoa struct {
	nome   string
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main(){

	pessoa1 := pessoa{"Roberto", 178}

	estudante1 := estudante{pessoa1, "Arquitetura", "Unicamp"}

	fmt.Println(estudante1)

	// a vantagem disso é que para acessar somente o nome/altura de um estudante, não é necessário dar um print em uma instancia de pessoa
	// é possivel acessa-lo diretamente pela instancia de estudante:
	fmt.Println(estudante1.nome, estudante1.altura)
}
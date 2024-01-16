package main

import "fmt"

// structs são coleções de campos, na qual cada campo tem um nome e um tipo
// seria tipo uma classe na orientação a objeto
// é possível criar instâncias dessa struct

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

// é possível ter structs dentro de structs
type endereco struct {
	logadouro string
	numero    int16
}

func main() {
	//  formas de instanciar uma struct

	var u usuario
	u.nome = "Daniel"
	u.idade = 23
	fmt.Println(u)

	endereco1 := endereco{"Rua dos Abacates", 22}

	user := usuario{"Juliana", 45, endereco1}
	fmt.Println(user)

	user2 := usuario{
		idade: 12,
	}
	fmt.Println(user2)
}

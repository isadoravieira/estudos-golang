package main

import "fmt"

// a diferença de metodo para uma função, é que um método sempre vai estar associado a algo, como uma struct ou interface

type usuario struct {
	nome  string
	idade uint
}

// o seguinte método está associado a uma struc 
// e ele basicamente diz que todo usuario possui uma função salvar
func (u usuario) salvar(){
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados.\n", u.nome)
}

func main() {
	usuario1 := usuario{"Fernanda",35}
	usuario1.salvar()
}

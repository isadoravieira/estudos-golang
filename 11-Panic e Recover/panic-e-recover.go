package main

import "fmt"

// a função panic interrompe o fluxo normal do programa -> para tudo -> entra em panico
// diferente do erro, que continnua executando  normalmente o programa
// para recuperar um programa que está entrando em panico, utiliza-se a cláusula recover 

func recuperarFuncao(){
	if r := recover(); r != nil{
		fmt.Println("Função recuperada com sucesso!")
	}
}

func alunoApravado(n1, n2 uint) bool{
	defer recuperarFuncao()

	media := (n1 + n2) / 2
	
	if media > 6 {
		fmt.Println(true)
	} else if media < 6 {
		fmt.Println(false)
	}
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main(){

	alunoApravado(6, 6)
	fmt.Println("Fim do programa.")
}
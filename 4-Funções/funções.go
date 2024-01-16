package main

import "fmt"

// entre parenteses vão os parametros  e após isso vai o tipo do retorno
func somar(n1 int8, n2 int8) int8{
	return n1 + n2
}

// é permiitido retornar mais de um valor
func calculos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main(){
	somar := somar(25, 23)
	fmt.Println(somar)

	resultadoSoma, resultadoSubtracao := calculos(32, 54)
	fmt.Println("Soma: ",resultadoSoma,"\nSubtração: ",resultadoSubtracao)
}
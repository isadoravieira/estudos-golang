package main

import (
	"errors"
	"fmt"
)

func main() {
	// existem quatros tipos de inteiros em go, a diferença é apenas o tamanho
	// cada um suporta a sus determinaada quantidade de bits
	// quando você não específica o tamanho ele irá usar como base a arquitetura do seu pc
	// int8

	// unsygned int - int sem sinal
	// uint

	// int32 = rune -> equivale aos caracteres da tabela ascii

	// existem dois tipos de numeros reais em go
	// seguem a mesma lógica dos inteiros
	// float32

	// ao usar aspas simples e printar, é mostrado o numero na tabela ascii
	char := 'R'
	fmt.Println(char)

	// existe o tipo error em go
	// o valor zero dele é nil
	// nil é um tipo de dados que serve como valor zero para muitas coisas: interface, erro, ponteiro
	var erro error
	fmt.Println(erro)

	// essa variável não retona uma string, ela retorna um erro 
	var erro2 error = errors.New("Erro interno ...")
	fmt.Println(erro2)
}
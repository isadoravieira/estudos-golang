package main

import "fmt"

func main() {

	// são estruturas de chave-valor
	fmt.Println("Maps:")

	pessoa := map[string]string{ // o tipo da chave vai dentro do colchetes e o tipo do valor vai ao lado
		"nome":      "Pedro",
		"sobrenome": "Rogerio",
	}

	fmt.Println(pessoa["nome"]) // chamando um campo especifico

	// maps aninhados:
	pessoa2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Luis",
			"ultimmo":  "Ricardo",
		},
		"abacaxi":{
			"gosta ?": "sim",
		},
		"curso": {
			"nome":   "Engenharia Química",
			"campus": "Campus São Paulo",
		},
	}

	fmt.Println(pessoa2)
	// por padrão, o go retorna as chaves de um map em ordem alfabética (no caso de strings) ou em ordem crescente (no caaso de números)
	// SAÍDA DA PESSOA2:
	// map[abacaxi:map[gosta ?:sim] curso:map[campus:Campus São Paulo nome:Engenharia Química] nome:map[primeiro:Luis ultimmo:Ricardo]]

	delete(pessoa2, "nome") // deletando uma chave e respectivamente, o seu valor
}

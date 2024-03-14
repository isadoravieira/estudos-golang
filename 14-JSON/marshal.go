package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func cachorroJson() {
	// json.Marshal() // transforma um struct/map em um json
	// json.Unmarshal() // transforma um json em uma struct/map

	c := cachorro{"Erick", "Vira-lata", 3}

	cachorroJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	// para printar um json: 
	fmt.Println(bytes.NewBuffer(cachorroJson))
}

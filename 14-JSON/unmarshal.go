package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// a struct com os campos em json devem estar no mesmo arquivo do unmarshal
type cachorro2 struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	cachorro := `{"nome":"Laila", "raca":"Vira-lata", "idade":13}`

	var c cachorro2

	if err := json.Unmarshal([]byte(cachorro), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}

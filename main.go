package main

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
}

func main() {
	var req Request

	if data, err := os.ReadFile("data/test.json"); err != nil {
		log.Fatalln("Falhou ao abrir arquivo de teste:", err)

	} else if err = json.Unmarshal(data, &req); err != nil {
		log.Fatalln("Falhou ao converter a requisição:", err)
	}

	res, err := req.Create()
	if err != nil {
		log.Fatalln("Falhou ao gerar a massa de testes:", err)
	}

	if err = res.Save("data/test.csv"); err != nil {
		log.Fatalln("Falhou ao salvar o arquivo com a massa de testes:", err)
	}

	log.Println("Massa de testes gerada com sucesso!")
}

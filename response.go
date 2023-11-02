package main

import (
	"log"
	"os"
	"strings"
)

type Response []string

func (res *Response) Add(line string) {
	(*res) = append((*res), line)
}

func (res *Response) Save(filename string) error {
	content := strings.Join(*res, "\n")

	if err := os.WriteFile("data/test.csv", []byte(content), 0666); err != nil {
		log.Fatalln("Falhou ao gerar arquivo com a massa de testes:", err)
	}

	return nil
}

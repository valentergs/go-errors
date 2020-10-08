package main

import (
	"fmt"
	"log"
)

type errorString string

// Aqui a funçao Error() sofre um by-pass atraves de uma string que aramazena a mensagem para imprimir um erro customizado.
func (es errorString) Error() string {
	return fmt.Sprintf("Aqui está o erro: %s", string(es))
}

func main() {
	n, err := soma(1, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

func soma(i, j int) (int, error) {
	n := i * j
	if n != i+j {
		// Aqui é possivel assignar uma mensagem ao esso do tipo errorString
		var sErr errorString = "a soma não funcionou"
		log.Fatalln(0, sErr)
	}
	return n, nil
}

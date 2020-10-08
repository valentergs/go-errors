package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/valentergs/go-errors/copiar"
)

func main() {
	destino := "arq2.txt"
	origem := "arq1.txt"
	err := copiar.CopiarArq(destino, origem)
	if err != nil {
		// A checagem do tipo erro (errors.Is), geralmente é feito no main(), não na função de origem
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Arquivo %s não existe\n", origem)
		}
		log.Println("erro ao exec copiarArq():", err)
	}
}

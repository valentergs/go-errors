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
	var pathErr *os.PathError
	err := copiar.CopiarArq(destino, origem)
	if err != nil {
		// A checagem do tipo erro (errors.Is ou errors.As), geralmente é feito no main(), não na função de origem
		if errors.As(err, &pathErr) {
			fmt.Printf("ARQUIVO: %s, OPERATION: %s, ORIGINAL ERROR: %s\n", pathErr.Path, pathErr.Op, err)
		}
		log.Println("erro ao exec copiarArq():", err)
	}
}

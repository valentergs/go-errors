package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/*
		fOne, err := os.Open("fileOne.txt")
		if err != nil {
			fmt.Errorf("Não consegui abrir o arquivo: %w", err)
		}
		fOne.Close()

		fTwo, err := os.Create("fileTwo.txt")
		if err != nil {
			fmt.Errorf("Não consegui criar o arquivo: %w", err)
		}
		fTwo.Close()

		if _, err := io.Copy(fTwo, fOne); err != nil {
			fmt.Errorf("não consegui copiar: %w", err)
		}
	*/

	destino := "arq2.txt"
	origem := "arq1.txt"
	err := copiarArq(destino, origem)
	if err != nil {
		// A checagem do tipo erro (errors.Is), geralmente é feito no main(), não na função de origem
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Arquivo %s não existe\n", origem)
		}
		log.Println("erro ao exec copiarArq():", err)
	}
}

func copiarArq(destino, origem string) error {

	// Abrir o arquivo de origem
	fOne, err := os.Open(origem)
	if err != nil {
		return fmt.Errorf("Não consegui abrir o arquivo: %w", err)
	}
	defer fOne.Close()

	// Criar e abrir arquivo de destino
	fTwo, err := os.Create(destino)
	if err != nil {
		return fmt.Errorf("Não consegui criar o arquivo: %w", err)
	}
	defer fTwo.Close()

	// Copiar arquivo no arquivo de destino o arquivo de origem
	i, err := io.Copy(fTwo, fOne)
	if err != nil {
		return fmt.Errorf("não consegui copiar: %w", err)
	}
	fmt.Println("quantos bytes foram copiados", i)

	// Retorna nil se não houver nenhum erro
	return nil
}

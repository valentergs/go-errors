package copiar

import (
	"fmt"
	"io"
	"os"
)

// CopiarArq copia o conteudo de um arquivo de origem para um arquivo de destino e verifica possiveis erros
func CopiarArq(destino, origem string) error {

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
	n, err := io.Copy(fTwo, fOne)
	if err != nil {
		return fmt.Errorf("não consegui copiar: %w", err)
	}
	fmt.Println("quantos bytes foram copiados", n)

	// Retorna nil se não houver nenhum erro
	return nil
}

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("somefile.txt")

	// O jeito mas comum de tratar um erro é iniciar um panic()
	//panic(err)

	// O jeito recomendado é utilizar fmt.Errorf() e errors.Is, que compara erros.
	// if errors.Is(err, os.ErrPermission) {
	// 	err = fmt.Errorf("você não tem permissão para abrir arquivo: %w", err)
	// 	log.Println(err)
	// } else if errors.Is(err, os.ErrNotExist) {
	// 	err = fmt.Errorf("arquivo não existe %w", err)
	// 	log.Println(err)
	// }

	var pathError *os.PathError

	// // errors.As assinala ao "Target", nesse caso *os.PathError, o valor do erro pretendido.
	// if errors.As(err, &pathError) {
	// 	// return fmt.Errorf("Arquivo : %w e error path %s", err, pathError.Path)
	// 	fmt.Printf("Error: %s\nPath: %s\nPathOp: %s\nPathErr: %s\n", err, pathError.Path, pathError.Op, pathError.Err)
	// } else {
	// 	fmt.Println(err)
	// }

	switch {
	case errors.Is(err, os.ErrPermission):
		err = fmt.Errorf("você não tem permissão para abrir arquivo: %w", err)
		log.Println(err)
	case errors.Is(err, os.ErrNotExist):
		err = fmt.Errorf("arquivo não existe %w", err)
		log.Println(err)
	case errors.As(err, &pathError):
		fmt.Printf("Error: %s\nPath: %s\nPathOp: %s\nPathErr: %s\n", err, pathError.Path, pathError.Op, pathError.Err)
	case err != nil:
		fmt.Println(err)
	}

	defer f.Close()

	//fmt.Println(err)

}

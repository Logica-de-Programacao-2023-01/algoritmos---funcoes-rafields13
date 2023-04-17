// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string com todas as strings concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
)

func concmaisvirgula(slice []string) (string, error) {

	if len(slice) == 0 {

		return "", fmt.Errorf("não tem nenhum texto para ser lido")

	}

	soma := slice[0]

	for i := 1; i < len(slice); i++ {

		soma += ", "

		soma += slice[i]

	}

	return soma, nil

}

func main() {

	result, err := concmaisvirgula([]string{"Hello", "world!"})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("A concatenação desses elementos separados por vírgula é: %s", result)

}

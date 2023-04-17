// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo apenas as strings que começam com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func apenasmaiusculas(slice []string) ([]string, error) {

	if len(slice) == 0 {

		return []string{""}, fmt.Errorf("o Slice está vazio")

	}

	var mapamaiusculas = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var maiusculas []string

	for i := 0; i < len(slice); i++ {

		for j := 0; j < len(slice); j++ {

			if strings.HasPrefix(slice[i], mapamaiusculas[j]) {

				maiusculas = append(maiusculas, slice[i])

			}

		}

	}

	return maiusculas, nil

}

func main() {

	result, err := apenasmaiusculas([]string{"Hello", "world"})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("As palavras que começam apenas com letra maiúscula são: %s", result)

}

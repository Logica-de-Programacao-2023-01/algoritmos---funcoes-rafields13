// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func crescente(slice []int) ([]int, error) {

	if len(slice) == 0 {

		return []int{0}, fmt.Errorf("o Slice está vazio")

	}

	var ordenados []int

	for i := 0; i < len(slice); i++ {

		for j := len(slice); j >= 0; j-- {

			if slice[i] > slice[j] {

				slice[i], slice[j] = slice[j], slice[i]

				ordenados = append(ordenados, slice[j])

			}

		}

	}

	return ordenados, nil

}

func main() {

	result, err := crescente([]int{1, 2, 7, 3, 4, 6, 5})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("Os valores de forma ordenada são: %d", result)

}

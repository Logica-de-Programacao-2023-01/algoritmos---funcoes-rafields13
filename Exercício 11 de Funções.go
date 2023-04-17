// Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo e falso caso contrário. Caso o número seja negativo, retorne um erro.

package main

import "fmt"

func verprimo(x int) (bool, error) {

	if x < 0 {

		return false, fmt.Errorf("o número é negativo")

	}

	slice := []int

	for i := 0; i < x

}

func main() {

}

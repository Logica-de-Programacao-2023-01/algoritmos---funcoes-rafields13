// Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string. Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func palavras(s string) ([]string, error) {

	if strings.Count(s, "") == 0 {

		return []string{""}, fmt.Errorf("não tem nenhum texto para ser lido")

	}

}

func main() {

}

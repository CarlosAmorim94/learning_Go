package main

import (
	"errors"
	"fmt"
)

func main() {
	value, error := sum(9, 2)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(value)

}

func sum(a, b int) (int, error) {
	if a+b > 10 {
		return 0, errors.New("Valor maior que 10")
	}
	return a + b, nil
}

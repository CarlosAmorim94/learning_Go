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

	fmt.Println(sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func sum(a, b int) (int, error) {
	if a+b > 10 {
		return 0, errors.New("Valor maior que 10")
	}
	return a + b, nil
}
func sum2(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

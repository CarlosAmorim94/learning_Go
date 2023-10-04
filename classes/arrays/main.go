package main

import "fmt"

func main() {
		var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[1])
	fmt.Println(meuArray[len(meuArray)-1])
	fmt.Println("Tamanho do meu array", len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[1:]), cap(s[1:]), s[1:])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	salarios := map[string]int {
		"Carlos": 1000,
		"João": 2000,
		"Maria": 3000,
	}

	fmt.Println(salarios)
	fmt.Println(salarios["Carlos"])

	delete(salarios, "Carlos")
	salarios["Car"] = 5000


	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

}
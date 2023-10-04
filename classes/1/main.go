package main

import "fmt"

const a = "Hello World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Carlos"
	e float64 = 12.3
	f ID      = 123456
)

func main() {
	a := "x"
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

	fmt.Printf("O tipo de E é %T\n", f)
	fmt.Printf("O tipo de E é %T\n", e)


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
}
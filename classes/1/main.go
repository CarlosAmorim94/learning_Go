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

	fmt.Printf("O tipo de E é %T", f)
	fmt.Printf("O tipo de E é %T", e)
}
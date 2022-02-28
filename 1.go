package main

import "fmt"

const prefixoOla = "Olá, "

func Hi(parameter string) string {
	if parameter == "" {
		parameter = "Mundo"
	}
	return prefixoOla + parameter
}

func main() {
	fmt.Println(Hi("Walisson"))
}

package main

import "fmt"

func Hi() string {
	return "Olá, Mundo"
}

func HiWithParameters(parameter string) string {
	return "Olá, " + parameter
}

func main() {
	fmt.Println(Hi())
}

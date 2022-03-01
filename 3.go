package main

func Repetir(param string, qtd int) (word string) {

	for i := 0; i < qtd; i++ {
		word += param
	}
	return
}

package main

func Soma(param []int) (resultado int) {
	// O range trabalha com chave e valor, o underscore é para ignorar um parameter, no caso o chave
	for _, numero := range param {
		resultado += numero

	}

	return

}

package main

func Soma(param []int) (resultado int) {
	// O range trabalha com chave e valor, o underscore é para ignorar um parameter, no caso o chave
	for _, numero := range param {
		resultado += numero

	}

	return

}

func SomaTudo(numeroParaSomar ...[]int) (resultados []int) {
	for _, quantidadeDeNumeros := range numeroParaSomar {
		resultados = append(resultados, Soma(quantidadeDeNumeros))

	}
	return
}

func SomaTodoOResto(soma ...[]int) (resultado []int) {

	for _, quantidadeDeNumeros := range soma {
		if len(quantidadeDeNumeros) == 0 {
			resultado = append(resultado, 0)
		} else {
			semPrimeiro := quantidadeDeNumeros[1:]
			resultado = append(resultado, Soma(semPrimeiro))
		}
	}
	return
}

//https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/arrays-e-slices

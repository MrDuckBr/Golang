package main

const espanhol = "Espanhol"
const frances = "Frances"
const prefixoOla = "Olá, "
const prefixoEspanhol = "Hola, "
const prefixoFrances = "Bonjour, "

func Hi(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoIdiomas(idioma) + nome

}
func prefixoIdiomas(idioma string) (prefixo string) {

	switch idioma {
	case frances:
		prefixo = prefixoFrances
	case espanhol:
		prefixo = prefixoEspanhol
	default:
		prefixo = prefixoOla
	}
	return

}

func main() {

}

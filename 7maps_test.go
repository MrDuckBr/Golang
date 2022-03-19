//https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/maps
package main

import (
	"testing"
)

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "Isso é um teste"}

	t.Run("Palavra Conhecida", func(t *testing.T) {

		resultado, _ := dicionario.Busca("teste")
		esperado := "Isso é um teste"

		comparaString(t, resultado, esperado)
	})

	t.Run("Palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("Não Existe")

		comparaErros(t, erroNaoEncontrado, err)
	})

}

func comparaString(t *testing.T, esperado, resultado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado: %s Esperado: %s", resultado, esperado)
	}

}

func comparaErros(t *testing.T, esperado, erro error) {
	t.Helper()

	if esperado != erro {
		t.Errorf("Resultado: %s Esperado: %s", erro, esperado)
	}

}

func TestAdiciona(t *testing.T) {
	t.Run("Palavra nova", func(t *testing.T) {

		dicionario := Dicionario{}
		palavra := "Teste adiciona"
		definicao := "Isso é apenas um teste"

		err := dicionario.Adiciona(palavra, definicao)

		comparaErros(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("Palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")

		comparaErros(t, err, erroPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)

	})

}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)

	if err != nil {
		t.Fatal("Não foi possivel encontrar a palavra buscada")
	}
	if definicao != resultado {
		t.Errorf("Resultado: %s Esperado: %s", resultado, definicao)
	}
}

func TestUpdate(t *testing.T) {

	t.Run("palavra existente", func(t *testing.T) {

		palavra := " teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		novaDefinicao := "Nova Definicao"

		err := dicionario.Atualiza(palavra, novaDefinicao)
		comparaErros(t, err, nil)

		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("Palavra Nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)

		comparaErros(t, err, erroPalavraInexistente)
	})

}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definicao de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != erroNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}

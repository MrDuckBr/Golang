package main

import "testing"

func TestSoma(t *testing.T) {

	t.Run("5 Numeros somados", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado: %d  esperado: %d numeros: %v", resultado, esperado, numeros)
		}
	})

	t.Run("3 Numeros somados", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado: %d  esperado: %d numeros: %v", resultado, esperado, numeros)
		}
	})
}

// https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/arrays-e-slices

package main

import (
	"reflect"
	"testing"
)

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

func TestSomaTudo(t *testing.T) {

	t.Run("Resultado de dois vetores", func(t *testing.T) {
		resultado := SomaTudo([]int{1, 2}, []int{0, 9})
		esperado := []int{3, 9}
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado: %v  esperado: %v ", resultado, esperado)
		}

	})
}

func TestSomaTodoOResto(t *testing.T) {

	verificaSoma := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado: %v  esperado: %v ", resultado, esperado)
		}

	}

	t.Run("Somando todo o resto", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSoma(t, resultado, esperado)
	})

	t.Run("Slice Vazio", func(t *testing.T) {

		resultado := SomaTodoOResto([]int{}, []int{0, 9})
		esperado := []int{0, 9}
		verificaSoma(t, resultado, esperado)

	})
}

// https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/arrays-e-slices

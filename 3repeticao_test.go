package main

import "testing"

func TestRepetir(t *testing.T) {
	resultado := Repetir("a", 3)
	esperado := repeatEsperado("a", 3)

	if esperado != resultado {
		t.Errorf("Esperado: %s , Resultado: %s", esperado, resultado)
	}

}

func repeatEsperado(a string, b int) string {
	var palavra string
	for i := 0; i < b; i++ {
		palavra += a
	}
	return palavra
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 3)
	}
}

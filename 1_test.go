package main

import "testing"

func TestHi(t *testing.T) {
	resultado := Hi()
	esperado := "Olá, Mundo"

	if resultado != esperado {
		t.Errorf("esperado %s, resultado %s", esperado, resultado)
	}
}

func TestHiWithParameters(t *testing.T) {
	resultado := HiWithParameters("Walisson")
	esperado := "Olá, Walisson"

	if resultado != esperado {
		t.Errorf("esperado %s, resultado %s", esperado, resultado)
	}

}

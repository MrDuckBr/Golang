package main

import "testing"

func TestHi(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, esperado, resultado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("esperado %s, resultado %s", esperado, resultado)
		}

	}

	t.Run("Quando vazia ela tem que imprimir 'ola mundo'", func(t *testing.T) {
		resultado := Hi("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, esperado, resultado)

	})

	t.Run("Quando passar uma string deve completar 'ola, walisson'", func(t *testing.T) {
		resultado := Hi("Walisson", "")
		esperado := "Olá, Walisson"
		verificaMensagemCorreta(t, esperado, resultado)

	})

	t.Run("Ola  em espanhol", func(t *testing.T) {
		resultado := Hi("Walisson", "Espanhol")
		esperado := "Hola, Walisson"
		verificaMensagemCorreta(t, esperado, resultado)

	})

	t.Run("Ola  em frances", func(t *testing.T) {
		resultado := Hi("Walisson", "Frances")
		esperado := "Bonjour, Walisson"
		verificaMensagemCorreta(t, esperado, resultado)

	})

}
